package responses

import (
	"encoding/json"
)

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

func (r Response) toString() string {
	v, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}

	return string(v)
}
