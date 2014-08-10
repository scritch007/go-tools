package tools

import (
	"net/http"
)

//Get the client ip address...
func GetClientAddr(r *http.Request) string {
	ip := r.Header.Get("X-Real-IP")
	if 0 == len(ip) {
		ip = r.Header.Get("X-Forwarded-For")
		if 0 == len(ip) {
			ip = r.RemoteAddr
		}
	}
	return ip
}
