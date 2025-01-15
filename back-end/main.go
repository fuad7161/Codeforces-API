package main

import (
	"fmt"
	"net/http"

	"gitgub.com/fuad7161/routes"
)

func main() {
	// tc, err := renders.CreateTemplateCache()
	// if err != nil {
	// 	log.Fatal("Cannot create template cache")
	// 	return
	// }

	routes.Routes()
	fmt.Println("Server is running at :8080")
	http.ListenAndServe(":8080", nil)
}
