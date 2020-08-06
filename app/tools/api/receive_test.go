package api

import (
	"encoding/json"
	"testing"
)

func TestReceive(t *testing.T) {
	test := `{"Code":0,"Msg":"成功","Data":{"A":"a","B":1,"T":"2019-10-18T14:23:34.966294+08:00"}}`
	recv := ReceiveInterface{}
	err := json.Unmarshal([]byte(test), &recv)
	if err != nil {
		t.Error(err.Error())
		return
	}
	data := recv.Data.(map[string]interface{})
	t.Log(data)

	test2 := `{"Code":0,"Msg":"成功","Data":[{"A":"a","B":1,"T":"2019-10-18T14:23:00.401788+08:00"},{"A":"aa","B":11,"T":"2019-10-18T14:23:00.401788+08:00"}]}`
	recv2 := ReceiveInterface{}
	err = json.Unmarshal([]byte(test2), &recv2)
	if err != nil {
		t.Error(err.Error())
		return
	}
	data2 := recv2.Data.([]interface{})
	for k, v := range data2 {
		t.Log(k)
		t.Log(v.(map[string]interface{}))
	}
}

func TestReceive1(t *testing.T) {
	test := `{"Code":0,"Msg":"成功","Data":{"A":"a","B":1,"T":"2019-10-18T14:23:34.966294+08:00"}}`
	recv := Receive{}
	err := json.Unmarshal([]byte(test), &recv)
	if err != nil {
		t.Error(err.Error())
		return
	}

	t.Log(recv)
}

func TestReceive2(t *testing.T) {
	test := `{"Code":0,"Msg":"成功","Data":[{"A":"a","B":1,"T":"2019-10-18T14:23:00.401788+08:00"},{"A":"aa","B":11,"T":"2019-10-18T14:23:00.401788+08:00"}]}`
	recv := ReceiveArray{}
	err := json.Unmarshal([]byte(test), &recv)
	if err != nil {
		t.Error(err.Error())
		return
	}

	t.Log(recv)
}
