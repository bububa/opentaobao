package fenxiao

import (
	"sync"
)

// CnskuRelationOperateOption 结构体
type CnskuRelationOperateOption struct {
	// 操作类型(默认为DELET)
	OperateType string `json:"operate_type,omitempty" xml:"operate_type,omitempty"`
}

var poolCnskuRelationOperateOption = sync.Pool{
	New: func() any {
		return new(CnskuRelationOperateOption)
	},
}

// GetCnskuRelationOperateOption() 从对象池中获取CnskuRelationOperateOption
func GetCnskuRelationOperateOption() *CnskuRelationOperateOption {
	return poolCnskuRelationOperateOption.Get().(*CnskuRelationOperateOption)
}

// ReleaseCnskuRelationOperateOption 释放CnskuRelationOperateOption
func ReleaseCnskuRelationOperateOption(v *CnskuRelationOperateOption) {
	v.OperateType = ""
	poolCnskuRelationOperateOption.Put(v)
}
