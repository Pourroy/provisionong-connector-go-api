package model

import (
	"fmt"
	"github.com/Pourroy/provisionong-connector-go-api/scr/configuration/logger"
	"github.com/Pourroy/provisionong-connector-go-api/scr/configuration/rest_err"
)

func (ud *UserDomain) CreateUser() *rest_err.RestErr {
	logger.Info("Init CreateUser Model")
	ud.EncryptPassword()

	fmt.Println(ud)
	return nil
}
