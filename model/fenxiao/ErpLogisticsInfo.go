package fenxiao

// ErpLogisticsInfo 结构体
type ErpLogisticsInfo struct {
	// 发货类型 CN=菜鸟发货,SC的商家仓发货
	ConsignType string `json:"consign_type,omitempty" xml:"consign_type,omitempty"`
	// 商品的最小库存单位Sku的id
	SkuId string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 子订单类型:标示该子交易单来源交易，还是BMS增加的，枚举值(00=交易，10=BMS绑定)
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 如是菜鸟仓，则将菜鸟仓的区域仓code进行填充，如是商家仓发货则填充SC
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// 货品仓储code
	ItemCode string `json:"item_code,omitempty" xml:"item_code,omitempty"`
	// 货品仓储id
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 组合商品ID
	CombineItemId string `json:"combine_item_id,omitempty" xml:"combine_item_id,omitempty"`
	// 组合商品Code
	CombineItemCode string `json:"combine_item_code,omitempty" xml:"combine_item_code,omitempty"`
	// 货品条码
	BarCode string `json:"bar_code,omitempty" xml:"bar_code,omitempty"`
	// 推荐配送公司编码
	DeliveryCps string `json:"delivery_cps,omitempty" xml:"delivery_cps,omitempty"`
	// 仓商家编码
	BizStoreCode string `json:"biz_store_code,omitempty" xml:"biz_store_code,omitempty"`
	// 推荐配送erp编码
	BizDeliveryCode string `json:"biz_delivery_code,omitempty" xml:"biz_delivery_code,omitempty"`
	// 仓配建议类型
	BizSdType string `json:"biz_sd_type,omitempty" xml:"biz_sd_type,omitempty"`
	// 预计发货地址-最小行政地址编码
	SendDivisionCode string `json:"send_division_code,omitempty" xml:"send_division_code,omitempty"`
	// 预计发货地址-文案描述-国家
	SendCountry string `json:"send_country,omitempty" xml:"send_country,omitempty"`
	// 预计发货地址-文案描述-省份
	SendState string `json:"send_state,omitempty" xml:"send_state,omitempty"`
	// 预计发货地址-文案描述-城市
	SendCity string `json:"send_city,omitempty" xml:"send_city,omitempty"`
	// 预计发货地址-文案描述-地区
	SendDistrict string `json:"send_district,omitempty" xml:"send_district,omitempty"`
	// 预计发货地址-文案描述-街道/镇
	SendTown string `json:"send_town,omitempty" xml:"send_town,omitempty"`
	// CP黑名单，逗号分隔
	BlackDeliveryCps string `json:"black_delivery_cps,omitempty" xml:"black_delivery_cps,omitempty"`
	// CP白名单，逗号分隔
	WhiteDeliveryCps string `json:"white_delivery_cps,omitempty" xml:"white_delivery_cps,omitempty"`
	// 商品数字编号
	NumIid int64 `json:"num_iid,omitempty" xml:"num_iid,omitempty"`
	// 应发数量
	NeedConsignNum int64 `json:"need_consign_num,omitempty" xml:"need_consign_num,omitempty"`
	// 采购单子单号
	SubOrderId int64 `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
	// 采购单号
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 组合货品比例
	ItemRatio int64 `json:"item_ratio,omitempty" xml:"item_ratio,omitempty"`
}
