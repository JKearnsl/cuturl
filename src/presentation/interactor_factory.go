package presentation

import (
	"cuturl/src/application"
)

type InteractorFactory interface {
	MakeUrl() application.MakeUrl
	GetUrl() application.GetUrl
}
