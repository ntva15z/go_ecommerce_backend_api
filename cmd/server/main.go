package main

import (
	"github.com/ntva15z/go-ecommerce-backend-api/internal/initialize"
)

func main() {
	initialize.Run() // listen and serve on 0.0.0.0:8888 (for windows "localhost:8888")
}
