package main

import (
	"fmt"

	"github.com/kajanm/learningGo/models"
	"github.com/kajanm/learningGo/theory"
)

func main() {
	u := models.User{
		Id:        2,
		FirstName: "Panda",
		LastName:  "Karady",
	}
	fmt.Println(u)

	port2, err := theory.StartWebServer(3000)
	fmt.Println(port2, err)
}
