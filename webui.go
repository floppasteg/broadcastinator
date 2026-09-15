package main

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func loadInterface() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("ui"))
	mux.Handle("/", fs)
	chk(http.ListenAndServe(":8080", mux))
}

func loadAPI() {
	route := gin.Default()
	route.Use(cors.New(cors.Config{
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"*"},
		AllowOrigins: []string{"*"},
	}))
	route.POST("/masterctrl/start", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"status":    "success",
			"data":      "",
			"timestamp": time.Now().UnixMilli(),
		})
	})
	route.POST("/masterctrl/shutdown-soft", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"status":    "success",
			"data":      "",
			"timestamp": time.Now().UnixMilli(),
		})
	})
	route.POST("/masterctrl/shutdown-hard", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"status":    "success",
			"data":      "",
			"timestamp": time.Now().UnixMilli(),
		})
	})
	route.POST("/masterctrl/e-stop", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"status":    "success",
			"data":      "",
			"timestamp": time.Now().UnixMilli(),
		})
	})
	route.Run(":8088")
}
