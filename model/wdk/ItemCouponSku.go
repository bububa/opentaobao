package wdk

// ItemCouponSku 
/* model for simplify = false
type ItemCouponSku struct {

    // 商品的skuCode
    
    SkuCode   string `json:"sku_code,omitempty"`
    

    // 商品名称
    
    SkuName   string `json:"sku_name,omitempty"`
    

    // 淘宝item和shop的对应关系， k-itemId, v-shopId
    
    ItemShopRelation   string `json:"item_shop_relation,omitempty"`
    

}
*/

// ItemCouponSku 
type ItemCouponSku struct {

    // 商品的skuCode
    SkuCode   string `json:"sku_code,omitempty"`

    // 商品名称
    SkuName   string `json:"sku_name,omitempty"`

    // 淘宝item和shop的对应关系， k-itemId, v-shopId
    ItemShopRelation   string `json:"item_shop_relation,omitempty"`

}
