package tbtrade

import (
	"sync"
)

// OrderDetailCursorTopDto 结构体
type OrderDetailCursorTopDto struct {
	// 商家昵称
	SellerNick string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
	// 虚拟主订单id
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 虚拟子订单id
	SubOrderId string `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
	// 店铺名称
	ShopName string `json:"shop_name,omitempty" xml:"shop_name,omitempty"`
	// 商品名称（转化item）
	ConvItemName string `json:"conv_item_name,omitempty" xml:"conv_item_name,omitempty"`
	// 广告点击时间
	ClickTime string `json:"click_time,omitempty" xml:"click_time,omitempty"`
	// 针对巨量2.0，图片素材宏参（下发原始素材id）
	Mid1 string `json:"mid1,omitempty" xml:"mid1,omitempty"`
	// 针对巨量2.0，标题素材宏参（下发原始素材id）
	Mid2 string `json:"mid2,omitempty" xml:"mid2,omitempty"`
	// 针对巨量2.0，视频素材宏参（下发原始素材id）
	Mid3 string `json:"mid3,omitempty" xml:"mid3,omitempty"`
	// 同步clickid
	SynClickId string `json:"syn_click_id,omitempty" xml:"syn_click_id,omitempty"`
	// 同步requestid
	SynRequestId string `json:"syn_request_id,omitempty" xml:"syn_request_id,omitempty"`
	// 异步clickid
	AsynClickId string `json:"asyn_click_id,omitempty" xml:"asyn_click_id,omitempty"`
	// 异步requestid
	AsynRequestId string `json:"asyn_request_id,omitempty" xml:"asyn_request_id,omitempty"`
	// 监测唯一id
	MonitorUniqueId string `json:"monitor_unique_id,omitempty" xml:"monitor_unique_id,omitempty"`
	// 转化时间
	ConvTime string `json:"conv_time,omitempty" xml:"conv_time,omitempty"`
	// 商品名称（点击item）
	ItemName string `json:"item_name,omitempty" xml:"item_name,omitempty"`
	// 商品名称（下单item）
	PayItemName string `json:"pay_item_name,omitempty" xml:"pay_item_name,omitempty"`
	// 订单创建时间
	OrderCreateTime string `json:"order_create_time,omitempty" xml:"order_create_time,omitempty"`
	// 订单支付时间
	OrderPayTime string `json:"order_pay_time,omitempty" xml:"order_pay_time,omitempty"`
	// 订单退款时间
	OrderRefundTime string `json:"order_refund_time,omitempty" xml:"order_refund_time,omitempty"`
	// 流量通Pro订单服务费率
	CommissionRate string `json:"commission_rate,omitempty" xml:"commission_rate,omitempty"`
	// 成交商品主图url
	ItemPicUrl string `json:"item_pic_url,omitempty" xml:"item_pic_url,omitempty"`
	// 广告投放位置，仅channel_id=1时返回
	CSite string `json:"c_site,omitempty" xml:"c_site,omitempty"`
	// 字节创意样式，仅channel_id=1时返回
	CType string `json:"c_type,omitempty" xml:"c_type,omitempty"`
	// B站稿件id，仅channel_id=108时返回
	BVid string `json:"b_vid,omitempty" xml:"b_vid,omitempty"`
	// 商家ID
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 品牌ID
	BrandId int64 `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
	// 商品id（点击item）
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 商品id（转化item）
	ConvItemId int64 `json:"conv_item_id,omitempty" xml:"conv_item_id,omitempty"`
	// 转化类型，1-收藏，2-加购
	ConvType int64 `json:"conv_type,omitempty" xml:"conv_type,omitempty"`
	// 待支付金额（单位分）
	DuePayFee int64 `json:"due_pay_fee,omitempty" xml:"due_pay_fee,omitempty"`
	// 渠道(媒体)id，1-巨量2.0
	ChannelId int64 `json:"channel_id,omitempty" xml:"channel_id,omitempty"`
	// 投放账户id
	AdvertiserId int64 `json:"advertiser_id,omitempty" xml:"advertiser_id,omitempty"`
	// 针对巨量2.0，项目id
	ProjectId int64 `json:"project_id,omitempty" xml:"project_id,omitempty"`
	// 针对巨量2.0，广告id
	PromotionId int64 `json:"promotion_id,omitempty" xml:"promotion_id,omitempty"`
	// 计划id
	CampaignId int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
	// 广告id
	AdgroupId int64 `json:"adgroup_id,omitempty" xml:"adgroup_id,omitempty"`
	// 创意id
	CreativeId int64 `json:"creative_id,omitempty" xml:"creative_id,omitempty"`
	// 商品id（下单item）
	PayItemId int64 `json:"pay_item_id,omitempty" xml:"pay_item_id,omitempty"`
	// 订单状态：1-拍下；2-支付；3-退款；4-预售；5-收货
	OrderStatus int64 `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// 支付金额（单位分）
	PayFee int64 `json:"pay_fee,omitempty" xml:"pay_fee,omitempty"`
	// 流量通Pro订单为实际支付金额（不含运费、税费），UDsmart订单与支付金额一致
	ActualPayFee int64 `json:"actual_pay_fee,omitempty" xml:"actual_pay_fee,omitempty"`
	// 订单退款金额（单位分）
	OrderRefundFee int64 `json:"order_refund_fee,omitempty" xml:"order_refund_fee,omitempty"`
	// 是否广告订单:1-是，0-不是
	IsAdOrder int64 `json:"is_ad_order,omitempty" xml:"is_ad_order,omitempty"`
	// 定金金额（单位分）
	PresaleDepositFee int64 `json:"presale_deposit_fee,omitempty" xml:"presale_deposit_fee,omitempty"`
	// 尾款金额（单位分）
	RemainingPaymentFee int64 `json:"remaining_payment_fee,omitempty" xml:"remaining_payment_fee,omitempty"`
	// 是否其他平台归因，0-否，1-是
	IsAttrByOthers int64 `json:"is_attr_by_others,omitempty" xml:"is_attr_by_others,omitempty"`
	// 购买件数
	BuyAmount int64 `json:"buy_amount,omitempty" xml:"buy_amount,omitempty"`
	// 流量通Pro订单成交预估服务费，实际服务费以后台结算为准
	CommissionFee int64 `json:"commission_fee,omitempty" xml:"commission_fee,omitempty"`
}

var poolOrderDetailCursorTopDto = sync.Pool{
	New: func() any {
		return new(OrderDetailCursorTopDto)
	},
}

// GetOrderDetailCursorTopDto() 从对象池中获取OrderDetailCursorTopDto
func GetOrderDetailCursorTopDto() *OrderDetailCursorTopDto {
	return poolOrderDetailCursorTopDto.Get().(*OrderDetailCursorTopDto)
}

// ReleaseOrderDetailCursorTopDto 释放OrderDetailCursorTopDto
func ReleaseOrderDetailCursorTopDto(v *OrderDetailCursorTopDto) {
	v.SellerNick = ""
	v.OrderId = ""
	v.SubOrderId = ""
	v.ShopName = ""
	v.ConvItemName = ""
	v.ClickTime = ""
	v.Mid1 = ""
	v.Mid2 = ""
	v.Mid3 = ""
	v.SynClickId = ""
	v.SynRequestId = ""
	v.AsynClickId = ""
	v.AsynRequestId = ""
	v.MonitorUniqueId = ""
	v.ConvTime = ""
	v.ItemName = ""
	v.PayItemName = ""
	v.OrderCreateTime = ""
	v.OrderPayTime = ""
	v.OrderRefundTime = ""
	v.CommissionRate = ""
	v.ItemPicUrl = ""
	v.CSite = ""
	v.CType = ""
	v.BVid = ""
	v.SellerId = 0
	v.BrandId = 0
	v.ItemId = 0
	v.ConvItemId = 0
	v.ConvType = 0
	v.DuePayFee = 0
	v.ChannelId = 0
	v.AdvertiserId = 0
	v.ProjectId = 0
	v.PromotionId = 0
	v.CampaignId = 0
	v.AdgroupId = 0
	v.CreativeId = 0
	v.PayItemId = 0
	v.OrderStatus = 0
	v.PayFee = 0
	v.ActualPayFee = 0
	v.OrderRefundFee = 0
	v.IsAdOrder = 0
	v.PresaleDepositFee = 0
	v.RemainingPaymentFee = 0
	v.IsAttrByOthers = 0
	v.BuyAmount = 0
	v.CommissionFee = 0
	poolOrderDetailCursorTopDto.Put(v)
}
