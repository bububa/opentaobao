package lsttrade

import (
	"sync"
)

// AlibabaLstTradeSellerOrderListQueryResult 结构体
type AlibabaLstTradeSellerOrderListQueryResult struct {
	// 信息
	ContentList []Content `json:"content_list,omitempty" xml:"content_list>content,omitempty"`
	// 失败信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 系统自动生成
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// 数量
	Size int64 `json:"size,omitempty" xml:"size,omitempty"`
	// 当前页
	Page int64 `json:"page,omitempty" xml:"page,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaLstTradeSellerOrderListQueryResult = sync.Pool{
	New: func() any {
		return new(AlibabaLstTradeSellerOrderListQueryResult)
	},
}

// GetAlibabaLstTradeSellerOrderListQueryResult() 从对象池中获取AlibabaLstTradeSellerOrderListQueryResult
func GetAlibabaLstTradeSellerOrderListQueryResult() *AlibabaLstTradeSellerOrderListQueryResult {
	return poolAlibabaLstTradeSellerOrderListQueryResult.Get().(*AlibabaLstTradeSellerOrderListQueryResult)
}

// ReleaseAlibabaLstTradeSellerOrderListQueryResult 释放AlibabaLstTradeSellerOrderListQueryResult
func ReleaseAlibabaLstTradeSellerOrderListQueryResult(v *AlibabaLstTradeSellerOrderListQueryResult) {
	v.ContentList = v.ContentList[:0]
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.Total = 0
	v.Size = 0
	v.Page = 0
	v.Success = false
	poolAlibabaLstTradeSellerOrderListQueryResult.Put(v)
}
