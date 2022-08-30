package btrip

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
	// 舱位折扣
	CabinDiscount int64 `json:"cabin_discount,omitempty" xml:"cabin_discount,omitempty"`
}
