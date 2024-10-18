package main

import (
	"cuturl/src/application"
	"cuturl/src/application/common"
)

type IoC struct {
	UnitGateway common.UnitGateway
}

func NewIoc(unitGateway common.UnitGateway) IoC {
	return IoC{
		UnitGateway: unitGateway,
	}
}

func (ioc *IoC) MakeUrl() application.MakeUrl {
	return application.MakeUrl{
		UnitGateway: ioc.UnitGateway,
	}
}

func (ioc *IoC) GetUrl() application.GetUrl {
	return application.GetUrl{
		UnitGateway: ioc.UnitGateway,
	}
}
