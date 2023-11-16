package clog

import (
	"fmt"
	"net/http"
)

func HTTPResponse(req *http.Request, code int, additional ...string) string {
	responseStatus := fmt.Sprintf("%d %s", code, http.StatusText(code))
	category := code / 100

	switch category {
	case 1:
		responseStatus = fmt.Sprintf("\033[97m%s\033[0m", responseStatus)
	case 2:
		responseStatus = fmt.Sprintf("\033[32m%s\033[0m", responseStatus)
	case 3:
		responseStatus = fmt.Sprintf("\033[33m%s\033[0m", responseStatus)
	case 4:
		responseStatus = fmt.Sprintf("\033[31m%s\033[0m", responseStatus)
	case 5:
		responseStatus = fmt.Sprintf("\033[91m%s\033[0m", responseStatus)
	}

	out := fmt.Sprintf(
		"%s - \033[1m\"%s %s\"\033[0m %s",
		req.RemoteAddr, req.Method, req.URL.Path, responseStatus,
	)

	for _, text := range additional {
		out = fmt.Sprintf("%s %s", out, text)
	}

	return out
}
