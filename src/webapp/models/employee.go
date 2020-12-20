package models

type Employee struct {
	EmployeeId int64 `json:"employee_id"`
	EmployeeName string `json:"employee_name"`
	EmployeeEmail string `json:"employee_email"`
}