package wdk

import (
	"sync"
)

// Operator 结构体
type Operator struct {
	// 操作员编码
	OperatorCode string `json:"operator_code,omitempty" xml:"operator_code,omitempty"`
	// 操作员姓名
	OperatorName string `json:"operator_name,omitempty" xml:"operator_name,omitempty"`
}

var poolOperator = sync.Pool{
	New: func() any {
		return new(Operator)
	},
}

// GetOperator() 从对象池中获取Operator
func GetOperator() *Operator {
	return poolOperator.Get().(*Operator)
}

// ReleaseOperator 释放Operator
func ReleaseOperator(v *Operator) {
	v.OperatorCode = ""
	v.OperatorName = ""
	poolOperator.Put(v)
}
