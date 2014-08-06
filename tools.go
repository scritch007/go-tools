package tools

import (
	"io"
	"log"
	"strings"
	"unicode"
	"unicode/utf8"
)

var (
	LOG_DEBUG   *log.Logger
	LOG_INFO    *log.Logger
	LOG_WARNING *log.Logger
	LOG_ERROR   *log.Logger
)

func LogInit(
	traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	LOG_DEBUG = log.New(traceHandle,
		"DEBUG: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	LOG_INFO = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	LOG_WARNING = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	LOG_ERROR = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func Capitalize(s string) string {
	if s == "" {
		return ""
	}
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToUpper(r)) + s[n:]
}

func JsonToGolang(in *string) (out string) {
	res := strings.Split(*in, "_")
	out = ""
	for _, s := range res {
		out += Capitalize(s)
	}
	return out
}
