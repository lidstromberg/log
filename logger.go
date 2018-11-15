package log

import (
	"log"
	"os"
	"strconv"
	"strings"
)

//LogEvent writes a bunch of messages to the output source
func LogEvent(messages ...string) {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.LUTC)
	log.SetOutput(os.Stderr)
	log.Println(strings.Join(messages, ":"))
}

//IsLoggingOn uses a supplied string to determine if logging is on
func IsLoggingOn(key string) bool {
	result, err := strconv.ParseBool(key)

	if err != nil {
		log.Printf("IsLoggingOn could not parse %s", key)
		return false
	}

	return result
}
