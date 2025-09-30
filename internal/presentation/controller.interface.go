package controllers

type ControllerInterface interface {
	Handle(request interface{}) (int, interface{})
}
