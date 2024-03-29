package controller

import (
	"fmt"

	"github.com/Pourroy/provisionong-connector-go-api/scr/configuration/rest_err"
	"github.com/Pourroy/provisionong-connector-go-api/scr/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(reqContext *gin.Context) {
	var userRequest request.UserRequest

	if err := reqContext.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("There are some incorret fields, error=%s", err.Error()))

		reqContext.JSON(restErr.Code, restErr)
	}
}
