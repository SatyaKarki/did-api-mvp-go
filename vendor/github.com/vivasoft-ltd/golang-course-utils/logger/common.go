package logger

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"reflect"
	"runtime"
	"runtime/debug"
	"strings"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func NewFileLoggerClient(filePath string) *CustomLogger {
	logFile, err := openLogFile(filePath)
	if err != nil {
		return nil
	}
	logger = log.New()
	multiWriter := io.MultiWriter(os.Stdout, logFile)
	logger.SetOutput(multiWriter)
	return &CustomLogger{
		client: logger,
	}
}

func (r *CustomLogger) SetLogLevel(level logrus.Level) {
	r.client.Level = level
}

func (r *CustomLogger) SetLogFormatter(formatter logrus.Formatter) {
	r.client.Formatter = formatter
}

func (r *CustomLogger) SetLogJsonFormatter() {
	r.client.Formatter = &logrus.JSONFormatter{}
}

// Debug logs a message at level Debug on the CustomLogger.
func (r *CustomLogger) Debug(args ...interface{}) {
	if r.client.Level >= logrus.DebugLevel {
		entry := r.client.WithFields(logrus.Fields{})
		//entry.Data["file"] = fileInfo(2)
		entry.Debug(args...)
	}
}

// DebugWithFields Debug logs a message with fields at level Debug on the CustomLogger.
func (r *CustomLogger) DebugWithFields(l interface{}, f map[string]interface{}) {
	if r.client.Level >= logrus.DebugLevel {
		entry := r.client.WithFields(f)
		//entry.Data["file"] = fileInfo(2)
		entry.Debug(l)
	}
}

// Info logs a message at level Info on the CustomLogger.
func (r *CustomLogger) Info(args ...interface{}) {
	if r.client.Level >= logrus.InfoLevel {
		entry := r.client.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Info(args...)
	}
}

// InfoWithFields Debug logs a message with fields at level Debug on the CustomLogger.
func (r *CustomLogger) InfoWithFields(l interface{}, f map[string]interface{}) {
	if r.client.Level >= logrus.InfoLevel {
		entry := r.client.WithFields(f)
		//entry.Data["file"] = fileInfo(2)
		entry.Info(l)
	}
}

// Warn logs a message at level Warn on the CustomLogger.
func (r *CustomLogger) Warn(args ...interface{}) {
	if r.client.Level >= logrus.WarnLevel {
		entry := r.client.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Warn(args...)
	}
}

// WarnWithFields Debug logs a message with fields at level Debug on the CustomLogger.
func (r *CustomLogger) WarnWithFields(l interface{}, f map[string]interface{}) {
	if r.client.Level >= logrus.WarnLevel {
		entry := r.client.WithFields(f)
		entry.Data["file"] = fileInfo(2)
		entry.Warn(l)
	}
}

// StdError logs a message at level Error on the CustomLogger.
func (r *CustomLogger) StdError(args ...interface{}) {
	if r.client.Level >= logrus.ErrorLevel {
		entry := r.client.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Error(args...)
	}
}

// Error logs a message at level Error on the CustomLogger and sends alert to slack
//
// if 1 item in args then there will be no metadata
//
// if multiple items in args then 1st item will be treated as metadata and rest items will go for args
func (r *CustomLogger) Error(args ...interface{}) {
	if len(args) > 1 {
		args = args[1:]
	}

	if r.client.Level >= logrus.ErrorLevel {
		entry := r.client.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Error(args...)

	}
}

// Error logs a message at level Error on the CustomLogger. with res data and metaData
func (r *CustomLogger) ApiError(rs RequestResponseMap, metaData interface{}, args ...interface{}) {
	if r.client.Level >= logrus.ErrorLevel {
		entry := logger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Error(args...)
	}
}

// ErrorWithFields Debug logs a message with fields at level Debug on the CustomLogger.
func (r *CustomLogger) ErrorWithFields(l interface{}, f map[string]interface{}) {
	if r.client.Level >= logrus.ErrorLevel {
		entry := r.client.WithFields(f)
		entry.Data["file"] = fileInfo(2)
		entry.Error(l)
	}
}

// Fatal logs a message at level Fatal on the CustomLogger.
func (r *CustomLogger) Fatal(args ...interface{}) {
	if r.client.Level >= logrus.FatalLevel {
		entry := r.client.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Fatal(args...)

	}
}

// FatalWithFields Debug logs a message with fields at level Debug on the CustomLogger.
func (r *CustomLogger) FatalWithFields(l interface{}, f map[string]interface{}) {
	if r.client.Level >= logrus.FatalLevel {
		entry := r.client.WithFields(f)
		entry.Data["file"] = fileInfo(2)
		entry.Fatal(l)
	}
}

// Panic logs a message at level Panic on the CustomLogger.
func (r *CustomLogger) Panic(args ...interface{}) {
	if r.client.Level >= logrus.PanicLevel {
		entry := r.client.WithFields(logrus.Fields{
			"stack": string(debug.Stack()),
		})
		entry.Data["file"] = fileInfo(2)
		entry.Panic(args...)

	}
}

// PanicWithFields Debug logs a message with fields at level Debug on the CustomLogger.
func (r *CustomLogger) PanicWithFields(l interface{}, f fields) {
	if r.client.Level >= logrus.PanicLevel {
		entry := r.client.WithFields(logrus.Fields(f))
		entry.Data["file"] = fileInfo(2)
		entry.Panic(l)
	}
}

func (r *CustomLogger) fileInfo(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		if slash >= 0 {
			file = file[slash+1:]
		}
	}
	return fmt.Sprintf("%s:%d", file, line)
}

func (r *CustomLogger) fileAddressInfo(skip int) string {
	_, file, line, _ := runtime.Caller(skip)
	return fmt.Sprintf("%s:%d", file, line)
}

func (r *CustomLogger) processLog(args ...interface{}) string {
	var errMsgBuffer bytes.Buffer
	for _, arg := range args {
		if arg != nil {
			switch reflect.TypeOf(arg).Kind() {
			case reflect.String:
				errMsgBuffer.WriteString(arg.(string) + "\n")
			case reflect.Ptr:
				e := arg.(error)
				errMsgBuffer.WriteString(e.Error() + "\n")
			}
		}
	}

	return errMsgBuffer.String()
}

func openLogFile(filePath string) (*os.File, error) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	return file, nil
}
