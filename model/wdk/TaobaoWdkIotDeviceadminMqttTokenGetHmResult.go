package wdk

import (
	"sync"
)

// TaobaoWdkIotDeviceadminMqttTokenGetHmResult 结构体
type TaobaoWdkIotDeviceadminMqttTokenGetHmResult struct {
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// model
	Model *MqttDeviceInfoDto `json:"model,omitempty" xml:"model,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoWdkIotDeviceadminMqttTokenGetHmResult = sync.Pool{
	New: func() any {
		return new(TaobaoWdkIotDeviceadminMqttTokenGetHmResult)
	},
}

// GetTaobaoWdkIotDeviceadminMqttTokenGetHmResult() 从对象池中获取TaobaoWdkIotDeviceadminMqttTokenGetHmResult
func GetTaobaoWdkIotDeviceadminMqttTokenGetHmResult() *TaobaoWdkIotDeviceadminMqttTokenGetHmResult {
	return poolTaobaoWdkIotDeviceadminMqttTokenGetHmResult.Get().(*TaobaoWdkIotDeviceadminMqttTokenGetHmResult)
}

// ReleaseTaobaoWdkIotDeviceadminMqttTokenGetHmResult 释放TaobaoWdkIotDeviceadminMqttTokenGetHmResult
func ReleaseTaobaoWdkIotDeviceadminMqttTokenGetHmResult(v *TaobaoWdkIotDeviceadminMqttTokenGetHmResult) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.Success = false
	poolTaobaoWdkIotDeviceadminMqttTokenGetHmResult.Put(v)
}
