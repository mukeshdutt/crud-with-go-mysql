package main

import (
	"encoding/json"
	"gin-crud-app/domain"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AllEmployee(c *gin.Context) {

	emps := GetAllEmployees()

	if len(emps) == 0 {
		c.JSON(http.StatusNoContent, "no content available")
		return
	}

	// If records available then will proceed to get the details for employee
	var empList []EmployeeModel
	for _, emp := range emps {
		empInfo := EmployeeModel{
			ID:      emp.ID,
			Name:    emp.Name,
			Age:     emp.Age,
			Gender:  emp.Gender,
			Mobile:  emp.Mobile,
			Email:   emp.Email,
			City:    emp.City,
			State:   emp.State,
			Country: emp.Country,
		}
		empList = append(empList, empInfo)
	}
	c.JSON(http.StatusOK, ResponseModel{IsSuccess: true, Data: empList})
}

func EmployeeByID(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseModel{IsSuccess: false, Message: "ID value is not valid."})
		return
	}

	empInfo := GetEmployeeInfoByID(id)
	if (empInfo == domain.Employee{}) {
		c.JSON(http.StatusBadRequest, ResponseModel{IsSuccess: false, Message: "Employee record not found for this ID."})
		return
	}

	emp := EmployeeModel{
		ID:      empInfo.ID,
		Name:    empInfo.Name,
		Age:     empInfo.Age,
		Gender:  empInfo.Gender,
		Mobile:  empInfo.Mobile,
		Email:   empInfo.Email,
		City:    empInfo.City,
		State:   empInfo.State,
		Country: empInfo.Country,
	}
	c.JSON(http.StatusOK, ResponseModel{IsSuccess: true, Data: emp})
}

func SaveEmployee(c *gin.Context) {

	empData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusOK, ResponseModel{IsSuccess: false, Message: "Invalid payload"})
		return
	}

	var emp EmployeeModel
	err = json.Unmarshal(empData, &emp)
	if err != nil {
		c.JSON(http.StatusOK, ResponseModel{IsSuccess: false, Message: "Invalid payload"})
		return
	}

	demp := domain.Employee{
		Name:    emp.Name,
		Age:     emp.Age,
		Gender:  emp.Gender,
		Mobile:  emp.Mobile,
		Email:   emp.Email,
		City:    emp.City,
		State:   emp.State,
		Country: emp.Country,
	}
	isSaved := AddEmployee(demp)
	if isSaved {
		c.JSON(http.StatusOK, ResponseModel{IsSuccess: true, Message: "Employee Information Successfully Modified."})
	}
}

func UpdateEmployee(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseModel{IsSuccess: false, Message: "employee id is missing in the request"})
		return
	}

	empData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusOK, ResponseModel{IsSuccess: false, Message: "Invalid payload"})
		return
	}

	var emp EmployeeModel
	err = json.Unmarshal(empData, &emp)
	if err != nil {
		c.JSON(http.StatusOK, ResponseModel{IsSuccess: false, Message: "Invalid payload"})
		return
	}

	demp := domain.Employee{
		Name:    emp.Name,
		Age:     emp.Age,
		Gender:  emp.Gender,
		Mobile:  emp.Mobile,
		Email:   emp.Email,
		City:    emp.City,
		State:   emp.State,
		Country: emp.Country,
	}
	isSaved := EditEmployee(id, demp)
	if isSaved {
		c.JSON(http.StatusOK, ResponseModel{IsSuccess: true, Message: "Employee Successfully onboarded."})
	}
}

func DeleteEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseModel{IsSuccess: false, Message: "ID value is not valid."})
		return
	}

	isDeleted := RemoveEmployee(id)
	if isDeleted {
		c.JSON(http.StatusBadRequest, ResponseModel{IsSuccess: false, Message: "Employee record deleted succssfully."})
		return
	} else {
		c.JSON(http.StatusBadRequest, ResponseModel{IsSuccess: false, Message: "Record not updated due to some error."})
	}
}
