package api

type Receive struct {
	Code Code                   `json:"Code"`
	Msg  string                 `json:"Msg"`
	Data map[string]interface{} `json:"Data,omitempty"`
}

type ReceiveArray struct {
	Code Code                     `json:"Code"`
	Msg  string                   `json:"Msg"`
	Data []map[string]interface{} `json:"Data,omitempty"`
}

type ReceiveInterface struct {
	Code Code        `json:"Code"`
	Msg  string      `json:"Msg"`
	Data interface{} `json:"Data,omitempty"`
}
