package controller

import (
	"github.com/Pourroy/provisionong-connector-go-api/scr/configuration/rest_err"
	"github.com/gin-gonic/gin"
)

func CreateUser(reqContext *gin.Context) {
	err := rest_err.NewInternalServerError("Você chamou a rota de forma errada papai")
	reqContext.JSON(err.Code, err)
}