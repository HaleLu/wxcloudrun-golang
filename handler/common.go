package handler

import (
	"encoding/json"
	"net/http"
)

func GetOpenId(r *http.Request) string {
	return r.Header["X-Wx-Openid"][0]
}

// getAction 获取action
func getReq[T any](r *http.Request) (*T, error) {
	decoder := json.NewDecoder(r.Body)
	body := new(T)
	if err := decoder.Decode(body); err != nil {
		return nil, err
	}
	defer r.Body.Close()
	return body, nil
}

// JsonResult 返回结构
type JsonResult struct {
	Code     int         `json:"code"`
	ErrorMsg string      `json:"errorMsg,omitempty"`
	Data     interface{} `json:"data"`
}

// todo: 鉴权以后加
var AdminOpenId = []string{}
