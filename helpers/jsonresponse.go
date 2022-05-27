package helpers

import (
	"bootcamp/hacktiv8-assigment2/params"
)

func JsonResponse(status int, message *string, additionalInfo *string, payload interface{}) params.Response {
	var response params.Response
	if additionalInfo != nil {
		response.Status = status
		response.Message = message
		response.AdditionalInfo = additionalInfo
		response.Payload = nil
	} else {
		response.Status = status
		response.Message = message
		response.AdditionalInfo = nil
		response.Payload = payload
	}

	return response
}
