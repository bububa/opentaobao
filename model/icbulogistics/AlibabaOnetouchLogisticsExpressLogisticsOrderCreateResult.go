package icbulogistics

import (
	"sync"
)

// AlibabaOnetouchLogisticsExpressLogisticsOrderCreateResult 结构体
type AlibabaOnetouchLogisticsExpressLogisticsOrderCreateResult struct {
	// 返回结果描述
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 结果对象
	Values *ExpressFreightDto `json:"values,omitempty" xml:"values,omitempty"`
	// 返回结果编码
	ErrorCode int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaOnetouchLogisticsExpressLogisticsOrderCreateResult = sync.Pool{
	New: func() any {
		return new(AlibabaOnetouchLogisticsExpressLogisticsOrderCreateResult)
	},
}

// GetAlibabaOnetouchLogisticsExpressLogisticsOrderCreateResult() 从对象池中获取AlibabaOnetouchLogisticsExpressLogisticsOrderCreateResult
func GetAlibabaOnetouchLogisticsExpressLogisticsOrderCreateResult() *AlibabaOnetouchLogisticsExpressLogisticsOrderCreateResult {
	return poolAlibabaOnetouchLogisticsExpressLogisticsOrderCreateResult.Get().(*AlibabaOnetouchLogisticsExpressLogisticsOrderCreateResult)
}

// ReleaseAlibabaOnetouchLogisticsExpressLogisticsOrderCreateResult 释放AlibabaOnetouchLogisticsExpressLogisticsOrderCreateResult
func ReleaseAlibabaOnetouchLogisticsExpressLogisticsOrderCreateResult(v *AlibabaOnetouchLogisticsExpressLogisticsOrderCreateResult) {
	v.ErrorMessage = ""
	v.Values = nil
	v.ErrorCode = 0
	v.Success = false
	poolAlibabaOnetouchLogisticsExpressLogisticsOrderCreateResult.Put(v)
}
