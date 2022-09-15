package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type UserInfo struct {
	UserID   string `json:"userId"`
	Password string `json:"password"`
}

func Login() {
	var userID, password string
	var option int64
	fmt.Println("Enter Your User ID: ")
	fmt.Scanln(&userID)

	fmt.Println("Enter Your Password: ")
	fmt.Scanln(&password)

	userInfo := UserInfo{UserID: userID, Password: password}

	f, err := os.Open("userData.json")
	if err != nil {
		fmt.Println("Error while opening and reading the file")
	}

	defer f.Close()

	var userData UserInfo
	err = json.NewDecoder(f).Decode(&userData)
	if err != nil {
		fmt.Println("Error while reading the file")
	}

	if userData.UserID == userInfo.UserID && userData.Password == userInfo.Password {
		fmt.Println("Login Successfully")
		fmt.Println("1. Withdraw money.")
		fmt.Println("2. Deposit money.")
		fmt.Println("3. Request balance.")
		fmt.Println("4. Quit the program.")
		fmt.Scanln(&option)
	} else {
		fmt.Println("Invalid Credentials")
		os.Exit(1)
	}

	switch option {
	case 1:
		fmt.Println("Withdraw money")
	case 2:
		fmt.Println("Deposit money")
	case 3:
		fmt.Println("Request balance")
	case 4:
		fmt.Println("Quit the program")
		os.Exit(1)
	}
}

func CreateNewAccount() {
	var user UserInfo

	fmt.Printf("Please enter your user id: ")
	fmt.Scanln(&user.UserID)

	fmt.Printf("Please enter your password: ")
	fmt.Scanln(&user.Password)
}

func main() {
	fmt.Println("Hi! Welcome to Mr. Subham ATM Machine!")

	fmt.Println("Please select an option from the menu below:")
	fmt.Println("l -> Login")
	fmt.Println("c -> Create New Account")
	fmt.Println("q -> Quit")
	var choice string
	fmt.Scanln(&choice)

	switch choice {
	case "l":
		Login()
	case "c":
		CreateNewAccount()
	case "q":
		os.Exit(1)
	}

}
