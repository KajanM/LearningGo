package theory

import (
	"fmt"
)

func StartWebServer(port int) (int, error) {
	fmt.Println("Starting server...")
	fmt.Println("Server started on port", port)
	//return nil
	return port, nil
}
