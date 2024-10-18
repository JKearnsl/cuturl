package common

type Interactor interface {
	Execute(input *interface{}) (*interface{}, error)
}
