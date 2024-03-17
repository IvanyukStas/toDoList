package main

import(
	"fmt"
	"tdl/src/todolist"
)

func main(){
	fmt.Println("hello")

	u := todolist.NewUser("Tom")
	
	fmt.Println(u.GetUser())
}