package iot

// PromotionDisplayTop 结构体
type PromotionDisplayTop struct {
	// 单品级优惠信息
	PromotionInItemList []PromotionInItem `json:"promotion_in_item_list,omitempty" xml:"promotion_in_item_list>promotion_in_item,omitempty"`
	// 店铺级优惠信息
	PromotionInShopList []PromotionInShop `json:"promotion_in_shop_list,omitempty" xml:"promotion_in_shop_list>promotion_in_shop,omitempty"`
}
