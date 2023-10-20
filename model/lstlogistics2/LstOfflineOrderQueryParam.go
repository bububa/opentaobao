package lstlogistics2

import (
	"sync"
)

// LstOfflineOrderQueryParam 结构体
type LstOfflineOrderQueryParam struct {
	// 外部主订单号
	OutOrderId string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// 买家手机号
	BuyerMobile string `json:"buyer_mobile,omitempty" xml:"buyer_mobile,omitempty"`
}

var poolLstOfflineOrderQueryParam = sync.Pool{
	New: func() any {
		return new(LstOfflineOrderQueryParam)
	},
}

// GetLstOfflineOrderQueryParam() 从对象池中获取LstOfflineOrderQueryParam
func GetLstOfflineOrderQueryParam() *LstOfflineOrderQueryParam {
	return poolLstOfflineOrderQueryParam.Get().(*LstOfflineOrderQueryParam)
}

// ReleaseLstOfflineOrderQueryParam 释放LstOfflineOrderQueryParam
func ReleaseLstOfflineOrderQueryParam(v *LstOfflineOrderQueryParam) {
	v.OutOrderId = ""
	v.BuyerMobile = ""
	poolLstOfflineOrderQueryParam.Put(v)
}
