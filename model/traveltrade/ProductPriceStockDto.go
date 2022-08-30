package traveltrade

// ProductPriceStockDto 结构体
type ProductPriceStockDto struct {
	// 场次价库信息
	Sessions []ProductSessionDto `json:"sessions,omitempty" xml:"sessions>product_session_dto,omitempty"`
	// 日期。yyyy-MM-dd
	Date string `json:"date,omitempty" xml:"date,omitempty"`
	// 库存
	Stock int64 `json:"stock,omitempty" xml:"stock,omitempty"`
	// 产品结算单价。单位：分
	WholesalePrice int64 `json:"wholesale_price,omitempty" xml:"wholesale_price,omitempty"`
	// 产品销售单价。单位：分
	RetailPrice int64 `json:"retail_price,omitempty" xml:"retail_price,omitempty"`
	// 是否可售卖；true：可售卖
	CanSell bool `json:"can_sell,omitempty" xml:"can_sell,omitempty"`
}
