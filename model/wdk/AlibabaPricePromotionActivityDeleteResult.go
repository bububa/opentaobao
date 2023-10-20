package wdk

import (
	"sync"
)

// AlibabaPricePromotionActivityDeleteResult 结构体
type AlibabaPricePromotionActivityDeleteResult struct {
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

var poolAlibabaPricePromotionActivityDeleteResult = sync.Pool{
	New: func() any {
		return new(AlibabaPricePromotionActivityDeleteResult)
	},
}

// GetAlibabaPricePromotionActivityDeleteResult() 从对象池中获取AlibabaPricePromotionActivityDeleteResult
func GetAlibabaPricePromotionActivityDeleteResult() *AlibabaPricePromotionActivityDeleteResult {
	return poolAlibabaPricePromotionActivityDeleteResult.Get().(*AlibabaPricePromotionActivityDeleteResult)
}

// ReleaseAlibabaPricePromotionActivityDeleteResult 释放AlibabaPricePromotionActivityDeleteResult
func ReleaseAlibabaPricePromotionActivityDeleteResult(v *AlibabaPricePromotionActivityDeleteResult) {
	v.DataList = v.DataList[:0]
	v.Message = ""
	v.ResultCode = 0
	v.TotalRecord = 0
	v.IsSuccess = false
	poolAlibabaPricePromotionActivityDeleteResult.Put(v)
}
