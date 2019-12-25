package main

import (
	"fmt"

	"github.com/kajanm/learningGo/models"
)

func main() {
	u := models.User{
		FirstName: "Panda",
		LastName:  "Karady",
	}
	fmt.Println(u)
	u2, err := models.AddUser(u)
	fmt.Println(u2, err)
}
