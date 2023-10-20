package xhotelitem

import (
	"sync"
)

// TaobaoXhotelItemSelectionSellerStatExposureModule 结构体
type TaobaoXhotelItemSelectionSellerStatExposureModule struct {
	// 返回结果
	SellerStatExposureElementList []SellerStatExposureElementList `json:"seller_stat_exposure_element_list,omitempty" xml:"seller_stat_exposure_element_list>seller_stat_exposure_element_list,omitempty"`
}

var poolTaobaoXhotelItemSelectionSellerStatExposureModule = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelItemSelectionSellerStatExposureModule)
	},
}

// GetTaobaoXhotelItemSelectionSellerStatExposureModule() 从对象池中获取TaobaoXhotelItemSelectionSellerStatExposureModule
func GetTaobaoXhotelItemSelectionSellerStatExposureModule() *TaobaoXhotelItemSelectionSellerStatExposureModule {
	return poolTaobaoXhotelItemSelectionSellerStatExposureModule.Get().(*TaobaoXhotelItemSelectionSellerStatExposureModule)
}

// ReleaseTaobaoXhotelItemSelectionSellerStatExposureModule 释放TaobaoXhotelItemSelectionSellerStatExposureModule
func ReleaseTaobaoXhotelItemSelectionSellerStatExposureModule(v *TaobaoXhotelItemSelectionSellerStatExposureModule) {
	v.SellerStatExposureElementList = v.SellerStatExposureElementList[:0]
	poolTaobaoXhotelItemSelectionSellerStatExposureModule.Put(v)
}
