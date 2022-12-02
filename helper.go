package echohelpers

import (
	"fmt"
	"strings"

	"github.com/labstack/gommon/log"
)

// LogLevelFromString takes a string representation of a log level and attempts to convert it to an echo log level.
// A PR has been opened to merge this functionality into Echo's log package, but has not been merged to date.
// https://github.com/labstack/gommon/pull/53
func LogLevelFromString(lvl string) (*log.Lvl, error) {
	var l log.Lvl
	switch strings.ToUpper(lvl) {
	case "DEBUG":
		l = log.DEBUG
	case "INFO":
		l = log.INFO
	case "WARN":
		l = log.WARN
	case "ERROR":
		l = log.ERROR
	case "OFF":
		l = log.OFF
	default:
		return nil, fmt.Errorf("not a valid log level: '%s'", lvl)
	}

	return &l, nil
}
