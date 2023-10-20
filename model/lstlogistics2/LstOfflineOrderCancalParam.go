package lstlogistics2

import (
	"sync"
)

// LstOfflineOrderCancalParam 结构体
type LstOfflineOrderCancalParam struct {
	// 买家手机号
	BuyerMobile string `json:"buyer_mobile,omitempty" xml:"buyer_mobile,omitempty"`
	// 外部主订单号
	OutOrderId string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
}

var poolLstOfflineOrderCancalParam = sync.Pool{
	New: func() any {
		return new(LstOfflineOrderCancalParam)
	},
}

// GetLstOfflineOrderCancalParam() 从对象池中获取LstOfflineOrderCancalParam
func GetLstOfflineOrderCancalParam() *LstOfflineOrderCancalParam {
	return poolLstOfflineOrderCancalParam.Get().(*LstOfflineOrderCancalParam)
}

// ReleaseLstOfflineOrderCancalParam 释放LstOfflineOrderCancalParam
func ReleaseLstOfflineOrderCancalParam(v *LstOfflineOrderCancalParam) {
	v.BuyerMobile = ""
	v.OutOrderId = ""
	poolLstOfflineOrderCancalParam.Put(v)
}
