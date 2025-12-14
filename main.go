package main

import (
	"fmt"
	"taskmanager/data"
	"taskmanager/router"
)

func main() {
	err := data.ConnectWithDB()
	if err != nil {
		fmt.Println("Error")
	}
	r := router.RouterSetup()
	r.Run()
}
