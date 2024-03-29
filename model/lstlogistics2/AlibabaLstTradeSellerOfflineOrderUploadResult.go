package lstlogistics2

import (
	"sync"
)

// AlibabaLstTradeSellerOfflineOrderUploadResult 结构体
type AlibabaLstTradeSellerOfflineOrderUploadResult struct {
	// 系统自动生成
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 出货单id
	Content int64 `json:"content,omitempty" xml:"content,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaLstTradeSellerOfflineOrderUploadResult = sync.Pool{
	New: func() any {
		return new(AlibabaLstTradeSellerOfflineOrderUploadResult)
	},
}

// GetAlibabaLstTradeSellerOfflineOrderUploadResult() 从对象池中获取AlibabaLstTradeSellerOfflineOrderUploadResult
func GetAlibabaLstTradeSellerOfflineOrderUploadResult() *AlibabaLstTradeSellerOfflineOrderUploadResult {
	return poolAlibabaLstTradeSellerOfflineOrderUploadResult.Get().(*AlibabaLstTradeSellerOfflineOrderUploadResult)
}

// ReleaseAlibabaLstTradeSellerOfflineOrderUploadResult 释放AlibabaLstTradeSellerOfflineOrderUploadResult
func ReleaseAlibabaLstTradeSellerOfflineOrderUploadResult(v *AlibabaLstTradeSellerOfflineOrderUploadResult) {
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.Content = 0
	v.Success = false
	poolAlibabaLstTradeSellerOfflineOrderUploadResult.Put(v)
}
