package main

import (
	"fmt"
	"go-todo-practice/app/controllers"
	"go-todo-practice/app/models"
	"log"
)

func main() {
	fmt.Println(models.Db)
	controllers.StartMainServer()
	user, _ := models.GetUserByEmail("test@example.com")
	fmt.Println(user)

	session, err := user.CreateSession()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(session)

	valid, _ := session.CheckSession()
	fmt.Println(valid)
}
