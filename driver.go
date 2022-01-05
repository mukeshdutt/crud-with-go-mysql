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

func AllEmployees() []domain.Employee {
	var emps []domain.Employee

	result, err := getConnection().Query("select * from employee")
	if err != nil {
		panic(err)
	}
	for result.Next() {
		emp := domain.Employee{}
		err = result.Scan(&emp.ID, &emp.Name, &emp.Age, &emp.Gender, &emp.Email, &emp.Mobile, &emp.City, &emp.State)
		if err != nil {
			panic(err)
		}
		emps = append(emps, emp)
	}
	return emps
}

func EmployeeInfoByID(id int) domain.Employee {

	result := getConnection().QueryRow("select * from employee where id=$1", id)

	if result.Err() != nil {
		panic(result.Err().Error())
	}
	emp := domain.Employee{}
	err := result.Scan(&emp.ID, &emp.Name, &emp.Age, &emp.Gender, &emp.Email, &emp.Mobile, &emp.City, &emp.State)
	if err != nil {
		panic(err)
	}
	return emp
}

func AddEmployee(emp domain.Employee) bool {

	query := `insert into Employee(Name, Age, Gender, Email, Mobile, City, State) 
	values('$1', $2, '$3', '$4', '$5', '$6', '$7');
	`
	result, err := getConnection().Exec(query, emp.Name, emp.Age, emp.Gender, emp.Email, emp.Mobile, emp.City, emp.State)

	if err != nil {
		panic(err)
	}
	count, _ := result.RowsAffected()
	return count > 0
}

func EditEmployee(id int, emp domain.Employee) bool {
	query := "update Employee set Name=$2, Age=$3, Gender=$4, Email=$5, Mobile=$6, City=$7, State=$8 where id=$1"
	result, err := getConnection().Exec(query, id, emp.Name, emp.Age, emp.Gender, emp.Email, emp.Mobile, emp.City, emp.State)

	if err != nil {
		panic(err)
	}
	count, _ := result.RowsAffected()
	return count > 0
}

func RemoveEmployee(id int) bool {
	query := "delete from employee where id=$1"
	result, err := getConnection().Exec(query, id)

	if err != nil {
		panic(err)
	}
	count, _ := result.RowsAffected()
	return count > 0
}
