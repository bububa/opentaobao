package wdk

import (
	"sync"
)

// WdkMarketOpenResult 结构体
type WdkMarketOpenResult struct {
	// 123
	Datas []SyncActivitySkuResultBo `json:"datas,omitempty" xml:"datas>sync_activity_sku_result_bo,omitempty"`
	// 活动数据
	ActivityList []SyncActivityResultBo `json:"activity_list,omitempty" xml:"activity_list>sync_activity_result_bo,omitempty"`
	// 123123
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 123123
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误编码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 唯一码信息
	Data *UniqueDiscountCodeBO `json:"data,omitempty" xml:"data,omitempty"`
	// 123123
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolWdkMarketOpenResult = sync.Pool{
	New: func() any {
		return new(WdkMarketOpenResult)
	},
}

// GetWdkMarketOpenResult() 从对象池中获取WdkMarketOpenResult
func GetWdkMarketOpenResult() *WdkMarketOpenResult {
	return poolWdkMarketOpenResult.Get().(*WdkMarketOpenResult)
}

// ReleaseWdkMarketOpenResult 释放WdkMarketOpenResult
func ReleaseWdkMarketOpenResult(v *WdkMarketOpenResult) {
	v.Datas = v.Datas[:0]
	v.ActivityList = v.ActivityList[:0]
	v.ErrorCode = ""
	v.Message = ""
	v.ErrCode = ""
	v.Data = nil
	v.Success = false
	poolWdkMarketOpenResult.Put(v)
}
