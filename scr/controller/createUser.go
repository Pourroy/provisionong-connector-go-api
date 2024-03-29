package controller

import (
	"fmt"
	"log"

	"github.com/Pourroy/provisionong-connector-go-api/scr/configuration/validation"
	"github.com/Pourroy/provisionong-connector-go-api/scr/controller/model/request"
	"github.com/Pourroy/provisionong-connector-go-api/scr/controller/model/response"
	"github.com/gin-gonic/gin"
)

func CreateUser(reqContext *gin.Context) {
	log.Println("Init CreateUser controller")
	var userRequest request.UserRequest

	if err := reqContext.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("Error trying to marshal object error=%s\n", err.Error())
		errRest := validation.ValidateUserError(err)

		reqContext.JSON(errRest.Code, errRest)
		return
	}

	fmt.Println(userRequest)

	response := response.NewUserResponse(userRequest)
	reqContext.JSON(202, response)
}
