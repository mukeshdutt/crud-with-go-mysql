package main

import (
	"database/sql"
	"gin-crud-app/domain"

	_ "github.com/go-sql-driver/mysql"
)

func getConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:@MySQL567@/employeemanagement")
	if err != nil {
		panic(err)
	}
	return db
}

func GetAllEmployees() []domain.Employee {
	var emps []domain.Employee

	result, err := getConnection().Query("select * from employee")
	if err != nil {
		panic(err)
	}
	for result.Next() {
		emp := domain.Employee{}
		err = result.Scan(&emp.ID, &emp.Name, &emp.Age, &emp.Gender, &emp.Email, &emp.Mobile, &emp.City, &emp.State, &emp.Country)
		if err != nil {
			panic(err)
		}
		emps = append(emps, emp)
	}
	return emps
}

func GetEmployeeInfoByID(id int) domain.Employee {

	result := getConnection().QueryRow("select * from employee where id=?", id)

	if result.Err() != nil {
		panic(result.Err().Error())
	}
	emp := domain.Employee{}
	err := result.Scan(&emp.ID, &emp.Name, &emp.Age, &emp.Gender, &emp.Email, &emp.Mobile, &emp.City, &emp.State, &emp.Country)
	if err != nil {
		panic(err)
	}
	return emp
}

func AddEmployee(emp domain.Employee) bool {

	stmt, err := getConnection().Prepare(`insert into Employee(Name, Age, Gender, Email, Mobile, City, State, Country) values(?,?,?,?,?,?,?,?);`)
	if err != nil {
		panic(err)
	}

	result, err := stmt.Exec(emp.Name, emp.Age, emp.Gender, emp.Email, emp.Mobile, emp.City, emp.State, emp.Country)
	if err != nil {
		panic(err)
	}
	count, _ := result.RowsAffected()
	return count > 0
}

func EditEmployee(id int, emp domain.Employee) bool {

	stmt, err := getConnection().Prepare(`update Employee set Name=?, Age=?, Gender=?, Email=?, Mobile=?, City=?, State=?, Country=? where id=?`)
	if err != nil {
		panic(err)
	}

	result, err := stmt.Exec(emp.Name, emp.Age, emp.Gender, emp.Email, emp.Mobile, emp.City, emp.State, emp.Country, id)
	if err != nil {
		panic(err)
	}
	count, _ := result.RowsAffected()
	return count > 0
}

func RemoveEmployee(id int) bool {

	stmt, err := getConnection().Prepare("delete from Employee where id=?")
	if err != nil {
		panic(err)
	}

	result, err := stmt.Exec(id)
	if err != nil {
		panic(err)
	}
	count, _ := result.RowsAffected()
	return count > 0
}
