package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// MySQL database connection string
	db, err := sql.Open("mysql", "root:root12345@tcp(localhost:3306)/studidevsecops")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Create a new Gin router
	router := gin.Default()

	// Define a route for "/hello"
	router.GET("/hello", func(c *gin.Context) {
		greeting := c.Query("greeting")
		name := c.Query("name")

		insertData(db, greeting, name, time.Now().Unix())
		c.Data(200, "text/plain", []byte(greeting+" "+name+
			"!, Your data is saved in database"))
	})

	// Start the server on port 8080
	router.Run(":8080")
}

func insertData(db *sql.DB, greeting, name string, createdAt int64) {
	// Prepare the SQL statement with placeholders for values
	stmt, err := db.Prepare("INSERT INTO greeting (greeting, name, created_at) VALUES (?, ?, ?)")
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()

	// Execute the prepared statement with the actual values
	_, err = stmt.Exec(greeting, name, createdAt)
	if err != nil {
		fmt.Println(err)
	}
}
