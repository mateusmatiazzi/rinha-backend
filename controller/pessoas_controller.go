package controller

import (
	"github.com/gin-gonic/gin"
)

func PessoasController() {
	gin := gin.Default()

	gin.GET("/ping", nil)
	gin.Run(":8080")
}
