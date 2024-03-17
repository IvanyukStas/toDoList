package main

import(
	"fmt"
	"tdl/src/todolist"
	"tdl/src/utils"
)

func main(){
	fmt.Println("hello")

	u := todolist.NewUser("Tom")
	fmt.Println("Tite: ")
	title := utils.GetStringFromCLI()
	fmt.Println("description: ")
	description := utils.GetStringFromCLI()
	t := todolist.NewTask(title, description)
	
	
	
	fmt.Println(u.GetUser())
	fmt.Println(*t)
}