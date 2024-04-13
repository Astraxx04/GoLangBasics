package user
import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName string
	birthDate string
	createdAt time.Time
}

type Admin struct{
	email string
	password string
	User // you can use anonymous embedding wo specifing a name
}

func (usr *User) OutputUserDetails() { //Receiver arguement attached to a struct
	fmt.Println(usr.firstName, usr.lastName, usr.birthDate, usr.createdAt)
}

func (usr *User) ClearUserName() {
	usr.firstName = ""
	usr.lastName = ""
}

func NewAdmin(email, password string) (*Admin, error) {
	if(email == "" || password == "") {
		return nil, errors.New("all the values are required")
	}
	return &Admin{
		email: email,
		password: password,
		User: User{
			firstName: "Admin",
			lastName: "Admin",
			birthDate: "--/--/----",
			createdAt: time.Now(),
		},
	}, nil
}

func New(firstName, lastName, birthDate string) (*User, error) {
	if(firstName == "" || lastName == "" || birthDate == "") {
		return nil, errors.New("all the values are required")
	}
	return &User{
		firstName: firstName,
		lastName: lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}

//You can embed a struct inside another struct