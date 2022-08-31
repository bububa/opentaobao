package alitripmerchant

// VoucherInfoVo 结构体
type VoucherInfoVo struct {
	// 预约须知
	AppointmentNotice []BaseDescVo `json:"appointment_notice,omitempty" xml:"appointment_notice>base_desc_vo,omitempty"`
	// 可约日期
	AvailableDateList []string `json:"available_date_list,omitempty" xml:"available_date_list>string,omitempty"`
	// 关联酒店集合
	HotelList []string `json:"hotel_list,omitempty" xml:"hotel_list>string,omitempty"`
	// 关联rpCode集合
	RateCodeList []string `json:"rate_code_list,omitempty" xml:"rate_code_list>string,omitempty"`
	// 加价规则
	MarkupRuleList []string `json:"markup_rule_list,omitempty" xml:"markup_rule_list>string,omitempty"`
	// 代金券id
	VoucherId string `json:"voucher_id,omitempty" xml:"voucher_id,omitempty"`
	// 券状态
	VoucherStatus string `json:"voucher_status,omitempty" xml:"voucher_status,omitempty"`
	// 使用时间
	VoucherUsedDate string `json:"voucher_used_date,omitempty" xml:"voucher_used_date,omitempty"`
	// 订单id
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 券码名称
	VoucherName string `json:"voucher_name,omitempty" xml:"voucher_name,omitempty"`
	// 有效期开始时间
	EffectiveDate string `json:"effective_date,omitempty" xml:"effective_date,omitempty"`
	// 有效期结束时间
	ExpireDate string `json:"expire_date,omitempty" xml:"expire_date,omitempty"`
	// 头图
	Img string `json:"img,omitempty" xml:"img,omitempty"`
	// 雅高merchantId
	MerchantId string `json:"merchant_id,omitempty" xml:"merchant_id,omitempty"`
	// skuId
	SkuId string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 酒店数量
	HotelNumber int64 `json:"hotel_number,omitempty" xml:"hotel_number,omitempty"`
	// 券价格
	VoucherPrice int64 `json:"voucher_price,omitempty" xml:"voucher_price,omitempty"`
	// 券码数量
	VoucherNumber int64 `json:"voucher_number,omitempty" xml:"voucher_number,omitempty"`
	// 可住几晚
	NightAccount int64 `json:"night_account,omitempty" xml:"night_account,omitempty"`
}
