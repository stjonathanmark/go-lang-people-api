package web

import (
	"fmt"
	"net/http"
	"strconv"
)

func HandleError(w http.ResponseWriter, err error, message string, httpStatus int) (hasError bool) {
	if err != nil {
		w.WriteHeader(httpStatus)
		fmt.Fprint(w, message, err)
		hasError = true
	}
	return
}

func Offset(pageNumber int, pageSize int) (offset int) {
	if pageNumber > 1 {
		offset = (pageNumber - 1) * pageSize
	}
	return
}

func getIntQueryParam(r *http.Request, name string) (value int) {
	queryVal, err := strconv.ParseInt(r.URL.Query().Get(name), 0, 32)
	if err == nil {
		value = int(queryVal)
	}
	return
}

type Error struct {
	Message string
}

func (e Error) Error() string {
	return e.Message
}
