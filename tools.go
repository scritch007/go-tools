package tools

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"strings"
	"unicode"
	"unicode/utf8"
)

func ComputeHmac256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
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

func CaseInsensitiveContains(s, substr string) bool {
	s, substr = strings.ToUpper(s), strings.ToUpper(substr)
	return strings.Contains(s, substr)
}
