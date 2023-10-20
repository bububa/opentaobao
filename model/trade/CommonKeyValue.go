package trade

import (
	"sync"
)

// CommonKeyValue 结构体
type CommonKeyValue struct {
	// 扩展key。例如：.cpCode（物流品牌）。详以接入文档中描述的场景对接
	Key string `json:"key,omitempty" xml:"key,omitempty"`
	// 扩展value。例如：传运单号。详以接入文档中描述的场景对接
	Value string `json:"value,omitempty" xml:"value,omitempty"`
}

var poolCommonKeyValue = sync.Pool{
	New: func() any {
		return new(CommonKeyValue)
	},
}

// GetCommonKeyValue() 从对象池中获取CommonKeyValue
func GetCommonKeyValue() *CommonKeyValue {
	return poolCommonKeyValue.Get().(*CommonKeyValue)
}

// ReleaseCommonKeyValue 释放CommonKeyValue
func ReleaseCommonKeyValue(v *CommonKeyValue) {
	v.Key = ""
	v.Value = ""
	poolCommonKeyValue.Put(v)
}
