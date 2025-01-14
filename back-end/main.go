package main

import (
	"fmt"
	"net/http"

	"gitgub.com/fuad7161/routes"
)

func main() {
	routes.Routes()
	fmt.Println("Server is running at :8080")
	http.ListenAndServe(":8080", nil)
}
