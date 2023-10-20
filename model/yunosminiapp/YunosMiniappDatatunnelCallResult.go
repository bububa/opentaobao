package yunosminiapp

import (
	"sync"
)

// YunosMiniappDatatunnelCallResult 结构体
type YunosMiniappDatatunnelCallResult struct {
	// cp对应的code
	CpCode string `json:"cp_code,omitempty" xml:"cp_code,omitempty"`
	// 结果详细内容
	Detail string `json:"detail,omitempty" xml:"detail,omitempty"`
}

var poolYunosMiniappDatatunnelCallResult = sync.Pool{
	New: func() any {
		return new(YunosMiniappDatatunnelCallResult)
	},
}

// GetYunosMiniappDatatunnelCallResult() 从对象池中获取YunosMiniappDatatunnelCallResult
func GetYunosMiniappDatatunnelCallResult() *YunosMiniappDatatunnelCallResult {
	return poolYunosMiniappDatatunnelCallResult.Get().(*YunosMiniappDatatunnelCallResult)
}

// ReleaseYunosMiniappDatatunnelCallResult 释放YunosMiniappDatatunnelCallResult
func ReleaseYunosMiniappDatatunnelCallResult(v *YunosMiniappDatatunnelCallResult) {
	v.CpCode = ""
	v.Detail = ""
	poolYunosMiniappDatatunnelCallResult.Put(v)
}
