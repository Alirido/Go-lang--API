package main
 
import (
    "net/http"
    "database/sql"
    "github.com/labstack/echo"
    "fmt"
    "log"
    _"github.com/lib/pq"
)

func main() {
	// var err 
	db, err := sql.Open("postgres", "user=postgres password=root dbname=books_database sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		panic(err)
	} else {
		fmt.Println("DB Connected...")
	}

	e := echo.New()

	type Employee struct {
		Id     string `json:"id"`
		Name   string `json:"name"`
		Salary string `json: "salary"`
		Age    string `json : "age"`
	}
	type Employees struct {
		Employees []Employee `json:"employees"`
	}

	e.POST("/employee", func(c echo.Context) error {
		u := new(Employee)
		if err := c.Bind(u); err != nil {
			return err
		}
		sqlStatement := "INSERT INTO employees (name, salry,age)VALUES ($1, $2, $3)"
		res, err := db.Query(sqlStatement, u.Name, u.Salary, u.Age)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
			return c.JSON(http.StatusCreated, u)
		}
		return c.String(http.StatusOK, "ok")
	})

	e.PUT("/employee", func(c echo.Context) error {
		u := new(Employee)
		if err := c.Bind(u); err != nil {
			return err
		}
		sqlStatement := "UPDATE employees SET name=$1,salary=$2,age=$3 WHERE id=$5"
		res, err := db.Query(sqlStatement, u.Name, u.Salary, u.Age, u.Id)
		if err != nil {
			fmt.Println(err)
			//return c.JSON(http.StatusCreated, u);
		} else {
			fmt.Println(res)
			return c.JSON(http.StatusCreated, u)
		}
		return c.String(http.StatusOK, u.Id)
	})

	e.DELETE("/employee/:id", func(c echo.Context) error {
		id := c.Param("id")
		sqlStatement := "DELETE FROM employees WHERE id = $1"
		res, err := db.Query(sqlStatement, id)
		if err != nil {
			fmt.Println(err)
			//return c.JSON(http.StatusCreated, u);
		} else {
			fmt.Println(res)
			return c.JSON(http.StatusOK, "Deleted")
		}
		return c.String(http.StatusOK, id+"Deleted")
	})

	e.GET("/employee", func(c echo.Context) error {
		sqlStatement := "SELECT id, name, salary, age FROM employees order by id"
		rows, err := db.Query(sqlStatement)
		if err != nil {
			fmt.Println(err)
			// return c.JSON(http.StatusCreated, u)
			return c.JSON(http.StatusCreated, "Error disini")
		}
		defer rows.Close()
		result := Employees{}
	
		for rows.Next() {
			employee := Employee{}
			err2 := rows.Scan(&employee.Id, &employee.Name, &employee.Salary, &employee.Age)
			// Exit if we get an error
			if err2 != nil {
				return err2
			}
			result.Employees = append(result.Employees, employee)
		}
		return c.JSON(http.StatusCreated, result)
	
		// return c.String(http.StatusOK, "ok")
	})

	e.Logger.Fatal(e.Start(":1323"))
}