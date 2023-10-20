package icbulogistics

import (
	"sync"
)

// AlibabaOnetouchLogisticsExpressOrderDetailGetResult 结构体
type AlibabaOnetouchLogisticsExpressOrderDetailGetResult struct {
	// 返回结果描述
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 返回结果编码
	ErrorCode int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 结果数据
	Data *OrderDetailDto `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaOnetouchLogisticsExpressOrderDetailGetResult = sync.Pool{
	New: func() any {
		return new(AlibabaOnetouchLogisticsExpressOrderDetailGetResult)
	},
}

// GetAlibabaOnetouchLogisticsExpressOrderDetailGetResult() 从对象池中获取AlibabaOnetouchLogisticsExpressOrderDetailGetResult
func GetAlibabaOnetouchLogisticsExpressOrderDetailGetResult() *AlibabaOnetouchLogisticsExpressOrderDetailGetResult {
	return poolAlibabaOnetouchLogisticsExpressOrderDetailGetResult.Get().(*AlibabaOnetouchLogisticsExpressOrderDetailGetResult)
}

// ReleaseAlibabaOnetouchLogisticsExpressOrderDetailGetResult 释放AlibabaOnetouchLogisticsExpressOrderDetailGetResult
func ReleaseAlibabaOnetouchLogisticsExpressOrderDetailGetResult(v *AlibabaOnetouchLogisticsExpressOrderDetailGetResult) {
	v.ErrorMessage = ""
	v.ErrorCode = 0
	v.Data = nil
	v.Success = false
	poolAlibabaOnetouchLogisticsExpressOrderDetailGetResult.Put(v)
}
