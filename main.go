package main

import (
	"eLeMeS/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.Student(r)

	if err := r.Run(); err != nil {
		return
	}
}
