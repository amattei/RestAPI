package main

import "github.com/undoRep/clientServiceAPI"

func main() {
	client := myservice.NewBasicAuthClient("admin", "94295957-94ee-4085-8345-5eb14fdd6c4c")
	todo := Todo{
		Content: "New Todo",
		Done:    true,
	}
	// Add a todo
	_ := client.AddTodo(&todo)

	// Fetch a todo
	t, _ := client.GetTodo(1)
}
