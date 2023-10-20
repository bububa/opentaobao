package omniorder

import (
	"sync"
)

// TaobaoOmniDealerOdersListData 结构体
type TaobaoOmniDealerOdersListData struct {
	// 外部系统订单id
	OutOrderId string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// 卖家id
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 主订单id
	BizOrderId int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
}

var poolTaobaoOmniDealerOdersListData = sync.Pool{
	New: func() any {
		return new(TaobaoOmniDealerOdersListData)
	},
}

// GetTaobaoOmniDealerOdersListData() 从对象池中获取TaobaoOmniDealerOdersListData
func GetTaobaoOmniDealerOdersListData() *TaobaoOmniDealerOdersListData {
	return poolTaobaoOmniDealerOdersListData.Get().(*TaobaoOmniDealerOdersListData)
}

// ReleaseTaobaoOmniDealerOdersListData 释放TaobaoOmniDealerOdersListData
func ReleaseTaobaoOmniDealerOdersListData(v *TaobaoOmniDealerOdersListData) {
	v.OutOrderId = ""
	v.SellerId = 0
	v.BizOrderId = 0
	poolTaobaoOmniDealerOdersListData.Put(v)
}
