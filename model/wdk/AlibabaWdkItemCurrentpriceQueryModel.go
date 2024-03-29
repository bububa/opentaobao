package wdk

import (
	"sync"
)

// AlibabaWdkItemCurrentpriceQueryModel 结构体
type AlibabaWdkItemCurrentpriceQueryModel struct {
	// 单品促销，finalPrice对应的促销活动
	ItemPromotionList []PromotionInfoDto `json:"item_promotion_list,omitempty" xml:"item_promotion_list>promotion_info_dto,omitempty"`
	// 商品池促销
	ShopPromotionList []PromotionInfoDto `json:"shop_promotion_list,omitempty" xml:"shop_promotion_list>promotion_info_dto,omitempty"`
	// 会员单品促销活动，memberFinalPrice对应的促销活动
	MemberItemPromotionList []PromotionInfoDto `json:"member_item_promotion_list,omitempty" xml:"member_item_promotion_list>promotion_info_dto,omitempty"`
	// 商家编码
	MerchantCode string `json:"merchant_code,omitempty" xml:"merchant_code,omitempty"`
	// 经营店Code
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 条码
	BarCode string `json:"bar_code,omitempty" xml:"bar_code,omitempty"`
	// 商品名称
	ItemName string `json:"item_name,omitempty" xml:"item_name,omitempty"`
	// 商品编码
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// 渠道，APP／POS
	OrderChannelCode string `json:"order_channel_code,omitempty" xml:"order_channel_code,omitempty"`
	// 会员价，单位为分。      * 对应IC标memberPrice的价格，如果没有对应的标，该价格为空
	MemberPrice int64 `json:"member_price,omitempty" xml:"member_price,omitempty"`
	// 促销执行价格，单位为分
	FinalPrice int64 `json:"final_price,omitempty" xml:"final_price,omitempty"`
	// 买赠活动信息
	MzPromotion *MzPromotionDto `json:"mz_promotion,omitempty" xml:"mz_promotion,omitempty"`
	// 会员促销执行价，单位为分。
	MemberFinalPrice int64 `json:"member_final_price,omitempty" xml:"member_final_price,omitempty"`
	// IC的商品ID
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 基础售价,单位为分
	AuctionPrice int64 `json:"auction_price,omitempty" xml:"auction_price,omitempty"`
	// 渠道店ID
	ShopId int64 `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
}

var poolAlibabaWdkItemCurrentpriceQueryModel = sync.Pool{
	New: func() any {
		return new(AlibabaWdkItemCurrentpriceQueryModel)
	},
}

// GetAlibabaWdkItemCurrentpriceQueryModel() 从对象池中获取AlibabaWdkItemCurrentpriceQueryModel
func GetAlibabaWdkItemCurrentpriceQueryModel() *AlibabaWdkItemCurrentpriceQueryModel {
	return poolAlibabaWdkItemCurrentpriceQueryModel.Get().(*AlibabaWdkItemCurrentpriceQueryModel)
}

// ReleaseAlibabaWdkItemCurrentpriceQueryModel 释放AlibabaWdkItemCurrentpriceQueryModel
func ReleaseAlibabaWdkItemCurrentpriceQueryModel(v *AlibabaWdkItemCurrentpriceQueryModel) {
	v.ItemPromotionList = v.ItemPromotionList[:0]
	v.ShopPromotionList = v.ShopPromotionList[:0]
	v.MemberItemPromotionList = v.MemberItemPromotionList[:0]
	v.MerchantCode = ""
	v.StoreId = ""
	v.BarCode = ""
	v.ItemName = ""
	v.SkuCode = ""
	v.OrderChannelCode = ""
	v.MemberPrice = 0
	v.FinalPrice = 0
	v.MzPromotion = nil
	v.MemberFinalPrice = 0
	v.ItemId = 0
	v.AuctionPrice = 0
	v.ShopId = 0
	poolAlibabaWdkItemCurrentpriceQueryModel.Put(v)
}
