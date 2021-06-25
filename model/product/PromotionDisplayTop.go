package product

// PromotionDisplayTop 
type PromotionDisplayTop struct {

    // 单品级优惠信息
    PromotionInItem   []PromotionInItem `json:"promotion_in_item,omitempty"`

    // 店铺级优惠信息
    PromotionInShop   []PromotionInShop `json:"promotion_in_shop,omitempty"`

}
