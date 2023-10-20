package bus

import (
	"sync"
)

// B2BOrderQueryRp 结构体
type B2BOrderQueryRp struct {
	// 错误代码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误描述
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 订单对象
	B2bBusOrderInfo *B2BBusOrderInfo `json:"b2b_bus_order_info,omitempty" xml:"b2b_bus_order_info,omitempty"`
	// 是否查询成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolB2BOrderQueryRp = sync.Pool{
	New: func() any {
		return new(B2BOrderQueryRp)
	},
}

// GetB2BOrderQueryRp() 从对象池中获取B2BOrderQueryRp
func GetB2BOrderQueryRp() *B2BOrderQueryRp {
	return poolB2BOrderQueryRp.Get().(*B2BOrderQueryRp)
}

// ReleaseB2BOrderQueryRp 释放B2BOrderQueryRp
func ReleaseB2BOrderQueryRp(v *B2BOrderQueryRp) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.B2bBusOrderInfo = nil
	v.Success = false
	poolB2BOrderQueryRp.Put(v)
}
