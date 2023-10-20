package traveltrade

import (
	"sync"
)

// OrderTipRuleInfo 结构体
type OrderTipRuleInfo struct {
	// 字段描述
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 字段正则表达式，字段值必须和正则匹配，否则上传服务信息会报错
	Regex string `json:"regex,omitempty" xml:"regex,omitempty"`
	// 字段是否必须
	Require bool `json:"require,omitempty" xml:"require,omitempty"`
}

var poolOrderTipRuleInfo = sync.Pool{
	New: func() any {
		return new(OrderTipRuleInfo)
	},
}

// GetOrderTipRuleInfo() 从对象池中获取OrderTipRuleInfo
func GetOrderTipRuleInfo() *OrderTipRuleInfo {
	return poolOrderTipRuleInfo.Get().(*OrderTipRuleInfo)
}

// ReleaseOrderTipRuleInfo 释放OrderTipRuleInfo
func ReleaseOrderTipRuleInfo(v *OrderTipRuleInfo) {
	v.Msg = ""
	v.Regex = ""
	v.Require = false
	poolOrderTipRuleInfo.Put(v)
}
