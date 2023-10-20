package tbrefund

import (
	"sync"
)

// Operation 结构体
type Operation struct {
	// 操作编码
	OperationCode string `json:"operation_code,omitempty" xml:"operation_code,omitempty"`
	// 操作提示文案
	Tips string `json:"tips,omitempty" xml:"tips,omitempty"`
}

var poolOperation = sync.Pool{
	New: func() any {
		return new(Operation)
	},
}

// GetOperation() 从对象池中获取Operation
func GetOperation() *Operation {
	return poolOperation.Get().(*Operation)
}

// ReleaseOperation 释放Operation
func ReleaseOperation(v *Operation) {
	v.OperationCode = ""
	v.Tips = ""
	poolOperation.Put(v)
}
