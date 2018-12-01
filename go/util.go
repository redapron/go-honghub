package swagger

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	MessageCode        string      `json:"messageCode,omitempty"`
	MessageDescription string      `json:"messageDescription,omitempty"`
	Data               interface{} `json:"data,omitempty"`
}

func JSONResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(Response{
			MessageCode:        "00",
			MessageDescription: "StatusOK",
			Data:               data,
		})
	} else {
		json.NewEncoder(w).Encode(Response{
			MessageCode:        "00",
			MessageDescription: "StatusOK",
		})
	}
	return
}
