package helpers

import (
	"bootcamp/hacktiv8-assigment2/params"
)

func JsonResponse(status int, errorMessage *string, payload interface{}) params.Response {
	var response params.Response
	if errorMessage != nil {
		response.Status = status
		response.ErrorMessage = errorMessage
		response.Payload = nil
	} else {
		response.Status = status
		response.ErrorMessage = nil
		response.Payload = payload
	}

	return response
}
