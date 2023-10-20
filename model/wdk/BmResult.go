package wdk

import (
	"sync"
)

// BmResult 结构体
type BmResult struct {
	// 结果数据
	DataList []AlibabaWdkBmCouponQueryData `json:"data_list,omitempty" xml:"data_list>alibaba_wdk_bm_coupon_query_data,omitempty"`
	// 详细结果
	PublishResults []SkuStockPublishResult `json:"publish_results,omitempty" xml:"publish_results>sku_stock_publish_result,omitempty"`
	// 错误码，当失败返回错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 描述信息，当成功返回OK，当失败返回错误描述
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 额外属性
	ExtData string `json:"ext_data,omitempty" xml:"ext_data,omitempty"`
	// 库存数量
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolBmResult = sync.Pool{
	New: func() any {
		return new(BmResult)
	},
}

// GetBmResult() 从对象池中获取BmResult
func GetBmResult() *BmResult {
	return poolBmResult.Get().(*BmResult)
}

// ReleaseBmResult 释放BmResult
func ReleaseBmResult(v *BmResult) {
	v.DataList = v.DataList[:0]
	v.PublishResults = v.PublishResults[:0]
	v.ErrorCode = ""
	v.Message = ""
	v.ExtData = ""
	v.Data = 0
	v.Success = false
	poolBmResult.Put(v)
}
