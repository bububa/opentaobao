package wdk

import (
	"sync"
)

// TaobaoWdkIotDeviceadminMqttDeviceGetwithtokenHmResult 结构体
type TaobaoWdkIotDeviceadminMqttDeviceGetwithtokenHmResult struct {
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// model
	Model *MqttDeviceInfoDto `json:"model,omitempty" xml:"model,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoWdkIotDeviceadminMqttDeviceGetwithtokenHmResult = sync.Pool{
	New: func() any {
		return new(TaobaoWdkIotDeviceadminMqttDeviceGetwithtokenHmResult)
	},
}

// GetTaobaoWdkIotDeviceadminMqttDeviceGetwithtokenHmResult() 从对象池中获取TaobaoWdkIotDeviceadminMqttDeviceGetwithtokenHmResult
func GetTaobaoWdkIotDeviceadminMqttDeviceGetwithtokenHmResult() *TaobaoWdkIotDeviceadminMqttDeviceGetwithtokenHmResult {
	return poolTaobaoWdkIotDeviceadminMqttDeviceGetwithtokenHmResult.Get().(*TaobaoWdkIotDeviceadminMqttDeviceGetwithtokenHmResult)
}

// ReleaseTaobaoWdkIotDeviceadminMqttDeviceGetwithtokenHmResult 释放TaobaoWdkIotDeviceadminMqttDeviceGetwithtokenHmResult
func ReleaseTaobaoWdkIotDeviceadminMqttDeviceGetwithtokenHmResult(v *TaobaoWdkIotDeviceadminMqttDeviceGetwithtokenHmResult) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.Success = false
	poolTaobaoWdkIotDeviceadminMqttDeviceGetwithtokenHmResult.Put(v)
}
