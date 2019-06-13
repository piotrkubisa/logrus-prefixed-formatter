package prefixed_test

import (
	"testing"

	. "github.com/piotrkubisa/logrus-prefixed-formatter"
	"github.com/sirupsen/logrus"
)

type LogOutput struct {
	buffer string
}

func (o *LogOutput) Write(p []byte) (int, error) {
	o.buffer += string(p[:])
	return len(p), nil
}

func (o *LogOutput) GetValue() string {
	return o.buffer
}

func TestFormatter(t *testing.T) {
	var formatter *TextFormatter
	var log *logrus.Logger
	var output *LogOutput

	beforeEach := func() {
		output = new(LogOutput)
		formatter = new(TextFormatter)
		log = logrus.New()
		log.Out = output
		log.Formatter = formatter
		log.Level = logrus.DebugLevel
	}

	t.Run("logfmt output", func(t *testing.T) {
		t.Run("should output simple message", func(t *testing.T) {
			beforeEach()

			formatter.DisableTimestamp = true
			log.Debug("test")
			got := output.GetValue()
			expected := "level=debug msg=test\n"
			if got != expected {
				t.Errorf("Incorrect output; expected: %s, got: %s", expected, got)
				t.FailNow()
			}
		})

		t.Run("should output message with additional field", func(t *testing.T) {
			beforeEach()

			formatter.DisableTimestamp = true
			log.WithFields(logrus.Fields{"animal": "walrus"}).Debug("test")
			got := output.GetValue()
			expected := "level=debug msg=test animal=walrus\n"
			if got != expected {
				t.Errorf("Incorrect output; expected: %s, got: %s", expected, got)
				t.FailNow()
			}
		})
	})

	t.Run("Formatted output", func(t *testing.T) {
		t.Run("should output formatted message", func(t *testing.T) {
			beforeEach()

			formatter.DisableTimestamp = true
			formatter.ForceFormatting = true
			log.Debug("test")
			got := output.GetValue()
			expected := "DEBUG test\n"
			if got != expected {
				t.Errorf("Incorrect output; expected: %s, got: %s", expected, got)
				t.FailNow()
			}
		})
	})
}
