package wdk

// MarketResult 结构体
type MarketResult struct {
	// 结果数据
	Datas []ItemBuyGiftSku `json:"datas,omitempty" xml:"datas>item_buy_gift_sku,omitempty"`
	// 处理结果
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误编码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误编码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 报名活动成功的商品详情
	Data *ItemCouponSku `json:"data,omitempty" xml:"data,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
