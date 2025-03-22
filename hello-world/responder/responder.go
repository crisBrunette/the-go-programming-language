package responder

import (
	"fmt"
	"net/http"
)

func HeaderWriter(req *http.Request) string {
	//var response string
	response := ""
	for _, value := range req.Header {
		for _, value := range value {
			response = fmt.Sprintln(response, value)
		}
	}
	return response
}
