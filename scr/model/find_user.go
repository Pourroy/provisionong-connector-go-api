package model

import (
	"github.com/Pourroy/provisionong-connector-go-api/scr/configuration/logger"
	"github.com/Pourroy/provisionong-connector-go-api/scr/configuration/rest_err"
)

func (ud *UserDomain) FindUser(string) (*UserDomain, *rest_err.RestErr) {
	logger.Info("Init FindUser Model")

	return nil, nil
}
