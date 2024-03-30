package model

import (
	"github.com/Pourroy/provisionong-connector-go-api/scr/configuration/logger"
	"github.com/Pourroy/provisionong-connector-go-api/scr/configuration/rest_err"
)

func (*UserDomain) CreateUser(UserDomain) *rest_err.RestErr {
	logger.Info("Init CreateUser Model")
	return nil
}