package dao

import (
	"github.com/bugbountychris8691/go-mvc-rest-api/src/webapp/errors"
	"fmt"
	"github.com/bugbountychris8691/go-mvc-rest-api/src/webapp/models"
	"net/http"
)

var(
	employees = map[int64]*models.Employee{
		101 : {EmployeeId: 101, EmployeeName: "Christopher", EmployeeEmail: "chrisg@email.com"},
	}
)

func FindEmployeeById(employeeId int64)(*models.Employee, *errors.AppError) {
	if employee := employees[employeeId]; employee != nil {
		return employee, nil
	}

	//return nil, errors.New(fmt.Sprintf("Employee %v was not found", employeeId))

	return nil, &errors.AppError{
		Message: fmt.Sprintf("Employee %v not found", employeeId),
		StatusCode: http.StatusNotFound,
		Status: "Not Found",
	}
}
