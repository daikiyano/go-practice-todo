package main

import "go-todo-practice/app/models"

func main() {
	//fmt.Println(config.Config.Port)
	//fmt.Println(config.Config.SQLDriver)
	//fmt.Println(config.Config.DbName)
	//fmt.Println(config.Config.LogFile)
	//log.Println("test")
	//fmt.Println(models.Db)
	//
	//u := &models.User{}
	//u.Name = "daiki"
	//u.Email = "daiki@gmail.com"
	//u.PassWord = "testtest"
	//fmt.Println(u)
	//u.CreateUser()
	//u, _ := models.GetUser(1)
	//
	//u.Name = "Test2"
	//u.Email = "test2@example.com"
	//u.UpdateUser()
	//u.DeleteUser()
	//u, _ = models.GetUser(1)
	//fmt.Println(u)
	//user, _ := models.GetUser(3)
	//user.CreateTodo("third Todo")
	//t, _ := models.GetTodo(3)
	//fmt.Println(t)
	//
	//user, _ := models.GetUser(2)
	//user.CreateTodo("Second Todo")
	//todos, _ := models.GetTodos()
	//for _, v := range todos {
	//	fmt.Println(v)
	//}

	//user2, _ := models.GetUser(2)
	//todos, _ := user2.GetTodosByUser()
	//for _, v := range todos {
	//	fmt.Println(v)
	//}

	t, _ := models.GetTodo(2)
	t.Content = "update todo"
	t.UpdateTodo()
}
