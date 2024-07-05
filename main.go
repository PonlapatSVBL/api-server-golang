package main

import (
	_ "github.com/go-sql-driver/mysql"
	employeecontroller "github.com/svbl/golang-api/controller"
	initializers "github.com/svbl/golang-api/initializers"
	"github.com/svbl/golang-api/mockup"
)

func main() {
	initializers.LoadEnv()
	initializers.ConnectDb()

	// employeecontroller.GetEmployee3()
	employeecontroller.GetEmployee4()

	// employeecontroller.ListEmployee()

	mockup.GetBooks()
}
