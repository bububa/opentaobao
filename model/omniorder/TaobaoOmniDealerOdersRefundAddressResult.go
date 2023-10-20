package omniorder

import (
	"sync"
)

// TaobaoOmniDealerOdersRefundAddressResult 结构体
type TaobaoOmniDealerOdersRefundAddressResult struct {
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 是否成功
	Success string `json:"success,omitempty" xml:"success,omitempty"`
	// 经销商订单退货地址
	Data *ScbRefundAddressDto `json:"data,omitempty" xml:"data,omitempty"`
}

var poolTaobaoOmniDealerOdersRefundAddressResult = sync.Pool{
	New: func() any {
		return new(TaobaoOmniDealerOdersRefundAddressResult)
	},
}

// GetTaobaoOmniDealerOdersRefundAddressResult() 从对象池中获取TaobaoOmniDealerOdersRefundAddressResult
func GetTaobaoOmniDealerOdersRefundAddressResult() *TaobaoOmniDealerOdersRefundAddressResult {
	return poolTaobaoOmniDealerOdersRefundAddressResult.Get().(*TaobaoOmniDealerOdersRefundAddressResult)
}

// ReleaseTaobaoOmniDealerOdersRefundAddressResult 释放TaobaoOmniDealerOdersRefundAddressResult
func ReleaseTaobaoOmniDealerOdersRefundAddressResult(v *TaobaoOmniDealerOdersRefundAddressResult) {
	v.Message = ""
	v.Success = ""
	v.Data = nil
	poolTaobaoOmniDealerOdersRefundAddressResult.Put(v)
}
