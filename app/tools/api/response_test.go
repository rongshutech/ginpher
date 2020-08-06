package api

import (
	"testing"
	"time"
)

type testStruct struct {
	A string
	B int
	T time.Time
}

//{"Code":0,"Msg":"成功","Data":{"A":"a","B":1,"T":"2019-10-18T14:23:34.966294+08:00"}}
func TestResponse(t *testing.T) {
	data := testStruct{
		A: "a",
		B: 1,
		T: time.Now(),
	}

	resp := NewResponse(Success, data)
	ts, err := resp.JsonMarshal()
	if err != nil {
		t.Error(err.Error())
		return
	}

	t.Log(string(ts))
}

/*
{"Code":0,"Msg":"成功","Data":[{"A":"a","B":1,"T":"2019-10-18T14:23:00.401788+08:00"},{"A":"aa","B":11,"T":"2019-10-18T14:23:00.401788+08:00"}]}
*/
func TestArrayResponse(t *testing.T) {
	data := []testStruct{
		{
			A: "a",
			B: 1,
			T: time.Now(),
		},
		{
			A: "aa",
			B: 11,
			T: time.Now(),
		},
	}

	resp := NewResponse(Success, data)
	ts, err := resp.JsonMarshal()
	if err != nil {
		t.Error(err.Error())
		return
	}

	t.Log(string(ts))

}
