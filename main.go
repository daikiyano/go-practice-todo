package main

import (
	"fmt"
	"go-todo-practice/app/controllers"
	"go-todo-practice/app/models"
)

func main() {
	fmt.Println(models.Db)
	controllers.StartMainServer()
}
