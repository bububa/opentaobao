package lstlogistics2

import (
	"sync"
)

// AlibabaLstTradeSellerOfflineOrderQueryResult 结构体
type AlibabaLstTradeSellerOfflineOrderQueryResult struct {
	// 系统自动生成
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 主订单
	Content *Content `json:"content,omitempty" xml:"content,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaLstTradeSellerOfflineOrderQueryResult = sync.Pool{
	New: func() any {
		return new(AlibabaLstTradeSellerOfflineOrderQueryResult)
	},
}

// GetAlibabaLstTradeSellerOfflineOrderQueryResult() 从对象池中获取AlibabaLstTradeSellerOfflineOrderQueryResult
func GetAlibabaLstTradeSellerOfflineOrderQueryResult() *AlibabaLstTradeSellerOfflineOrderQueryResult {
	return poolAlibabaLstTradeSellerOfflineOrderQueryResult.Get().(*AlibabaLstTradeSellerOfflineOrderQueryResult)
}

// ReleaseAlibabaLstTradeSellerOfflineOrderQueryResult 释放AlibabaLstTradeSellerOfflineOrderQueryResult
func ReleaseAlibabaLstTradeSellerOfflineOrderQueryResult(v *AlibabaLstTradeSellerOfflineOrderQueryResult) {
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.Content = nil
	v.Success = false
	poolAlibabaLstTradeSellerOfflineOrderQueryResult.Put(v)
}
