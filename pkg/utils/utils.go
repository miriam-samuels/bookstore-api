package utils

import(
	"encoding/json"
	"net/http"
	"io/ioutil"
)
// unmarshall data (request) received from db from json
func ParseBody(r *http.Request, x interface{}){
	// ioutil.ReadAll to read the body of request
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		// unmarxhall body
		if err := json.Unmarshal([]byte(body), x); err != nil{
			return
		}
	}
} 