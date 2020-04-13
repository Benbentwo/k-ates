package util

import (
	"bytes"
	stackdriver "github.com/TV4/logrus-stackdriver-formatter"
	"github.com/fatih/color"
	"github.com/pkg/errors"
	"github.com/rickar/props"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

type LogLevel int
type FormatLayoutType string

const (
	DEBUG LogLevel = 0
	INFO  LogLevel = 1
	WARN  LogLevel = 2
	ERROR LogLevel = 3

	FormatLayoutJSON        FormatLayoutType = "json"
	FormatLayoutText        FormatLayoutType = "text"
	FormatLayoutStackdriver FormatLayoutType = "stackdriver"
)

var (
	white      = color.New(color.FgWhite).SprintFunc()
	logger     *logrus.Entry
	labelsPath = "/etc/labels"
)

func Log(level LogLevel, messages ...interface{}) {
	col, msg := getPretext(level)
	var messageTotal = ""
	for _, message := range messages {
		if msg, typeok := message.(string); typeok {
			messageTotal += msg
		} else {
			white(message)
		}
	}
	Logger().Infof("%s: %s", col(msg), messageTotal)
}

func Logger() *logrus.Entry {
	err := initializeLogger()
	if err != nil {
		logrus.Warnf("error initializing logrus %v", err)
	}
	return logger
}

func initializeLogger() error {
	if logger == nil {

		// if we are inside a pod, record some useful info
		var fields logrus.Fields
		if exists, err := fileExists(labelsPath); err != nil {
			return errors.Wrapf(err, "checking if %s exists", labelsPath)
		} else if exists {
			f, err := os.Open(labelsPath)
			if err != nil {
				return errors.Wrapf(err, "opening %s", labelsPath)
			}
			labels, err := props.Read(f)
			if err != nil {
				return errors.Wrapf(err, "reading %s as properties", labelsPath)
			}
			app := labels.Get("app")
			if app != "" {
				fields["app"] = app
			}
			chart := labels.Get("chart")
			if chart != "" {
				fields["chart"] = labels.Get("chart")
			}
		}
		logger = logrus.WithFields(fields)

		format := os.Getenv("JX_LOG_FORMAT")
		if format == "json" {
			setFormatter(FormatLayoutJSON)
		} else if format == "stackdriver" {
			setFormatter(FormatLayoutStackdriver)
		} else {
			setFormatter(FormatLayoutText)
		}
	}
	return nil
}

func getPretext(level LogLevel) (func(...interface{}) string, string) {
	var col func(...interface{}) string
	var msg string
	switch level {
	case DEBUG:
		msg = "DEBUG"
		col = color.New(color.FgHiCyan).SprintFunc()
	case INFO:
		msg = "INFO"
		col = color.New(color.FgGreen).SprintFunc()
	case WARN:
		msg = "WARN"
		col = color.New(color.FgYellow).SprintFunc()
	case ERROR:
		msg = "ERROR"
		col = color.New(color.FgHiRed).SprintFunc()
	}
	return col, msg
}

// copied from utils to avoid circular import
func fileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, errors.Wrapf(err, "failed to check if file exists %s", path)
}

// JenkinsXTextFormat lets use a custom text format
type TextFormat struct {
	ShowInfoLevel   bool
	ShowTimestamp   bool
	TimestampFormat string
}

func (f *TextFormat) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer

	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	level := strings.ToUpper(entry.Level.String())
	switch level {
	case "INFO":
		if f.ShowInfoLevel {
			b.WriteString((level))
			b.WriteString(": ")
		}
	case "WARNING":
		b.WriteString(ColorWarning(level))
		b.WriteString(": ")
	case "DEBUG":
		b.WriteString(ColorDebug(level))
		b.WriteString(": ")
	default:
		b.WriteString(ColorError(level))
		b.WriteString(": ")
	}
	if f.ShowTimestamp {
		b.WriteString(entry.Time.Format(f.TimestampFormat))
		b.WriteString(" - ")
	}

	b.WriteString(entry.Message)

	if !strings.HasSuffix(entry.Message, "\n") {
		b.WriteByte('\n')
	}
	return b.Bytes(), nil

}

// NewTextFormat creates the default text formatter
func NewTextFormat() *TextFormat {
	return &TextFormat{
		ShowInfoLevel:   false,
		ShowTimestamp:   false,
		TimestampFormat: "2006-01-02 15:04:05",
	}
}

// setFormatter sets the logrus format to use either text or JSON formatting
func setFormatter(layout FormatLayoutType) {
	switch layout {
	case FormatLayoutJSON:
		logrus.SetFormatter(&logrus.JSONFormatter{})
	case FormatLayoutStackdriver:
		logrus.SetFormatter(stackdriver.NewFormatter())
	default:
		logrus.SetFormatter(NewTextFormat())
	}
}
