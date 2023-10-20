package ascm

import (
	"sync"
)

// SettlementGatewayResult 结构体
type SettlementGatewayResult struct {
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// errorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// result
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

var poolSettlementGatewayResult = sync.Pool{
	New: func() any {
		return new(SettlementGatewayResult)
	},
}

// GetSettlementGatewayResult() 从对象池中获取SettlementGatewayResult
func GetSettlementGatewayResult() *SettlementGatewayResult {
	return poolSettlementGatewayResult.Get().(*SettlementGatewayResult)
}

// ReleaseSettlementGatewayResult 释放SettlementGatewayResult
func ReleaseSettlementGatewayResult(v *SettlementGatewayResult) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Result = false
	poolSettlementGatewayResult.Put(v)
}
