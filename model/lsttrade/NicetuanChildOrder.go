package lsttrade

// NicetuanChildOrder 
type NicetuanChildOrder struct {

    // 规格名称
    
    SkuName   string `json:"sku_name,omitempty" xml:"sku_name,omitempty"`
    

    // 子订单订单金额
    
    Amount   string `json:"amount,omitempty" xml:"amount,omitempty"`
    

    // 商品数量
    
    Quantity   int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
    

    // 子订单佣金金额
    
    RebateAmount   string `json:"rebate_amount,omitempty" xml:"rebate_amount,omitempty"`
    

    // 是否零售通商品
    
    LstYn   string `json:"lst_yn,omitempty" xml:"lst_yn,omitempty"`
    

    // 商品图片
    
    MerchandiseCoverImage   string `json:"merchandise_cover_image,omitempty" xml:"merchandise_cover_image,omitempty"`
    

    // 预计发货日期
    
    DeliveryDay   string `json:"delivery_day,omitempty" xml:"delivery_day,omitempty"`
    

    // 子订单号
    
    SubOrderCode   string `json:"sub_order_code,omitempty" xml:"sub_order_code,omitempty"`
    

    // 商品名称
    
    MerchandiseName   string `json:"merchandise_name,omitempty" xml:"merchandise_name,omitempty"`
    

    // skuID
    
    SkuId   string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
    

}
