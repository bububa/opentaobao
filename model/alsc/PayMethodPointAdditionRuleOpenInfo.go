package alsc

import (
	"sync"
)

// PayMethodPointAdditionRuleOpenInfo 结构体
type PayMethodPointAdditionRuleOpenInfo struct {
	// 方法ID
	MethodId string `json:"method_id,omitempty" xml:"method_id,omitempty"`
	// 方法名称
	MethodName string `json:"method_name,omitempty" xml:"method_name,omitempty"`
	// 是否开启
	Enable bool `json:"enable,omitempty" xml:"enable,omitempty"`
}

var poolPayMethodPointAdditionRuleOpenInfo = sync.Pool{
	New: func() any {
		return new(PayMethodPointAdditionRuleOpenInfo)
	},
}

// GetPayMethodPointAdditionRuleOpenInfo() 从对象池中获取PayMethodPointAdditionRuleOpenInfo
func GetPayMethodPointAdditionRuleOpenInfo() *PayMethodPointAdditionRuleOpenInfo {
	return poolPayMethodPointAdditionRuleOpenInfo.Get().(*PayMethodPointAdditionRuleOpenInfo)
}

// ReleasePayMethodPointAdditionRuleOpenInfo 释放PayMethodPointAdditionRuleOpenInfo
func ReleasePayMethodPointAdditionRuleOpenInfo(v *PayMethodPointAdditionRuleOpenInfo) {
	v.MethodId = ""
	v.MethodName = ""
	v.Enable = false
	poolPayMethodPointAdditionRuleOpenInfo.Put(v)
}
