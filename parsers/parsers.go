package parsers

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/golang-collections/go-datastructures/bitarray"
	"io"
	"strings"
	"unicode"
	"unicode/utf8"
)

type Table struct {
	Filename string
	Headers  []string
	Rows     [][]string
}

type Parser interface {
	Parse(io.Reader) (Table, error)
}

type CSVParser struct{}
type FixedWidthParser struct {
	Columns []string
}

type EmptyFileError struct{}

func (e EmptyFileError) Error() string {
	return "This file is empty"
}

func (p CSVParser) Parse(r io.Reader) (Table, error) {
	var t Table

	records, e := csv.NewReader(r).ReadAll()
	if e != nil {
		return t, e
	}
	if len(records) < 1 {
		return t, EmptyFileError{}
	}
	headers := records[0]
	var rows [][]string
	if len(records) > 1 {
		rows = records[1:]
	}

	t = Table{Headers: headers, Rows: rows}

	return t, nil
}

func readLines(r io.Reader) []string {
	scanner := bufio.NewScanner(r)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

type ColNotFound struct {
	column string
}

func (e ColNotFound) Error() string {
	return fmt.Sprintf("Column '%s' not found in file", e.column)
}

func (p FixedWidthParser) Parse(r io.Reader) (Table, error) {
	var t Table
	lines := readLines(r)
	if len(lines) < 1 {
		return t, EmptyFileError{}
	}
	headerLine := lines[0]

	var offsets []int
	var headers []string
	if len(p.Columns) < 1 {
		// no columns, infer offsets
		offsets = inferOffsets(lines)
		// infer headers
		headers = str2slice(headerLine, offsets)
	} else {
		var e error
		offsets, e = calculateOffsets(p.Columns, headerLine)
		if e != nil {
			return t, e
		}
		headers = p.Columns
	}

	var rows [][]string
	if len(lines) > 1 {
		for _, line := range lines[1:] {
			rows = append(rows, str2slice(line, offsets))
		}
	}
	t = Table{Headers: headers, Rows: rows}

	return t, nil
}

func inferOffsets(lines []string) []int {
	var bitmasks []bitarray.BitArray
	max := 0

	// Find max line length
	for _, line := range lines {
		c := utf8.RuneCountInString(line)
		if c > max {
			max = c
		}
	}

	// For each line, create a bitarray where each bit indicates
	// if that rune is a space (0) or not (1)
	for _, line := range lines {
		b := bitarray.NewBitArray(uint64(max))
		utf8line := []rune(line)
		for i := 0; i < len(utf8line); i++ {
			if !unicode.IsSpace(utf8line[i]) {
				b.SetBit(uint64(i))
			}
		}
		bitmasks = append(bitmasks, b)
	}

	// logical OR all the bitarrays, results in bitarray
	// where a bit is 0 if that rune is a space in every line
	combined := bitarray.NewBitArray(uint64(max))
	for _, m := range bitmasks {
		combined = combined.Or(m)
	}

	// detect edges by logical XORing neighbouring bits
	var edges []int
	edges = append(edges, 0)
	skip := true
	for i := 0; i < max-1; i++ {
		a, _ := combined.GetBit(uint64(i))
		b, _ := combined.GetBit(uint64(i + 1))
		if a != b {
			// we're only interested in leading edges (start of non-space sequence)
			if !skip {
				edges = append(edges, i+1)
			}
			skip = !skip
		}
	}
	return edges
}

func calculateOffsets(columns []string, header string) ([]int, error) {
	var offsets []int
	for _, col := range columns {
		if i := strings.Index(strings.ToLower(header), strings.ToLower(col)); i < 0 {
			return nil, ColNotFound{col}
		} else {
			// To get the proper index when column names contain non-ascii runes,
			// count the # of runes up to this index
			offsets = append(offsets, utf8.RuneCountInString(header[:i]))
		}
	}
	return offsets, nil
}

func str2slice(s string, offsets []int) []string {
	n := len(offsets)
	var row []string
	for j := 0; j < n-1; j++ {
		// To get the proper slice in case of non-ascii runes,
		// convert to rune slice, take slice, convert back
		row = append(row, string([]rune(s)[offsets[j]:offsets[j+1]]))
		//row = append(row, s[offsets[j]:offsets[j + 1]])
	}
	row = append(row, string([]rune(s)[offsets[n-1]:]))
	return row
}
