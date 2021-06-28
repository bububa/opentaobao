package iot

// PromotionDisplayTop 
/* model for simplify = false
type PromotionDisplayTop struct {

    // 单品级优惠信息
    
    PromotionInItemList  struct {
        PromotionInItem  []PromotionInItem `json:"promotion_in_item,omitempty"`
    } `json:"promotion_in_item_list,omitempty"`
    

    // 店铺级优惠信息
    
    PromotionInShopList  struct {
        PromotionInShop  []PromotionInShop `json:"promotion_in_shop,omitempty"`
    } `json:"promotion_in_shop_list,omitempty"`
    

}
*/

// PromotionDisplayTop 
type PromotionDisplayTop struct {

    // 单品级优惠信息
    PromotionInItemList   []PromotionInItem `json:"promotion_in_item_list,omitempty"`

    // 店铺级优惠信息
    PromotionInShopList   []PromotionInShop `json:"promotion_in_shop_list,omitempty"`

}
