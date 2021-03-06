package task

import (
	"fmt"
	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/service"
	"github.com/irisnet/explorer/backend/utils"
)

type UpdateValidator struct{}

func (task UpdateValidator) Name() string {
	return "update_validator"
}
func (task UpdateValidator) Start() {
	utils.RunTimer(conf.Get().Server.CronTimeValidators, utils.Sec, func() {
		if err := task.DoTask(); err != nil {
			logger.Error(fmt.Sprintf("%s fail", task.Name()), logger.String("err", err.Error()))
		} else {
			logger.Info(fmt.Sprintf("%s success", task.Name()))
		}
	})

}

func (task UpdateValidator) DoTask() error {
	validators, err := document.Validator{}.GetAllValidator()

	if err != nil {
		return err
	}

	validatorService := service.Get(service.Validator).(*service.ValidatorService)
	err = validatorService.UpdateValidators(validators)

	if err != nil {
		return err
	}

	return nil
}
