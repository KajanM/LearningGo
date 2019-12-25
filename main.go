package main

import (
	"fmt"

	"github.com/kajanm/learningGo/models"
)

func main() {
	u := models.User{
		Id:        2,
		FirstName: "Panda",
		LastName:  "Karady",
	}
	fmt.Println(u)
}
