package alitripmerchant

// MarkupInfoVo 结构体
type MarkupInfoVo struct {
	// 每日加价信息
	DailyMarkupPriceList []DailyMarkupPrice `json:"daily_markup_price_list,omitempty" xml:"daily_markup_price_list>daily_markup_price,omitempty"`
	// 加价描述
	MarkupDesc string `json:"markup_desc,omitempty" xml:"markup_desc,omitempty"`
	// 加价金额
	MarkupPrice int64 `json:"markup_price,omitempty" xml:"markup_price,omitempty"`
	// 加价金额总和
	TotalMarkupPrice int64 `json:"total_markup_price,omitempty" xml:"total_markup_price,omitempty"`
}
