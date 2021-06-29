package idleisv

// AppraiseIsvOrderDto 
type AppraiseIsvOrderDto struct {

    // 订单号
    
    BizOrderId   int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
    

    // 订单创建时间,时间戳,毫秒
    
    CreateTime   int64 `json:"create_time,omitempty" xml:"create_time,omitempty"`
    

    // 0:未知状态、1：订单已创建、2：订单已付款、3：已发货、4：交易成功、5：已退款、6：交易关闭
    
    OrderStatus   int64 `json:"order_status,omitempty" xml:"order_status,omitempty"`
    

    // 订单下单付款时间,时间戳,毫秒
    
    PayTime   int64 `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
    

    // 订单发货时间,时间戳,毫秒
    
    ShipTime   int64 `json:"ship_time,omitempty" xml:"ship_time,omitempty"`
    

    // 商品购买数量
    
    BuyAmount   int64 `json:"buy_amount,omitempty" xml:"buy_amount,omitempty"`
    

    // 买家收货地址
    
    BuyerAddress   *AppraiseIsvAddressDto `json:"buyer_address,omitempty" xml:"buyer_address,omitempty"`
    

    // 买家昵称
    
    BuyerNick   string `json:"buyer_nick,omitempty" xml:"buyer_nick,omitempty"`
    

    // 商品信息
    
    Item   *AppraiseIsvItemDto `json:"item,omitempty" xml:"item,omitempty"`
    

    // 实付金额, 单位分
    
    Payment   int64 `json:"payment,omitempty" xml:"payment,omitempty"`
    

    // 邮费
    
    PostFee   int64 `json:"post_fee,omitempty" xml:"post_fee,omitempty"`
    

    // 卖家昵称
    
    SellerNick   string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
    

    // 订单完结时间,时间戳,毫秒
    
    EndTime   int64 `json:"end_time,omitempty" xml:"end_time,omitempty"`
    

    // sku信息（格式： skuId|属性名:属性值;属性名:属性值）
    
    Sku   string `json:"sku,omitempty" xml:"sku,omitempty"`
    

    // 物流信息
    
    Logistics   *AppraiseIsvLogisticsDto `json:"logistics,omitempty" xml:"logistics,omitempty"`
    

}
