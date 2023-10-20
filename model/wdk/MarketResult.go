package wdk

import (
	"sync"
)

// MarketResult 结构体
type MarketResult struct {
	// 结果数据
	Datas []ItemBuyGiftSku `json:"datas,omitempty" xml:"datas>item_buy_gift_sku,omitempty"`
	// 处理结果
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误编码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误编码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 创建成功时会返回五道口优惠券活动id
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

var poolMarketResult = sync.Pool{
	New: func() any {
		return new(MarketResult)
	},
}

// GetMarketResult() 从对象池中获取MarketResult
func GetMarketResult() *MarketResult {
	return poolMarketResult.Get().(*MarketResult)
}

// ReleaseMarketResult 释放MarketResult
func ReleaseMarketResult(v *MarketResult) {
	v.Datas = v.Datas[:0]
	v.Message = ""
	v.ErrorCode = ""
	v.ResultCode = ""
	v.Data = 0
	v.Success = false
	v.IsSuccess = false
	poolMarketResult.Put(v)
}
