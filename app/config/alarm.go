/*
@Time: 2020-08-04 10:48
@Auth: xujiang
@File: alarm.go
@Software: GoLand
@Desc: TODO
*/
package config

type alarmConfig struct{
	EnablePanicAlarm bool
	WxAlarmConfig struct{
		Enable         bool
		CorpID         string `json:"CorpID"`
		AgentID        string `json:"AgentID"`
		Secret         string `json:"Secret"`
		EncodingAESKey string `json:"EncodingAESKey"`
	}
	AlarmMessage struct {
		Enable 		 bool
		Address string
		TemplateFile string
		MessageType string `json:"MessageType" default:"text"`
	}
}

var AlarmConfig *alarmConfig

func init() {
	registerConfig(AlarmConfig)
}


func (d *alarmConfig)getConfigKeyName()(string,interface{}){
	if AlarmConfig == nil{
		AlarmConfig = new(alarmConfig)
	}
	return "AlarmConfig",AlarmConfig
}