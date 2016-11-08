package main

import "fmt"

var currentID int

var todos Todos

func init() {

}

func RepoFindTod(ID int) Todo {
	for _, t := range todos {
		if t.ID == ID {
			return t
		}
	}
	return Todo{}
}

func RepoCreateTodo(t Todo) Todo {
	fmt.Println("RepoCreateTodo start...")
	currentID += 1
	t.ID = currentID
	todos = append(todos, t)
	return t
}

func RepoDestroyTodo(ID int) error {
	for i, t := range todos {
		if t.ID == ID {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not finnd Todo with ID of %d to delete", ID)
}
