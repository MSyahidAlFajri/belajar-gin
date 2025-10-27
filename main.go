package main

import (
	"github.com/MSyahidAlFajri/belajar-gin/src/config"
	route "github.com/MSyahidAlFajri/belajar-gin/src/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db := config.DB()

	route.Api(r, db)

	r.Run()
}
