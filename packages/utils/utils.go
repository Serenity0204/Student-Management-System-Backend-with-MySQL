package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(req *http.Request, myInterface interface{}) {
	body, err := ioutil.ReadAll(req.Body)
	if err == nil {
		err := json.Unmarshal([]byte(body), myInterface)
		if err != nil {
			return
		}
	}
}
