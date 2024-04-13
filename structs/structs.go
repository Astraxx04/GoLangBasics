package main

import(
	"fmt"
	"structs/user"
)

func getUserData(promptText string)(string) {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}


func main() {
	userFirstName := getUserData("Enter your first name: ")
	userLastName := getUserData("Enter your last name: ")
	userBirthDate := getUserData("Enter your birth date (MM/DD/YYYY): ")

	// var appUser *user.User

	appUser, err := user.New(userFirstName, userLastName, userBirthDate)
	if(err != nil) {
		fmt.Println(err)
		return
	}
	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()

	admin, err := user.NewAdmin("admin@gmail.com", "admin123")
	admin.OutputUserDetails()
	admin.ClearUserName()
	admin.OutputUserDetails()

	if(err != nil) {
		fmt.Println(err)
		return
	}
}



// Scanln stops waiting for a value when you pres enter
// The fields inside the struct should also be in caps to export