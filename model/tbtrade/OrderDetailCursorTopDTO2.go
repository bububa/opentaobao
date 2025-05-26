package tbtrade

import (
	"sync"
)

// OrderDetailCursorTopDTO2 结构体
type OrderDetailCursorTopDTO2 struct {
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
}

var poolOrderDetailCursorTopDTO2 = sync.Pool{
	New: func() any {
		return new(OrderDetailCursorTopDTO2)
	},
}

// GetOrderDetailCursorTopDTO2() 从对象池中获取OrderDetailCursorTopDTO2
func GetOrderDetailCursorTopDTO2() *OrderDetailCursorTopDTO2 {
	return poolOrderDetailCursorTopDTO2.Get().(*OrderDetailCursorTopDTO2)
}

// ReleaseOrderDetailCursorTopDTO2 释放OrderDetailCursorTopDTO2
func ReleaseOrderDetailCursorTopDTO2(v *OrderDetailCursorTopDTO2) {
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
	poolOrderDetailCursorTopDTO2.Put(v)
}
