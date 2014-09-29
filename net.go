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

//Look into the headers or into the Query parameters for the desired piece of information
func GetParameter(r *http.Request, parameter string) string {
	atoken := r.Header.Get("access_token")
	if 0 == len(atoken) {
		//Now look into the Query Parameters
		atoken = r.URL.Query().Get("access_token")
	}
	return atoken
}
