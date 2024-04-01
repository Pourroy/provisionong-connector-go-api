package controller

import (
	"github.com/Pourroy/provisionong-connector-go-api/scr/configuration/logger"
	"github.com/Pourroy/provisionong-connector-go-api/scr/configuration/validation"
	"github.com/Pourroy/provisionong-connector-go-api/scr/controller/model/request"
	"github.com/Pourroy/provisionong-connector-go-api/scr/controller/model/response"
	"github.com/Pourroy/provisionong-connector-go-api/scr/model"
	//"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func CreateUser(reqContext *gin.Context) {
	logger.Info("Init CreateUser controller")
	var userRequest request.UserRequest

	if err := reqContext.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user information", err)
		errRest := validation.ValidateUserError(err)

		reqContext.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	if err := domain.CreateUser(); err != nil {
		reqContext.JSON(err.Code, err)
		return
	}

	response := response.NewUserResponse(userRequest)
	logger.Info("User created successfully")
	reqContext.JSON(202, response)
}
