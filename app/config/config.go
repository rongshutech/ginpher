/*
@Time: 2020-08-03 10:26
@Auth: xujiang
@File: config.go
@Software: GoLand
@Desc: TODO
*/
package config

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"reflect"
	"strconv"
)

func InitConfig(env string){
	if env == "online"{
		//get config from online
		filePath := "./config/online.json"
		tmpConfigReader := viper.New()
		tmpConfigReader.SetConfigFile(filePath)
		err := tmpConfigReader.ReadInConfig()
		if err != nil{
			log.Errorf("Read remote config error.Error is %s",err.Error())
		}
		if !tmpConfigReader.GetBool("security"){
			err := viper.AddRemoteProvider(tmpConfigReader.GetString("provider"),
				tmpConfigReader.GetString("endpoint"),tmpConfigReader.GetString("path"))
			if err != nil {
				log.Error(err)
			}
		}else{
			err := viper.AddSecureRemoteProvider(tmpConfigReader.GetString("provider"),
				tmpConfigReader.GetString("endpoint"),tmpConfigReader.GetString("path"),
				tmpConfigReader.GetString("secretkeystring"))
			if err != nil {
				log.Error(err)
			}
		}
		viper.SetConfigType(tmpConfigReader.GetString("type"))
		err = viper.ReadRemoteConfig()
		if err != nil{
			log.Error(err)
		}
	}else{
		//get config from local file.
		filePath := "./config/default.json"
		if len(env) != 0{
			filePath = fmt.Sprintf("./config/%s.json",env)
		}
		viper.SetConfigFile(filePath)
		viper.SetConfigType("json")
		err := viper.ReadInConfig()
		if err != nil{
			log.Errorf("Read local config file error, error is %s.",err.Error())
		}
		for _,v := range configMap{
			name,value := v.(configHandler).getConfigKeyName()
			err = viper.UnmarshalKey(name,value)
			setDefaultValue(value)
			if err != nil{
				log.Error(err)
			}
		}
	}

}

var configMap map[string]interface{}

type configHandler interface {
	getConfigKeyName()(string,interface{})
}

func registerConfig(v interface{}){
	if configMap == nil {
		configMap = make(map[string]interface{})
	}
	name,value := v.(configHandler).getConfigKeyName()
	configMap[name] = value
}

func setDefaultValue(val interface{}) {
	t := reflect.TypeOf(val).Elem()
	for i := 0; i < t.NumField(); i++ {
		if _, exists := t.Field(i).Tag.Lookup("default"); !exists {
			continue
		}
		switch  reflect.ValueOf(val).Elem().Field(i).Kind(){
		case reflect.Int,reflect.Int8,reflect.Int16,reflect.Int32,reflect.Int64:
			if reflect.ValueOf(val).Elem().Field(i).Int() == 0 {
				tagVal := reflect.ValueOf(t.Field(i).Tag.Get("default"))
				intTag,_ := strconv.ParseInt(tagVal.String(),10,64)
				reflect.ValueOf(val).Elem().Field(i).SetInt(int64(intTag))
			}
		case reflect.Uint,reflect.Uint8,reflect.Uint16,reflect.Uint32,reflect.Uint64:
			if reflect.ValueOf(val).Elem().Field(i).Uint() == 0 {
				tagVal := reflect.ValueOf(t.Field(i).Tag.Get("default"))
				uintTag,_ := strconv.ParseUint(tagVal.String(),10,64)
				reflect.ValueOf(val).Elem().Field(i).SetUint(uintTag)
			}
		default:
			tagVal := reflect.ValueOf(t.Field(i).Tag.Get("default"))
			reflect.ValueOf(val).Elem().Field(i).Set(tagVal)
		}

	}
}