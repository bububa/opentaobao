package wdk

// AlibabaWdkItemCurrentpriceQueryModel 
type AlibabaWdkItemCurrentpriceQueryModel struct {

    // 商家编码
    
    MerchantCode   string `json:"merchant_code,omitempty" xml:"merchant_code,omitempty"`
    

    // 会员价，单位为分。      * 对应IC标memberPrice的价格，如果没有对应的标，该价格为空
    
    MemberPrice   int64 `json:"member_price,omitempty" xml:"member_price,omitempty"`
    

    // 促销执行价格，单位为分
    
    FinalPrice   int64 `json:"final_price,omitempty" xml:"final_price,omitempty"`
    

    // 经营店Code
    
    StoreId   string `json:"store_id,omitempty" xml:"store_id,omitempty"`
    

    // 买赠活动信息
    
    MzPromotion   *MzPromotionDto `json:"mz_promotion,omitempty" xml:"mz_promotion,omitempty"`
    

    // 条码
    
    BarCode   string `json:"bar_code,omitempty" xml:"bar_code,omitempty"`
    

    // 会员促销执行价，单位为分。
    
    MemberFinalPrice   int64 `json:"member_final_price,omitempty" xml:"member_final_price,omitempty"`
    

    // IC的商品ID
    
    ItemId   int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
    

    // 单品促销，finalPrice对应的促销活动
    
    ItemPromotionList   []PromotionInfoDto `json:"item_promotion_list,omitempty" xml:"item_promotion_list,omitempty"`
    

    // 商品名称
    
    ItemName   string `json:"item_name,omitempty" xml:"item_name,omitempty"`
    

    // 基础售价,单位为分
    
    AuctionPrice   int64 `json:"auction_price,omitempty" xml:"auction_price,omitempty"`
    

    // 渠道店ID
    
    ShopId   int64 `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
    

    // 商品池促销
    
    ShopPromotionList   []PromotionInfoDto `json:"shop_promotion_list,omitempty" xml:"shop_promotion_list,omitempty"`
    

    // 会员单品促销活动，memberFinalPrice对应的促销活动
    
    MemberItemPromotionList   []PromotionInfoDto `json:"member_item_promotion_list,omitempty" xml:"member_item_promotion_list,omitempty"`
    

    // 商品编码
    
    SkuCode   string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
    

    // 渠道，APP／POS
    
    OrderChannelCode   string `json:"order_channel_code,omitempty" xml:"order_channel_code,omitempty"`
    

}
