package controllers

import (
	"encoding/json"
	"github.com/bugbountychris8691/go-mvc-rest-api/src/webapp/errors"
	"github.com/bugbountychris8691/go-mvc-rest-api/src/webapp/services"
	"net/http"
	"strconv"
)

func GetEmployeeById(writer http.ResponseWriter, request *http.Request) {
	idFromRequest := request.URL.Query().Get("employee_id")
	employeeId, err := strconv.ParseInt(idFromRequest, 10, 64)
	// or
	//employeeId, err := strconv.ParseInt(request.URL.Query().Get("employee_id"), 10, 64)

	if err != nil {
		apiErr := &errors.AppError{
			Message: "employee_id must be a number",
			StatusCode: http.StatusBadRequest,
			Status: "Bad Request",
		}
		jsonValue, _ := json.Marshal(apiErr)
		writer.WriteHeader(apiErr.StatusCode)
		writer.Write(jsonValue)
		return
	}

	employee, apiErr := services.FindEmployeeById(employeeId)

	if apiErr != nil {
		jsonValue, _ := json.Marshal(apiErr)
		writer.WriteHeader(apiErr.StatusCode)
		writer.Write(jsonValue)
		return
	}

	jsonValue, _ := json.Marshal(employee)
	writer.Write(jsonValue)
}