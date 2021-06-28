package trade

// FastBuyPosPayRequest 
/* model for simplify = false
type FastBuyPosPayRequest struct {

    // 外部支付宝交易号
    
    AlipayTradeId   string `json:"alipay_trade_id,omitempty"`
    

    // 外部唯一订单号
    
    OutOrderId   string `json:"out_order_id,omitempty"`
    

    // 经营店id
    
    StoreId   string `json:"store_id,omitempty"`
    

    // 外部门店编码
    
    OutShopCode   string `json:"out_shop_code,omitempty"`
    

    // 实际支付金额
    
    PayFee   int64 `json:"pay_fee,omitempty"`
    

}
*/

// FastBuyPosPayRequest 
type FastBuyPosPayRequest struct {

    // 外部支付宝交易号
    AlipayTradeId   string `json:"alipay_trade_id,omitempty"`

    // 外部唯一订单号
    OutOrderId   string `json:"out_order_id,omitempty"`

    // 经营店id
    StoreId   string `json:"store_id,omitempty"`

    // 外部门店编码
    OutShopCode   string `json:"out_shop_code,omitempty"`

    // 实际支付金额
    PayFee   int64 `json:"pay_fee,omitempty"`

}
