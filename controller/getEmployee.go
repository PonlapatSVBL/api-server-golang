package employeecontroller

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/svbl/golang-api/initializers"
	"github.com/svbl/golang-api/models"
)

type EmployeeStruct models.EmployeeStruct

func GetEmployee() ([]EmployeeStruct, error) {
	query := "SELECT employee_id, employee_name FROM hms_api.comp_employee LIMIT 5"

	rows, err := initializers.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []EmployeeStruct

	for rows.Next() {
		var emp EmployeeStruct
		if err := rows.Scan(&emp.EmployeeId, &emp.EmployeeName); err != nil {
			log.Println(err)
			continue
		}
		employees = append(employees, emp)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	fmt.Println("employees", employees)
	for _, emp := range employees {
		fmt.Printf("employee_id: %s, employee_name: %s\n", emp.EmployeeId, emp.EmployeeName)
	}
	fmt.Println()

	return employees, nil
}

func GetEmployee2() ([]EmployeeStruct, error) {
	query := "SELECT employee_id, employee_name FROM hms_api.comp_employee LIMIT 5"

	rows, err := initializers.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []EmployeeStruct

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	columnCount := len(columns)

	for rows.Next() {
		var emp EmployeeStruct
		values := make([]interface{}, columnCount)
		valuePtrs := make([]interface{}, columnCount)

		for i := range values {
			valuePtrs[i] = &values[i]
		}

		if err := rows.Scan(valuePtrs...); err != nil {
			log.Println(err)
			continue
		}

		elem := reflect.ValueOf(&emp).Elem()
		for i, colName := range columns {
			fieldName := ToCamel(colName)
			field := elem.FieldByName(fieldName)
			if field.IsValid() {
				val := values[i]
				strVal := fmt.Sprintf("%v", val)
				field.SetString(strVal)
			}
		}

		employees = append(employees, emp)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	fmt.Println("employees", employees)
	for _, emp := range employees {
		fmt.Printf("employee_id: %s, employee_name: %s\n", emp.EmployeeId, emp.EmployeeName)
	}
	fmt.Println()

	return employees, nil
}

func GetEmployee3() ([]EmployeeStruct, error) {
	query := "SELECT * FROM hms_api.comp_employee LIMIT 2"

	rows, err := initializers.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []EmployeeStruct

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	columnCount := len(columns)

	for rows.Next() {
		var emp EmployeeStruct
		values := make([]interface{}, columnCount)
		valuePtrs := make([]interface{}, columnCount)

		for i := range values {
			valuePtrs[i] = &values[i]
		}

		if err := rows.Scan(valuePtrs...); err != nil {
			log.Println(err)
			continue
		}

		elem := reflect.ValueOf(&emp).Elem()
		for i, colName := range columns {
			fieldName := ToCamel(colName)
			field := elem.FieldByName(fieldName)
			if field.IsValid() {
				val := values[i]
				switch v := val.(type) {
				case []byte:
					field.SetString(string(v))
				case string:
					field.SetString(v)
				default:
					field.SetString(fmt.Sprintf("%v", v))
				}
			}
		}

		employees = append(employees, emp)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	// fmt.Println("employees", employees)
	fmt.Println()
	/* for _, emp := range employees {
		fmt.Printf("employee_id: %s, employee_name: %s\n", emp.EmployeeId, emp.EmployeeName)
	} */
	for _, emp := range employees {
		v := reflect.ValueOf(emp)
		typeOfS := v.Type()

		for i := 0; i < v.NumField(); i++ {
			fmt.Printf("%s: %v\n", typeOfS.Field(i).Name, v.Field(i).Interface())
		}
		fmt.Println()
	}
	fmt.Println()

	return employees, nil
}

func GetEmployee4() ([]EmployeeStruct, error) {
	query := "SELECT employee_id, employee_name FROM hms_api.comp_employee LIMIT 2"

	rows, err := initializers.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []EmployeeStruct

	columns, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}
	columnCount := len(columns)

	for rows.Next() {
		values := make([]interface{}, columnCount)
		valuePtrs := make([]interface{}, columnCount)

		for i := range values {
			valuePtrs[i] = &values[i]
		}

		if err := rows.Scan(valuePtrs...); err != nil {
			log.Println(err)
			continue
		}

		var emp EmployeeStruct
		elem := reflect.ValueOf(&emp).Elem()

		for i, col := range columns {
			fieldName := ToCamel(col)
			field := elem.FieldByName(fieldName)
			if field.IsValid() {
				switch v := values[i].(type) {
				case []byte:
					field.SetString(string(v))
				case string:
					field.SetString(v)
				default:
					field.SetString(fmt.Sprintf("%v", v))
				}
			}
		}

		employees = append(employees, emp)
	}

	fmt.Println()
	/* for _, emp := range employees {
		v := reflect.ValueOf(emp)
		typeOfS := v.Type()

		for i := 0; i < v.NumField(); i++ {
			fieldValue := v.Field(i).Interface()
			if fieldValue != nil && fieldValue != "" {
				fmt.Printf("%s: %v\n", typeOfS.Field(i).Name, fieldValue)
			}
		}
		fmt.Println()
	} */
	/* jsonBytes, _ := json.Marshal(employees)
	fmt.Println(string(jsonBytes)) */
	for _, emp := range employees {
		jsonBytes, _ := json.Marshal(emp)
		fmt.Println(string(jsonBytes))
		fmt.Println("< ========================================================== >")
	}
	fmt.Println()

	return employees, nil
}

func ToCamel(s string) string {
	parts := strings.Split(s, "_")
	for i := range parts {
		parts[i] = strings.Title(parts[i])
	}
	return strings.Join(parts, "")
}
