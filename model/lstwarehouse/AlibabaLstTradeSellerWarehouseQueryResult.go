package lstwarehouse

import (
	"sync"
)

// AlibabaLstTradeSellerWarehouseQueryResult 结构体
type AlibabaLstTradeSellerWarehouseQueryResult struct {
	// 记录
	ContentList []Content `json:"content_list,omitempty" xml:"content_list>content,omitempty"`
	// 系统自动生成
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 总数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// 每页最大记录数
	Size int64 `json:"size,omitempty" xml:"size,omitempty"`
	// 页码
	Page int64 `json:"page,omitempty" xml:"page,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaLstTradeSellerWarehouseQueryResult = sync.Pool{
	New: func() any {
		return new(AlibabaLstTradeSellerWarehouseQueryResult)
	},
}

// GetAlibabaLstTradeSellerWarehouseQueryResult() 从对象池中获取AlibabaLstTradeSellerWarehouseQueryResult
func GetAlibabaLstTradeSellerWarehouseQueryResult() *AlibabaLstTradeSellerWarehouseQueryResult {
	return poolAlibabaLstTradeSellerWarehouseQueryResult.Get().(*AlibabaLstTradeSellerWarehouseQueryResult)
}

// ReleaseAlibabaLstTradeSellerWarehouseQueryResult 释放AlibabaLstTradeSellerWarehouseQueryResult
func ReleaseAlibabaLstTradeSellerWarehouseQueryResult(v *AlibabaLstTradeSellerWarehouseQueryResult) {
	v.ContentList = v.ContentList[:0]
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.Total = 0
	v.Size = 0
	v.Page = 0
	v.Success = false
	poolAlibabaLstTradeSellerWarehouseQueryResult.Put(v)
}
