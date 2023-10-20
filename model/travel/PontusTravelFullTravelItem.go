package travel

import (
	"sync"
)

// PontusTravelFullTravelItem 结构体
type PontusTravelFullTravelItem struct {
	// 预定规则
	BookingRules []PontusTravelBookingRuleInfo `json:"booking_rules,omitempty" xml:"booking_rules>pontus_travel_booking_rule_info,omitempty"`
	// 行程信息
	Itineraries []PontusTravelItemItineraryInfo `json:"itineraries,omitempty" xml:"itineraries>pontus_travel_item_itinerary_info,omitempty"`
	// sku信息
	SkuInfos []PontusTravelItemSkuInfo `json:"sku_infos,omitempty" xml:"sku_infos>pontus_travel_item_sku_info,omitempty"`
	// 产品亮点
	HighLights []HighLights `json:"high_lights,omitempty" xml:"high_lights>high_lights,omitempty"`
	// 创建时间
	Created string `json:"created,omitempty" xml:"created,omitempty"`
	// 修改时间
	Modified string `json:"modified,omitempty" xml:"modified,omitempty"`
	// 卖家昵称
	SellerNick string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
	// 商品特征数据
	Features string `json:"features,omitempty" xml:"features,omitempty"`
	// 新版行程信息
	RefTrip string `json:"ref_trip,omitempty" xml:"ref_trip,omitempty"`
	// 商品基本信息
	BaseInfo *PontusTravelItemBaseInfo `json:"base_info,omitempty" xml:"base_info,omitempty"`
	// 自由行相关信息
	FreedomItemExt *PontusTravelFreedomItemExt `json:"freedom_item_ext,omitempty" xml:"freedom_item_ext,omitempty"`
	// 跟团游相关信息
	GroupItemExt *PontusTravelGroupItemExt `json:"group_item_ext,omitempty" xml:"group_item_ext,omitempty"`
	// 商品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 商品状态。0,1正常;-1:用户删除;-2:用户下架;-3 小二下架;-4 小二删除;-5 从未上架;-9 被处罚
	ItemStatus int64 `json:"item_status,omitempty" xml:"item_status,omitempty"`
	// 商品类型
	ItemType int64 `json:"item_type,omitempty" xml:"item_type,omitempty"`
	// 退改规则信息
	RefundInfo *PontusTravelItemRefundInfo `json:"refund_info,omitempty" xml:"refund_info,omitempty"`
	// 销售属性信息
	SaleInfo *PontusTravelItemSaleInfo `json:"sale_info,omitempty" xml:"sale_info,omitempty"`
	// 卖家id
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 同城活动商品相关信息
	TcwlItemExt *TcwlItemExt `json:"tcwl_item_ext,omitempty" xml:"tcwl_item_ext,omitempty"`
	// 航旅度假TOP API3.0 邮轮扩展信息结构
	CruiseItemExt *CruiseItemExt `json:"cruise_item_ext,omitempty" xml:"cruise_item_ext,omitempty"`
}

var poolPontusTravelFullTravelItem = sync.Pool{
	New: func() any {
		return new(PontusTravelFullTravelItem)
	},
}

// GetPontusTravelFullTravelItem() 从对象池中获取PontusTravelFullTravelItem
func GetPontusTravelFullTravelItem() *PontusTravelFullTravelItem {
	return poolPontusTravelFullTravelItem.Get().(*PontusTravelFullTravelItem)
}

// ReleasePontusTravelFullTravelItem 释放PontusTravelFullTravelItem
func ReleasePontusTravelFullTravelItem(v *PontusTravelFullTravelItem) {
	v.BookingRules = v.BookingRules[:0]
	v.Itineraries = v.Itineraries[:0]
	v.SkuInfos = v.SkuInfos[:0]
	v.HighLights = v.HighLights[:0]
	v.Created = ""
	v.Modified = ""
	v.SellerNick = ""
	v.Features = ""
	v.RefTrip = ""
	v.BaseInfo = nil
	v.FreedomItemExt = nil
	v.GroupItemExt = nil
	v.ItemId = 0
	v.ItemStatus = 0
	v.ItemType = 0
	v.RefundInfo = nil
	v.SaleInfo = nil
	v.SellerId = 0
	v.TcwlItemExt = nil
	v.CruiseItemExt = nil
	poolPontusTravelFullTravelItem.Put(v)
}
