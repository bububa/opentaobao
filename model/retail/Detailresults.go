package retail

import (
	"sync"
)

// Detailresults 结构体
type Detailresults struct {
	// outOrderId
	OutOrderId string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// errorMessage
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolDetailresults = sync.Pool{
	New: func() any {
		return new(Detailresults)
	},
}

// GetDetailresults() 从对象池中获取Detailresults
func GetDetailresults() *Detailresults {
	return poolDetailresults.Get().(*Detailresults)
}

// ReleaseDetailresults 释放Detailresults
func ReleaseDetailresults(v *Detailresults) {
	v.OutOrderId = ""
	v.ErrorCode = ""
	v.ErrorMessage = ""
	v.Success = false
	poolDetailresults.Put(v)
}
