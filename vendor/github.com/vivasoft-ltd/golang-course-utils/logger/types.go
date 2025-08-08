package logger

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

// fields wraps logrus.Fields, which is a map[string]interface{}
type fields logrus.Fields

// CustomLogger wraps logrus.Logger and provides additional functionality
type CustomLogger struct {
	client *logrus.Logger
}

type RequestResponseMap struct {
	Req     *http.Request
	ReqBody interface{}
	Res     *http.Response
	ResBody interface{}
}
