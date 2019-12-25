package theory

import (
	"errors"
	"fmt"
)

func StartWebServer(port int) error {
	fmt.Println("Starting server...")
	fmt.Println("Server started on port", port)
	//return nil
	return errors.New("something went wrong")
}
