package wdk

// ItemBuyGiftSku 
type ItemBuyGiftSku struct {
    // 限购信息
    LimitInfo   *LimitInfo `json:"limit_info,omitempty" xml:"limit_info,omitempty"`
    // 商品的skuCode
    SkuCode   string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
    // 买赠门槛数量
    BuyNum   int64 `json:"buy_num,omitempty" xml:"buy_num,omitempty"`
    // 赠品的skuCode
    GiftSkuCode   string `json:"gift_sku_code,omitempty" xml:"gift_sku_code,omitempty"`
    // 淘宝item和shop的对应关系， k-itemId, v-shopId
    ItemShopRelation   string `json:"item_shop_relation,omitempty" xml:"item_shop_relation,omitempty"`
    // 主商品名称
    SkuName   string `json:"sku_name,omitempty" xml:"sku_name,omitempty"`
    // 赠品名称
    GiftSkuName   string `json:"gift_sku_name,omitempty" xml:"gift_sku_name,omitempty"`
}
