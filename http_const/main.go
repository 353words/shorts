package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println(http.MethodPost)   // POST
	fmt.Println(http.StatusTeapot) // 418
}
