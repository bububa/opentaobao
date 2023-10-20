package tmallsc

import (
	"sync"
)

// Value 结构体
type Value struct {
	// 服务项
	ServiceSkuCode string `json:"service_sku_code,omitempty" xml:"service_sku_code,omitempty"`
	// 服务项展示名称
	DisplayName string `json:"display_name,omitempty" xml:"display_name,omitempty"`
	// 是否为补差服务项
	NeedExtPay bool `json:"need_ext_pay,omitempty" xml:"need_ext_pay,omitempty"`
}

var poolValue = sync.Pool{
	New: func() any {
		return new(Value)
	},
}

// GetValue() 从对象池中获取Value
func GetValue() *Value {
	return poolValue.Get().(*Value)
}

// ReleaseValue 释放Value
func ReleaseValue(v *Value) {
	v.ServiceSkuCode = ""
	v.DisplayName = ""
	v.NeedExtPay = false
	poolValue.Put(v)
}
