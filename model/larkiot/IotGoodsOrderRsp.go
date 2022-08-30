package larkiot

// IotGoodsOrderRsp 结构体
type IotGoodsOrderRsp struct {
	// 订单号
	GoodsOrderId string `json:"goods_order_id,omitempty" xml:"goods_order_id,omitempty"`
	// 影院内码
	CinemaLinkId string `json:"cinema_link_id,omitempty" xml:"cinema_link_id,omitempty"`
	// 外部订单号
	OutGoodsOrderId string `json:"out_goods_order_id,omitempty" xml:"out_goods_order_id,omitempty"`
}
