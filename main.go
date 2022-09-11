package main

import (
	"fmt"
	"log"
	"net/http"
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router := gin.Default()
	router.POST("/adduser", adduser)
	router.GET("/showusers", showusers)
	router.PATCH("/updatepassword", updatepassword)

	err := router.Run("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/tasneem")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println(" error connectting db .Ping")
		panic(err.Error)
	}
	fmt.Println("successfully connected to db ")
}

type update struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type users struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Pnum     int    `json:"pnum"`
	Mail     string `json:"mail"`
}

func adduser(c *gin.Context) {
	var NewUser users
	err := c.BindJSON(&NewUser)
	if err != nil {
		fmt.Println("errrrrrrrrrrrrrrrrrrrrrrrrrr: ", err)
		c.JSON(400, err)
		return
	}
	fmt.Println(NewUser)
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/new_schema")
	if err != nil {
		fmt.Println("dberr", err)
		panic(err.Error())
	}

	defer db.Close()
	query := "insert into users (name,password,pnum,mail) values (?,?,?,?);"
	insert, _ := db.Prepare(query)
	_, err = insert.Exec(NewUser.Name, NewUser.Password, NewUser.Pnum, NewUser.Mail)
	if err != nil {
		fmt.Println("errrrrot:  ", err)
		c.JSON(500, err)
	}
	fmt.Println("connected ")
}

func showusers(c *gin.Context) {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/new_schema")
	if err != nil {
		fmt.Println("dberr", err)
		panic(err.Error())
	}

	defer db.Close()

	var User = users{}
	var AllUsers = []users{}
	rows, err := db.Query("select * from users")
	for rows.Next() {
		err = rows.Scan(&User.Name, &User.Pnum, &User.Mail, &User.Password)
		AllUsers = append(AllUsers, User)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println(AllUsers)
	c.JSON(http.StatusOK, AllUsers)
}

func updatepassword(c *gin.Context) {
	var updateusers update
	err := c.BindJSON(&updateusers)
	if err != nil {
		fmt.Println("errrrrrrrrrrrrrrrrrrrrrrrrrr: ", err)
		c.JSON(400, err)
		return
	}
	fmt.Println(updateusers)
	fmt.Println("1")
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/new_schema")
	if err != nil {
		fmt.Println("dberr", err)
		fmt.Println("2")
		panic(err.Error())
	}

	defer db.Close()
	fmt.Println("3")

	stmt, err := db.Prepare("update users set password = ? where name=?;")
	fmt.Println("err1: ", err)
	res, err := stmt.Exec(updateusers.Password, updateusers.Name)
	fmt.Println("err2: ", err)
	a, err := res.RowsAffected()
	fmt.Println("err3: ", err)
	fmt.Println(a)
}

//func  deleteacc (c*gin.Context){
//	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/new_schema")
//	if err != nil {
//		fmt.Println("dberr", err)
//		panic(err.Error())
//	}

//	stmt, e := db.Prepare("delete from users where name=?;")
//	//res, e := stmt.Exec("")
//	a, e := res.RowsAffected()
//	fmt.Println(a)
//}
