package idle

import (
	"sync"
)

// TenderOrderInfoVo 结构体
type TenderOrderInfoVo struct {
	// 扩展信息
	ChannelData string `json:"channel_data,omitempty" xml:"channel_data,omitempty"`
	// 拍卖买家的用户nick
	BuyerNick string `json:"buyer_nick,omitempty" xml:"buyer_nick,omitempty"`
	// 服务商退货的原因
	ServiceRefundReason string `json:"service_refund_reason,omitempty" xml:"service_refund_reason,omitempty"`
	// 上门回收时间
	ShipTime string `json:"ship_time,omitempty" xml:"ship_time,omitempty"`
	// 渠道
	Channel string `json:"channel,omitempty" xml:"channel,omitempty"`
	// 订单主状态
	OrderStatus string `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// 卖家地址
	SellerAddress string `json:"seller_address,omitempty" xml:"seller_address,omitempty"`
	// 拍卖最终成交价格
	DealPrice string `json:"deal_price,omitempty" xml:"deal_price,omitempty"`
	// 买家区
	BuyerArea string `json:"buyer_area,omitempty" xml:"buyer_area,omitempty"`
	// 服务商评估价格
	AppraisePrice string `json:"appraise_price,omitempty" xml:"appraise_price,omitempty"`
	// 买家地址
	BuyerAddress string `json:"buyer_address,omitempty" xml:"buyer_address,omitempty"`
	// 买家城市信息
	BuyerCity string `json:"buyer_city,omitempty" xml:"buyer_city,omitempty"`
	// 卖家alipay信息
	SellerAlipayUserId string `json:"seller_alipay_user_id,omitempty" xml:"seller_alipay_user_id,omitempty"`
	// 卖家城市信息
	SellerCity string `json:"seller_city,omitempty" xml:"seller_city,omitempty"`
	// appkey
	AppKey string `json:"app_key,omitempty" xml:"app_key,omitempty"`
	// 扩展信息
	Attribute string `json:"attribute,omitempty" xml:"attribute,omitempty"`
	// 服务商绑定的用户id
	ServiceId string `json:"service_id,omitempty" xml:"service_id,omitempty"`
	// 卖家nick
	SellerNick string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
	// 起拍价格
	StartPrice string `json:"start_price,omitempty" xml:"start_price,omitempty"`
	// 买家省份
	BuyerProvince string `json:"buyer_province,omitempty" xml:"buyer_province,omitempty"`
	// 订单子状态
	OrderSubStatus string `json:"order_sub_status,omitempty" xml:"order_sub_status,omitempty"`
	// 卖家的电话信息
	SellerPhone string `json:"seller_phone,omitempty" xml:"seller_phone,omitempty"`
	// 卖家小区信息
	SellerCountry string `json:"seller_country,omitempty" xml:"seller_country,omitempty"`
	// 卖家的alipay账号信息
	SellerAlipayAccount string `json:"seller_alipay_account,omitempty" xml:"seller_alipay_account,omitempty"`
	// 卖家关闭原因
	SellerCloseReason string `json:"seller_close_reason,omitempty" xml:"seller_close_reason,omitempty"`
	// 拍卖的终止价格
	EndPrice string `json:"end_price,omitempty" xml:"end_price,omitempty"`
	// 购买人的真实姓名
	BuyerRealName string `json:"buyer_real_name,omitempty" xml:"buyer_real_name,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 创建环境
	Env string `json:"env,omitempty" xml:"env,omitempty"`
	// 卖家真实姓名
	SellerRealName string `json:"seller_real_name,omitempty" xml:"seller_real_name,omitempty"`
	// 服务商关闭的原因
	ServiceCloseReason string `json:"service_close_reason,omitempty" xml:"service_close_reason,omitempty"`
	// 卖家的地址
	SellerArea string `json:"seller_area,omitempty" xml:"seller_area,omitempty"`
	// 买家的地址信息
	BuyerCountry string `json:"buyer_country,omitempty" xml:"buyer_country,omitempty"`
	// 回收支付类型
	IdlePayType string `json:"idle_pay_type,omitempty" xml:"idle_pay_type,omitempty"`
	// 买家的电话信息
	BuyerPhone string `json:"buyer_phone,omitempty" xml:"buyer_phone,omitempty"`
	// 服务商的公司名称
	ServiceNick string `json:"service_nick,omitempty" xml:"service_nick,omitempty"`
	// 下单场景
	SceneType string `json:"scene_type,omitempty" xml:"scene_type,omitempty"`
	// 订单id2
	BizOrderId string `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 估价id
	AppraiseId string `json:"appraise_id,omitempty" xml:"appraise_id,omitempty"`
	// 服务商最终成交价格
	ServiceDealPrice string `json:"service_deal_price,omitempty" xml:"service_deal_price,omitempty"`
	// 上门类型
	ShipType string `json:"ship_type,omitempty" xml:"ship_type,omitempty"`
	// spuId
	SpuId string `json:"spu_id,omitempty" xml:"spu_id,omitempty"`
	// 卖家的退回原因
	SellerRefundReason string `json:"seller_refund_reason,omitempty" xml:"seller_refund_reason,omitempty"`
	// 卖家的地址信息
	SellerProvince string `json:"seller_province,omitempty" xml:"seller_province,omitempty"`
	// 拍场订单Id
	TenderOrderId string `json:"tender_order_id,omitempty" xml:"tender_order_id,omitempty"`
	// 服务费(分)
	ServiceFee string `json:"service_fee,omitempty" xml:"service_fee,omitempty"`
	// 买家支付宝流水号
	BuyerOutPayId string `json:"buyer_out_pay_id,omitempty" xml:"buyer_out_pay_id,omitempty"`
}

var poolTenderOrderInfoVo = sync.Pool{
	New: func() any {
		return new(TenderOrderInfoVo)
	},
}

// GetTenderOrderInfoVo() 从对象池中获取TenderOrderInfoVo
func GetTenderOrderInfoVo() *TenderOrderInfoVo {
	return poolTenderOrderInfoVo.Get().(*TenderOrderInfoVo)
}

// ReleaseTenderOrderInfoVo 释放TenderOrderInfoVo
func ReleaseTenderOrderInfoVo(v *TenderOrderInfoVo) {
	v.ChannelData = ""
	v.BuyerNick = ""
	v.ServiceRefundReason = ""
	v.ShipTime = ""
	v.Channel = ""
	v.OrderStatus = ""
	v.SellerAddress = ""
	v.DealPrice = ""
	v.BuyerArea = ""
	v.AppraisePrice = ""
	v.BuyerAddress = ""
	v.BuyerCity = ""
	v.SellerAlipayUserId = ""
	v.SellerCity = ""
	v.AppKey = ""
	v.Attribute = ""
	v.ServiceId = ""
	v.SellerNick = ""
	v.StartPrice = ""
	v.BuyerProvince = ""
	v.OrderSubStatus = ""
	v.SellerPhone = ""
	v.SellerCountry = ""
	v.SellerAlipayAccount = ""
	v.SellerCloseReason = ""
	v.EndPrice = ""
	v.BuyerRealName = ""
	v.GmtCreate = ""
	v.Env = ""
	v.SellerRealName = ""
	v.ServiceCloseReason = ""
	v.SellerArea = ""
	v.BuyerCountry = ""
	v.IdlePayType = ""
	v.BuyerPhone = ""
	v.ServiceNick = ""
	v.SceneType = ""
	v.BizOrderId = ""
	v.AppraiseId = ""
	v.ServiceDealPrice = ""
	v.ShipType = ""
	v.SpuId = ""
	v.SellerRefundReason = ""
	v.SellerProvince = ""
	v.TenderOrderId = ""
	v.ServiceFee = ""
	v.BuyerOutPayId = ""
	poolTenderOrderInfoVo.Put(v)
}
