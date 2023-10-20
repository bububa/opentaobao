package hotel

import (
	"sync"
)

// HotelPricesResult 结构体
type HotelPricesResult struct {
	// 错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码，不为0表示有异常
	ErrorCode int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 酒店报价信息
	Module *HotelPricesResponse `json:"module,omitempty" xml:"module,omitempty"`
}

var poolHotelPricesResult = sync.Pool{
	New: func() any {
		return new(HotelPricesResult)
	},
}

// GetHotelPricesResult() 从对象池中获取HotelPricesResult
func GetHotelPricesResult() *HotelPricesResult {
	return poolHotelPricesResult.Get().(*HotelPricesResult)
}

// ReleaseHotelPricesResult 释放HotelPricesResult
func ReleaseHotelPricesResult(v *HotelPricesResult) {
	v.ErrorMessage = ""
	v.ErrorCode = 0
	v.Module = nil
	poolHotelPricesResult.Put(v)
}
