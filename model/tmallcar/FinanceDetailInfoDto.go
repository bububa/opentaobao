package tmallcar

// FinanceDetailInfoDto 结构体
type FinanceDetailInfoDto struct {
	// 品牌名称
	BrandName string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
	// 车系名称
	SeriesName string `json:"series_name,omitempty" xml:"series_name,omitempty"`
	// 车型名称
	ModelName string `json:"model_name,omitempty" xml:"model_name,omitempty"`
	// 业务内容
	BizContent string `json:"biz_content,omitempty" xml:"biz_content,omitempty"`
	// 外部金融订单号
	OuterFinanceOrderId string `json:"outer_finance_order_id,omitempty" xml:"outer_finance_order_id,omitempty"`
	// 厂商名称
	ManufactureName string `json:"manufacture_name,omitempty" xml:"manufacture_name,omitempty"`
	// 商品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 订单金额
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// 品牌id
	BrandId int64 `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
	// 车系id
	SeriesId int64 `json:"series_id,omitempty" xml:"series_id,omitempty"`
	// 车型id
	ModelId int64 `json:"model_id,omitempty" xml:"model_id,omitempty"`
	// bookingShopInfo
	BookingShopInfo *AutoShopDto `json:"booking_shop_info,omitempty" xml:"booking_shop_info,omitempty"`
	// 购车人信息
	PurchaseUserInfo *PurchaseUserDto `json:"purchase_user_info,omitempty" xml:"purchase_user_info,omitempty"`
	// 收货地址
	DeliveryAddress *DeliveryAddressDto `json:"delivery_address,omitempty" xml:"delivery_address,omitempty"`
	// 金融方案信息
	FinanceInfo *AutoFinanceDto `json:"finance_info,omitempty" xml:"finance_info,omitempty"`
	// 是否支付定金
	EarnestPaid bool `json:"earnest_paid,omitempty" xml:"earnest_paid,omitempty"`
	// 是否有效订单
	Valid bool `json:"valid,omitempty" xml:"valid,omitempty"`
}
