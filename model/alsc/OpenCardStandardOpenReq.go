package alsc

// OpenCardStandardOpenReq 结构体
type OpenCardStandardOpenReq struct {
	// 品牌ID,与outBrandId不能同时为空
	BrandId string `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
	// 顾客ID
	CustomerId string `json:"customer_id,omitempty" xml:"customer_id,omitempty"`
	// 操作人ID
	OperatorId string `json:"operator_id,omitempty" xml:"operator_id,omitempty"`
	// 操作人名称
	OperatorName string `json:"operator_name,omitempty" xml:"operator_name,omitempty"`
	// 物理卡号，与卡模板ID不能同时为空
	PhysicalCardId string `json:"physical_card_id,omitempty" xml:"physical_card_id,omitempty"`
	// 门店ID，与outShopId不能同时为空
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 外部品牌ID，与brandId不能同时为空
	OutBrandId string `json:"out_brand_id,omitempty" xml:"out_brand_id,omitempty"`
	// 请求UUID ,用于做幂等
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 外部门店ID,与shopId不能同时为空
	OutShopId string `json:"out_shop_id,omitempty" xml:"out_shop_id,omitempty"`
	// 卡模板ID，与物理卡号不能同时为空
	CardTemplateId string `json:"card_template_id,omitempty" xml:"card_template_id,omitempty"`
	// 是否激活
	Active bool `json:"active,omitempty" xml:"active,omitempty"`
}
