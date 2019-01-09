package main

import (
	"encoding/json"
	"net/http"
	// "strings"
	"time"
	"database/sql"
	"github.com/gin-gonic/gin"
	"fmt"
)
import _ "github.com/go-sql-driver/mysql"

//Page
type Page struct {
	Title   string
	Link    string
	Snippet string
}

//Query
type Query struct {
	Items []Page
}

var myClient = &http.Client{Timeout: 10 * time.Second}
var userQuery = "Perritos"

func getJSON(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}

func home(c *gin.Context) {
	c.JSON(200, gin.H{
		"message":"Hello World",
	})
}

func post(c *gin.Context) {
	c.JSON(200, gin.H{
		"message":"Post",
	})
}

func main() {
	// var apiURL = "https://www.googleapis.com/customsearch/v1?q=QUERY&cx=006924283690115384884%3Aci07khskaey&key=AIzaSyDnjzE_wsZ7Bo2KekvjnDvTgZFFLkezhT4"
	// r := gin.Default()
	// r.GET("/:query", func(c *gin.Context) {
	// 	query := new(Query) // or &Foo{}
	// 	getJSON(strings.Replace(apiURL, "QUERY", c.Param("query"), 1), query)
	// 	c.JSON(200, query)
	// })
	// r.GET("/", home)
	// r.POST("/p", post)
	// r.Run() // listen and serve on 0.0.0.0:8080apiURL = strings.Replace(apiURL, "QUERY", userQuery, 1)

	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/dist")

	// insert
	stmt, err := db.Prepare("INSERT busqueda SET id_busqueda=?,palabra=?")
	if err != nil {
		panic(err.Error())
	}

	res, err := stmt.Exec("1", "ajax")
	if err != nil {
		panic(err.Error())
	}

	id, err := res.LastInsertId()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(id)

	// update
	stmt, err = db.Prepare("update busqueda set id_busqueda=? where id_busqueda=?")
	if err != nil {
		panic(err.Error())
	}

	res, err = stmt.Exec("2", id)
	if err != nil {
		panic(err.Error())
	}

	affect, err := res.RowsAffected()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(affect)

	// query
	rows, err := db.Query("SELECT * FROM busqueda")
	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		var id int
		var palabra string
		err = rows.Scan(&id, &palabra)
		if err != nil {
		panic(err.Error())
	}
		fmt.Println(id)
		fmt.Println(palabra)
	}

	// delete
	stmt, err = db.Prepare("delete from busqueda where id_busqueda=?")
	if err != nil {
		panic(err.Error())
	}

	res, err = stmt.Exec(id)
	if err != nil {
		panic(err.Error())
	}

	affect, err = res.RowsAffected()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(affect)

	defer db.Close()


	fmt.Println("End Connection")
}