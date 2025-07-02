package main

import (
	"fmt"
	"os/user"
	"time"
)

func main() {
	fmt.Println("Hello World!")
	currentUser, err := user.Current()
	if err != nil {
		fmt.Println("Error getting current user:", err)
		return
	}
	fmt.Println(currentUser.Name)
	t := time.Now()
	fmt.Println(t.String())
	fmt.Scanln()
}
