package main

import (
	"vualt/api/router"
)

func main() {
	r := router.New()
	r.Run("localhost:9000")
}

