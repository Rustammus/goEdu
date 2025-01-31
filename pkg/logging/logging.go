package logging

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"runtime"
)

type writerHook struct {
	Writer   []io.Writer
	LogLevel []logrus.Level
}

func (hook *writerHook) Levels() []logrus.Level {
	return hook.LogLevel
}

func (hook *writerHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}

	for _, w := range hook.Writer {
		w.Write([]byte(line))
	}
	return nil
}

func init() {
	l := logrus.New()
	l.SetReportCaller(true)
	l.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		DisableQuote:  true,
		FullTimestamp: true,
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			filename := path.Base(frame.File)
			return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s:%d", filename, frame.Line)
		},
	})
	//err := os.MkdirAll("logs", 0644)
	//if err != nil {
	//	panic(err)
	//}
	//allLogsFile, err := os.OpenFile("logs/all.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0640)
	//if err != nil {
	//	panic(err)
	//}

	l.SetOutput(io.Discard)
	l.AddHook(&writerHook{
		Writer:   []io.Writer{os.Stdout},
		LogLevel: logrus.AllLevels,
	})

	l.SetLevel(logrus.TraceLevel)

	logEntry = logrus.NewEntry(l)
}

var logEntry *logrus.Entry

type Logger struct {
	*logrus.Entry
}

func GetLogger() Logger {
	return Logger{logEntry}
}
