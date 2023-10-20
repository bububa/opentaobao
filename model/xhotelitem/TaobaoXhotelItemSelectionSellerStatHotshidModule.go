package xhotelitem

import (
	"sync"
)

// TaobaoXhotelItemSelectionSellerStatHotshidModule 结构体
type TaobaoXhotelItemSelectionSellerStatHotshidModule struct {
	// 热销标准酒店中卖家覆盖的数量
	CoveredHidAmount string `json:"covered_hid_amount,omitempty" xml:"covered_hid_amount,omitempty"`
	// 热销标准酒店中卖家可售的酒店数量
	CanSaleHidAmount string `json:"can_sale_hid_amount,omitempty" xml:"can_sale_hid_amount,omitempty"`
}

var poolTaobaoXhotelItemSelectionSellerStatHotshidModule = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelItemSelectionSellerStatHotshidModule)
	},
}

// GetTaobaoXhotelItemSelectionSellerStatHotshidModule() 从对象池中获取TaobaoXhotelItemSelectionSellerStatHotshidModule
func GetTaobaoXhotelItemSelectionSellerStatHotshidModule() *TaobaoXhotelItemSelectionSellerStatHotshidModule {
	return poolTaobaoXhotelItemSelectionSellerStatHotshidModule.Get().(*TaobaoXhotelItemSelectionSellerStatHotshidModule)
}

// ReleaseTaobaoXhotelItemSelectionSellerStatHotshidModule 释放TaobaoXhotelItemSelectionSellerStatHotshidModule
func ReleaseTaobaoXhotelItemSelectionSellerStatHotshidModule(v *TaobaoXhotelItemSelectionSellerStatHotshidModule) {
	v.CoveredHidAmount = ""
	v.CanSaleHidAmount = ""
	poolTaobaoXhotelItemSelectionSellerStatHotshidModule.Put(v)
}
