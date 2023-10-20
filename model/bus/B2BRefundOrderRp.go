package bus

import (
	"sync"
)

// B2BRefundOrderRp 结构体
type B2BRefundOrderRp struct {
	// results
	Results []string `json:"results,omitempty" xml:"results>string,omitempty"`
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// errorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// success1
	Success1 bool `json:"success1,omitempty" xml:"success1,omitempty"`
}

var poolB2BRefundOrderRp = sync.Pool{
	New: func() any {
		return new(B2BRefundOrderRp)
	},
}

// GetB2BRefundOrderRp() 从对象池中获取B2BRefundOrderRp
func GetB2BRefundOrderRp() *B2BRefundOrderRp {
	return poolB2BRefundOrderRp.Get().(*B2BRefundOrderRp)
}

// ReleaseB2BRefundOrderRp 释放B2BRefundOrderRp
func ReleaseB2BRefundOrderRp(v *B2BRefundOrderRp) {
	v.Results = v.Results[:0]
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Success = false
	v.Success1 = false
	poolB2BRefundOrderRp.Put(v)
}
