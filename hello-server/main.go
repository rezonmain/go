package main

import (
	"fmt"
	"net/http"
)

func main() {
	staticDir := http.Dir("./public")
	fileServer := http.FileServer(staticDir)
	http.Handle("/", fileServer)
	fmt.Println("Server is running on port 3000")
	http.ListenAndServe(":3000", nil)
}
