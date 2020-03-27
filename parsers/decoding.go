package parsers

import (
	"bytes"
	"fmt"
	"github.com/saintfish/chardet"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/encoding/unicode"
	"io"
	"io/ioutil"
	"os"
)

// Sadly, Go doesn't seem to support auto-detecting charsets, and we might get funky encodings in our input files
// Chardet can automatically detect charset for us, but returns charset as string. Here we convert
// the most common charsets in their string representation to actual encodings
var charsetEncodings = map[string]encoding.Encoding{
	"UTF-8":        unicode.UTF8,
	"ISO-8859-1":   charmap.ISO8859_1,
	"ISO-8859-2":   charmap.ISO8859_2,
	"ISO-8859-5":   charmap.ISO8859_5,
	"ISO-8859-6":   charmap.ISO8859_6,
	"ISO-8859-7":   charmap.ISO8859_7,
	"ISO-8859-8":   charmap.ISO8859_8,
	"ISO-8859-8-I": charmap.ISO8859_8I,
	"UTF-16BE":     unicode.UTF16(unicode.BigEndian, unicode.UseBOM),
	"UTF-16LE":     unicode.UTF16(unicode.LittleEndian, unicode.UseBOM),
	"windows-1251": charmap.Windows1251,
	"windows-1252": charmap.Windows1252,
	"windows-1253": charmap.Windows1253,
	"windows-1254": charmap.Windows1254,
	"windows-1255": charmap.Windows1255,
	"windows-1256": charmap.Windows1256,
}

func DecodeFile(path string) (io.Reader, error) {
	inFile, e := os.Open(path)
	if e != nil {
		return nil, e
	}
	defer inFile.Close()
	d := chardet.NewTextDetector()
	content, e := ioutil.ReadAll(inFile)
	if e != nil {
		return nil, e
	}
	res, e := d.DetectBest(content)
	if e != nil {
		return nil, e
	}
	enc, ok := charsetEncodings[res.Charset]
	if !ok {
		return nil, fmt.Errorf(
			"Could not find charset decoder for `%s`",
			res.Charset)
	}
	return enc.NewDecoder().Reader(bytes.NewReader(content)), nil
}
