package controllers

import (
	"fmt"
)

func HttpOk(body interface{}) (int, interface{}) {
	return 200, body
}

func HttpOkNoContent() (int, interface{}) {
	return 204, nil
}

func HttpBadRequest(message string) (int, map[string]interface{}) {
	return 400, map[string]interface{}{"message": message}
}

func HttpUnauthorized(message string) (int, map[string]interface{}) {
	return 401, map[string]interface{}{"message": message}
}

func HttpNotFound(message string) (int, map[string]interface{}) {
	return 404, map[string]interface{}{"message": message}
}

func HttpUnprocessableEntity(message string) (int, map[string]interface{}) {
	return 422, map[string]interface{}{"message": message}
}

func HttpInternalError(err error) (int, map[string]interface{}) {
	fmt.Printf("internal server error: %v", err)
	return 500, map[string]interface{}{"message": "internal server error"}
}
