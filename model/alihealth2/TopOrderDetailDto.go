package alihealth2

// TopOrderDetailDto 结构体
type TopOrderDetailDto struct {
	// 孔雀翎订单号
	CepOrderId int64 `json:"cep_order_id,omitempty" xml:"cep_order_id,omitempty"`
	// 孔雀翎店铺编码
	CepShopCode string `json:"cep_shop_code,omitempty" xml:"cep_shop_code,omitempty"`
	// 买家姓名
	BuyerName string `json:"buyer_name,omitempty" xml:"buyer_name,omitempty"`
	// 买家电话
	BuyerPhone string `json:"buyer_phone,omitempty" xml:"buyer_phone,omitempty"`
	// 收货地址，省
	AddrProvince string `json:"addr_province,omitempty" xml:"addr_province,omitempty"`
	// 收货地址，市
	AddrCity string `json:"addr_city,omitempty" xml:"addr_city,omitempty"`
	// 收货地址，区
	AddrDistrict string `json:"addr_district,omitempty" xml:"addr_district,omitempty"`
	// 收货地址，详细信息
	AddrDetail string `json:"addr_detail,omitempty" xml:"addr_detail,omitempty"`
	// 收货地址，经度，高德
	AddrLongitude int64 `json:"addr_longitude,omitempty" xml:"addr_longitude,omitempty"`
	// 收货地址，纬度，高德
	AddrLatitude int64 `json:"addr_latitude,omitempty" xml:"addr_latitude,omitempty"`
	// 订单状态（1已支付；2已接单；3已完成；4已取消）
	OrderStatus int64 `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// 退款状态（0未申请退款，默认值；1已申请退款；2同意退款；3拒绝退款；4已取消退款）
	RefundStatus int64 `json:"refund_status,omitempty" xml:"refund_status,omitempty"`
	// 物流状态（0未申请退款，默认值；1已申请退款；2同意退款；3拒绝退款；4已取消退款）
	DeliveryStatus int64 `json:"delivery_status,omitempty" xml:"delivery_status,omitempty"`
	// 期望送达时间
	ExpectTime string `json:"expect_time,omitempty" xml:"expect_time,omitempty"`
	// 物流单号
	ExpressNumber string `json:"express_number,omitempty" xml:"express_number,omitempty"`
	// 物流名称
	ExpressName string `json:"express_name,omitempty" xml:"express_name,omitempty"`
	// 物流编码
	ExpressCode string `json:"express_code,omitempty" xml:"express_code,omitempty"`
	// 商品集合
	Items []TopOrderGoodsDto `json:"items,omitempty" xml:"items>top_order_goods_dto,omitempty"`
	// o2o订单号
	O2oOrderNo string `json:"o2o_order_no,omitempty" xml:"o2o_order_no,omitempty"`
}
