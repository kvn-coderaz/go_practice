package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func parsebody (r *http.Request, x interface{}){
	if body, err:= ioutil.ReadAll(r.Body); err==nil{
		if err:= json.Unmarshal(body, x); err!=nil{
			return
		}
	}
}