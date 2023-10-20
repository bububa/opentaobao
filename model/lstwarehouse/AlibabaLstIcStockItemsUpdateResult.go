package lstwarehouse

import (
	"sync"
)

// AlibabaLstIcStockItemsUpdateResult 结构体
type AlibabaLstIcStockItemsUpdateResult struct {
	// 列表
	ContentList []Content `json:"content_list,omitempty" xml:"content_list>content,omitempty"`
	// errorMsg
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaLstIcStockItemsUpdateResult = sync.Pool{
	New: func() any {
		return new(AlibabaLstIcStockItemsUpdateResult)
	},
}

// GetAlibabaLstIcStockItemsUpdateResult() 从对象池中获取AlibabaLstIcStockItemsUpdateResult
func GetAlibabaLstIcStockItemsUpdateResult() *AlibabaLstIcStockItemsUpdateResult {
	return poolAlibabaLstIcStockItemsUpdateResult.Get().(*AlibabaLstIcStockItemsUpdateResult)
}

// ReleaseAlibabaLstIcStockItemsUpdateResult 释放AlibabaLstIcStockItemsUpdateResult
func ReleaseAlibabaLstIcStockItemsUpdateResult(v *AlibabaLstIcStockItemsUpdateResult) {
	v.ContentList = v.ContentList[:0]
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.Success = false
	poolAlibabaLstIcStockItemsUpdateResult.Put(v)
}
