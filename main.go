package main

import "github.com/Komei1009/todoList/router"

func main() {
	r := router.Getrouter()
	r.Run()
}