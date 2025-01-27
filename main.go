package main

import (
	"gin-orm/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.RestfulUser(r)

	//r.Run(":9999")
	err := r.Run(":9999")
	if err != nil {
		panic(err)
	}

}
