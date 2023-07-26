package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("Helloooooooo~")

	l := logrus.New()
	l.Info("log init")

	r := gin.Default()

	r.GET("/info", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": 1, "message": "Server is running!"})
	})

	r.Run(":999")
}
