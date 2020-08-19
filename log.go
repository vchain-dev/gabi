package gabi

import (
	"github.com/vchain-dev/gabi/revocation"
	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func init() {
	Logger = logrus.StandardLogger()
	revocation.Logger = Logger
}
