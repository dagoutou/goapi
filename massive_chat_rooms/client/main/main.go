package main

import (
	"fmt"
	"goapi/massive_chat_rooms/client"
	"os"
)

var (
	userId  int
	userPwd string
)

func main() {
	var key int
	var loop = true
	for loop {
		fmt.Println("----------Welcome To Massive Chat Room!-------------")
		fmt.Println("\t\t\t 1 Sign In")
		fmt.Println("\t\t\t 2 Register User")
		fmt.Println("\t\t\t 3 Sign Out")
		fmt.Println("\t\t\t Please Change(1-3):")
		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("Sign In rooms")
			loop = false
		case 2:
			fmt.Println("Register user")
			loop = false
		case 3:
			fmt.Println("Sign out")
			os.Exit(0)
		default:
			fmt.Println("Your input is incorrect, please re-enter")
		}
	}
	if key == 1 {
		fmt.Println("Please input user id")
		fmt.Scanf("%d", userId)
		fmt.Println("Please input user password")
		fmt.Scanf("%d", userPwd)
		err := client.Login(userId, userPwd)
		if err != nil {
			fmt.Println("Login success!")
		} else {
			fmt.Println("Login failed")
		}
	} else if key == 2 {
		//进行注册
		fmt.Println()
	}

}
