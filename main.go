package main

import (
	"fmt"
	"go-todo-practice/app/models"
)

func main() {
	fmt.Println(models.Db)
	//controllers.StartMainServer()
	user, _ := models.GetUserByEmail("test@example.com")
	fmt.Println(user)
}
