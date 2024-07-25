package ghost

import (
	"fmt"
	"strings"
)

func filter[T any](arr []T, validator func(T) bool) (ret []T) {
	for _, v := range arr {
		if validator(v) {
			ret = append(ret, v)
		}
	}

	return
}

func splitUrl(url string) []string {
	return filter(strings.Split(url, "/"), func(path string) bool { return path != "" })
}

func createResponse(statusCode uint, statusCodeText string, contentType string, message string) string {
	contentLength := len(message)

	response := fmt.Sprintf("HTTP/1.1 %d %v\r\n", statusCode, statusCodeText) +
		fmt.Sprintf("Content-Type: %v\r\n", contentType) +
		fmt.Sprintf("Content-Length: %d\r\n", contentLength) +
		"\r\n" +
		message

	return response
}

