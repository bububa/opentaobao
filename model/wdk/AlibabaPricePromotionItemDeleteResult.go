package wdk

import (
	"sync"
)

// AlibabaPricePromotionItemDeleteResult 结构体
type AlibabaPricePromotionItemDeleteResult struct {
	// data
	DataList []string `json:"data_list,omitempty" xml:"data_list>string,omitempty"`
	// msg
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	ResultCode int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// totalRecord
	TotalRecord int64 `json:"total_record,omitempty" xml:"total_record,omitempty"`
	// success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

var poolAlibabaPricePromotionItemDeleteResult = sync.Pool{
	New: func() any {
		return new(AlibabaPricePromotionItemDeleteResult)
	},
}

// GetAlibabaPricePromotionItemDeleteResult() 从对象池中获取AlibabaPricePromotionItemDeleteResult
func GetAlibabaPricePromotionItemDeleteResult() *AlibabaPricePromotionItemDeleteResult {
	return poolAlibabaPricePromotionItemDeleteResult.Get().(*AlibabaPricePromotionItemDeleteResult)
}

// ReleaseAlibabaPricePromotionItemDeleteResult 释放AlibabaPricePromotionItemDeleteResult
func ReleaseAlibabaPricePromotionItemDeleteResult(v *AlibabaPricePromotionItemDeleteResult) {
	v.DataList = v.DataList[:0]
	v.Message = ""
	v.ResultCode = 0
	v.TotalRecord = 0
	v.IsSuccess = false
	poolAlibabaPricePromotionItemDeleteResult.Put(v)
}
