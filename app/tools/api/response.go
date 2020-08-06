package api

import (
	"encoding/json"

	"github.com/vgmdj/utils/logger"
)

type Response struct {
	Code Code        `json:"Code"`
	Msg  string      `json:"Msg"`
	Data interface{} `json:"Data,omitempty"`
}

func NewResponse(code Code, data interface{}, msgs ...string) Response {
	var msg = code.Msg()
	if len(msgs) != 0 {
		msg = msgs[0]
	}

	if code != Success {
		logger.Info(code, code.Msg())
	}

	return Response{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

func (resp *Response) JsonMarshal() ([]byte, error) {
	bts, err := json.Marshal(resp)
	if err != nil {
		logger.Error(err.Error())
		return nil, nil
	}

	return bts, nil

}
