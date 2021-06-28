package product

// PromotionDisplayTop 
/* model for simplify = false
type PromotionDisplayTop struct {

    // 单品级优惠信息
    
    PromotionInItem  struct {
        PromotionInItem  []PromotionInItem `json:"promotion_in_item,omitempty"`
    } `json:"promotion_in_item,omitempty"`
    

    // 店铺级优惠信息
    
    PromotionInShop  struct {
        PromotionInShop  []PromotionInShop `json:"promotion_in_shop,omitempty"`
    } `json:"promotion_in_shop,omitempty"`
    

}
*/

// PromotionDisplayTop 
type PromotionDisplayTop struct {

    // 单品级优惠信息
    PromotionInItem   []PromotionInItem `json:"promotion_in_item,omitempty"`

    // 店铺级优惠信息
    PromotionInShop   []PromotionInShop `json:"promotion_in_shop,omitempty"`

}
