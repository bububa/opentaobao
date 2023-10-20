package wdk

import (
	"sync"
)

// PropField 结构体
type PropField struct {
	// 渠道属性value，取值见key定义
	Value string `json:"value,omitempty" xml:"value,omitempty"`
	// 渠道属性key, 取值为&#34;ONE_HOUR_STATUS&#34; 代表小时达，value=0表示不可售, value=1表示可售；&#34;WAVE_ARRIVE_STATUS&#34; 代表波次达，value=0表示不可售,value=1表示可售
	Key string `json:"key,omitempty" xml:"key,omitempty"`
}

var poolPropField = sync.Pool{
	New: func() any {
		return new(PropField)
	},
}

// GetPropField() 从对象池中获取PropField
func GetPropField() *PropField {
	return poolPropField.Get().(*PropField)
}

// ReleasePropField 释放PropField
func ReleasePropField(v *PropField) {
	v.Value = ""
	v.Key = ""
	poolPropField.Put(v)
}
