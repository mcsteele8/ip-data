package werror

import (
	"encoding/json"
	"fmt"
	"ip-data/tools/wlog"
	"net/http"
)

type Error struct {
	Message string `json:"message"`
}

func (e *Error) ReturnHttpError(w http.ResponseWriter, statusCode int) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(&e)
}

func DoHttpError(w http.ResponseWriter, statusCode int, message string) {
	wlog.New().Errorf("Http error code: %d, with message: %s", statusCode, message)
	err := Error{Message: message}
	err.ReturnHttpError(w, statusCode)
}

func Wrap(err error, errMessage string) error {
	return fmt.Errorf(errMessage+" | err: %w", err)
}
