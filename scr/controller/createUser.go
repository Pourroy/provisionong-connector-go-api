package controller

import (
	"github.com/Pourroy/provisionong-connector-go-api/scr/configuration/logger"
	"github.com/Pourroy/provisionong-connector-go-api/scr/configuration/validation"
	"github.com/Pourroy/provisionong-connector-go-api/scr/controller/model/request"
	"github.com/Pourroy/provisionong-connector-go-api/scr/controller/model/response"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateUser(reqContext *gin.Context) {
	logger.Info("Init CreateUser controller", zap.String("transaction", requestid.Get(reqContext)))
	var userRequest request.UserRequest

	if err := reqContext.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user information", err, zap.String("transaction", requestid.Get(reqContext)))
		errRest := validation.ValidateUserError(err)

		reqContext.JSON(errRest.Code, errRest)
		return
	}

	response := response.NewUserResponse(userRequest)
	logger.Info("User created successfully", zap.String("transactionId", requestid.Get(reqContext)))
	reqContext.JSON(202, response)
}
