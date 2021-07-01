package alihealth2

// TaobaoTradeDrugGetResult 结构体
type TaobaoTradeDrugGetResult struct {
	// 订单列表
	ResultList []TaobaoTradeDrugGetResult `json:"result_list,omitempty" xml:"result_list>taobao_trade_drug_get_result,omitempty"`
	// 返回记录数
	ResultSize int64 `json:"result_size,omitempty" xml:"result_size,omitempty"`
	// 总记录数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 送达时间要求_起始
	StartDeliveryTime string `json:"start_delivery_time,omitempty" xml:"start_delivery_time,omitempty"`
	// 用户地址对象
	UserAddress *UserAddress `json:"user_address,omitempty" xml:"user_address,omitempty"`
	// 下单商品列表
	GoodsList []OrderGoods `json:"goods_list,omitempty" xml:"goods_list>order_goods,omitempty"`
	// 分店ID
	StoreId int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 订单号
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 送货地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 配送费
	DeliveryPay string `json:"delivery_pay,omitempty" xml:"delivery_pay,omitempty"`
	// 送达时间要求_结束
	EndDeliveryTime string `json:"end_delivery_time,omitempty" xml:"end_delivery_time,omitempty"`
	// 创建时间
	CreateTime string `json:"create_time,omitempty" xml:"create_time,omitempty"`
	// 总支付金额
	TotalPay string `json:"total_pay,omitempty" xml:"total_pay,omitempty"`
	// 店铺名称
	StoreName string `json:"store_name,omitempty" xml:"store_name,omitempty"`
	// 店铺联系电话
	StoreContactPhone string `json:"store_contact_phone,omitempty" xml:"store_contact_phone,omitempty"`
	// 用户下单时的备注信息
	Note string `json:"note,omitempty" xml:"note,omitempty"`
	// 用户下单方式(PC/APP)
	From int64 `json:"from,omitempty" xml:"from,omitempty"`
	// 支付宝流水号
	AlipayStreamId string `json:"alipay_stream_id,omitempty" xml:"alipay_stream_id,omitempty"`
}
