package btrip

import (
	"sync"
)

// CabinInfo 结构体
type CabinInfo struct {
	// 改签费用信息
	ModifyPriceList []ModifyPrice `json:"modify_price_list,omitempty" xml:"modify_price_list>modify_price,omitempty"`
	// av
	LeftNum string `json:"left_num,omitempty" xml:"left_num,omitempty"`
	// 舱位代码
	Cabin string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	// 舱等
	CabinClass string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	// 舱位描述
	CabinDesc string `json:"cabin_desc,omitempty" xml:"cabin_desc,omitempty"`
	// 舱等折扣，八折用80表示
	ChildCabin string `json:"child_cabin,omitempty" xml:"child_cabin,omitempty"`
	// 商品Id
	OtaItemid string `json:"ota_itemid,omitempty" xml:"ota_itemid,omitempty"`
	// 改签航班的退改签规则
	ChangeOtaItemRuleRq *ChangeOtaItemRuleRq `json:"change_ota_item_rule_rq,omitempty" xml:"change_ota_item_rule_rq,omitempty"`
	// 舱位折扣
	CabinDiscount int64 `json:"cabin_discount,omitempty" xml:"cabin_discount,omitempty"`
}

var poolCabinInfo = sync.Pool{
	New: func() any {
		return new(CabinInfo)
	},
}

// GetCabinInfo() 从对象池中获取CabinInfo
func GetCabinInfo() *CabinInfo {
	return poolCabinInfo.Get().(*CabinInfo)
}

// ReleaseCabinInfo 释放CabinInfo
func ReleaseCabinInfo(v *CabinInfo) {
	v.ModifyPriceList = v.ModifyPriceList[:0]
	v.LeftNum = ""
	v.Cabin = ""
	v.CabinClass = ""
	v.CabinDesc = ""
	v.ChildCabin = ""
	v.OtaItemid = ""
	v.ChangeOtaItemRuleRq = nil
	v.CabinDiscount = 0
	poolCabinInfo.Put(v)
}
