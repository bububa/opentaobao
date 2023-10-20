package lstlogistics2

import (
	"sync"
)

// AlibabaLstTradeSellerOfflineOrderCancelResult 结构体
type AlibabaLstTradeSellerOfflineOrderCancelResult struct {
	// 系统自动生成
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 是否取消成功
	Content bool `json:"content,omitempty" xml:"content,omitempty"`
}

var poolAlibabaLstTradeSellerOfflineOrderCancelResult = sync.Pool{
	New: func() any {
		return new(AlibabaLstTradeSellerOfflineOrderCancelResult)
	},
}

// GetAlibabaLstTradeSellerOfflineOrderCancelResult() 从对象池中获取AlibabaLstTradeSellerOfflineOrderCancelResult
func GetAlibabaLstTradeSellerOfflineOrderCancelResult() *AlibabaLstTradeSellerOfflineOrderCancelResult {
	return poolAlibabaLstTradeSellerOfflineOrderCancelResult.Get().(*AlibabaLstTradeSellerOfflineOrderCancelResult)
}

// ReleaseAlibabaLstTradeSellerOfflineOrderCancelResult 释放AlibabaLstTradeSellerOfflineOrderCancelResult
func ReleaseAlibabaLstTradeSellerOfflineOrderCancelResult(v *AlibabaLstTradeSellerOfflineOrderCancelResult) {
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.Success = false
	v.Content = false
	poolAlibabaLstTradeSellerOfflineOrderCancelResult.Put(v)
}
