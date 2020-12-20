package services

import (
	"github.com/bugbountychris8691/go-mvc-rest-api/src/webapp/dao"
	"github.com/bugbountychris8691/go-mvc-rest-api/src/webapp/errors"
	"github.com/bugbountychris8691/go-mvc-rest-api/src/webapp/models"
)

func FindEmployeeById(employeeId int64)(*models.Employee, *errors.AppError) {
	return dao.FindEmployeeById(employeeId)
}