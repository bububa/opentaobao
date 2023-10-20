package itpolicy

import (
	"sync"
)

// ErrorFareRow 结构体
type ErrorFareRow struct {
	// 错误描述
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误行数
	RowNum int64 `json:"row_num,omitempty" xml:"row_num,omitempty"`
}

var poolErrorFareRow = sync.Pool{
	New: func() any {
		return new(ErrorFareRow)
	},
}

// GetErrorFareRow() 从对象池中获取ErrorFareRow
func GetErrorFareRow() *ErrorFareRow {
	return poolErrorFareRow.Get().(*ErrorFareRow)
}

// ReleaseErrorFareRow 释放ErrorFareRow
func ReleaseErrorFareRow(v *ErrorFareRow) {
	v.ErrorMsg = ""
	v.RowNum = 0
	poolErrorFareRow.Put(v)
}
