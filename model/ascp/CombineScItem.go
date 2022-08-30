package ascp

// CombineScItem 结构体
type CombineScItem struct {
	// 子条目
	SubScItems []SubScItem `json:"sub_sc_items,omitempty" xml:"sub_sc_items>sub_sc_item,omitempty"`
	// 组合货品erp货品id
	CombineScItemId string `json:"combine_sc_item_id,omitempty" xml:"combine_sc_item_id,omitempty"`
	// 组合货品商家编码
	CombineScItemCode string `json:"combine_sc_item_code,omitempty" xml:"combine_sc_item_code,omitempty"`
	// 组合货品名称
	CombineScItemName string `json:"combine_sc_item_name,omitempty" xml:"combine_sc_item_name,omitempty"`
	// 货主编码；优仓分配，长度不会超过32位，不含特殊字符
	OwnerCode string `json:"owner_code,omitempty" xml:"owner_code,omitempty"`
	// 品牌名称
	BrandName string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
	// 类目名称
	CategoryName string `json:"category_name,omitempty" xml:"category_name,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 币种，USD-美元，CNY-人民币，RUB-卢布，JPY-日元，EUR-欧元，GBP-英镑，HKD-港币，NZD-新西兰元，SGD-新加坡元，AUD-澳元，KRW-韩元，THB-泰铢
	Currency string `json:"currency,omitempty" xml:"currency,omitempty"`
	// 零售价（人民币-分）
	RetailPrice int64 `json:"retail_price,omitempty" xml:"retail_price,omitempty"`
}
