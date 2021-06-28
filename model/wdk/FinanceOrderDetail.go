package wdk

// FinanceOrderDetail 
/* model for simplify = false
type FinanceOrderDetail struct {

    // 币种
    
    Currency   string `json:"currency,omitempty"`
    

    // 未税销售净额
    
    UntaxSaleTotalAmount   int64 `json:"untax_sale_total_amount,omitempty"`
    

    // 含税销售净额
    
    SaleTotalAmount   int64 `json:"sale_total_amount,omitempty"`
    

    // 未税优惠金额
    
    UntaxDiscountAmount   int64 `json:"untax_discount_amount,omitempty"`
    

    // 含税优惠金额
    
    DiscountAmount   int64 `json:"discount_amount,omitempty"`
    

    // 未税金额
    
    UntaxAmount   int64 `json:"untax_amount,omitempty"`
    

    // 含税金额
    
    Amount   int64 `json:"amount,omitempty"`
    

    // 含税商品单价（元）
    
    UnitPrice   int64 `json:"unit_price,omitempty"`
    

    // 交易数量
    
    Quantity   int64 `json:"quantity,omitempty"`
    

    // 税率
    
    TaxRate   string `json:"tax_rate,omitempty"`
    

    // 销售渠道
    
    SaleChannel   string `json:"sale_channel,omitempty"`
    

    // 销售来源
    
    SaleSource   string `json:"sale_source,omitempty"`
    

    // 交易类型
    
    TradeType   string `json:"trade_type,omitempty"`
    

    // 商品名称
    
    SkuName   string `json:"sku_name,omitempty"`
    

    // 商品编码
    
    SkuCode   string `json:"sku_code,omitempty"`
    

    // 业务主订单id
    
    PTradeId   string `json:"p_trade_id,omitempty"`
    

    // 门店名称
    
    ShopName   string `json:"shop_name,omitempty"`
    

    // 门店编码
    
    ShopCode   string `json:"shop_code,omitempty"`
    

    // 交易时间，用户实际下单时间，格式：HH:mm:ss
    
    TradeTime   string `json:"trade_time,omitempty"`
    

    // 业务日期，用户实际下单日期，格式：yyyyMMdd
    
    BizDate   string `json:"biz_date,omitempty"`
    

    // 业务主键
    
    BizUk   string `json:"biz_uk,omitempty"`
    

    // 交易类型编码      * 88 - 销售      * 99 - 退款
    
    TradeTypeCode   int64 `json:"trade_type_code,omitempty"`
    

}
*/

// FinanceOrderDetail 
type FinanceOrderDetail struct {

    // 币种
    Currency   string `json:"currency,omitempty"`

    // 未税销售净额
    UntaxSaleTotalAmount   int64 `json:"untax_sale_total_amount,omitempty"`

    // 含税销售净额
    SaleTotalAmount   int64 `json:"sale_total_amount,omitempty"`

    // 未税优惠金额
    UntaxDiscountAmount   int64 `json:"untax_discount_amount,omitempty"`

    // 含税优惠金额
    DiscountAmount   int64 `json:"discount_amount,omitempty"`

    // 未税金额
    UntaxAmount   int64 `json:"untax_amount,omitempty"`

    // 含税金额
    Amount   int64 `json:"amount,omitempty"`

    // 含税商品单价（元）
    UnitPrice   int64 `json:"unit_price,omitempty"`

    // 交易数量
    Quantity   int64 `json:"quantity,omitempty"`

    // 税率
    TaxRate   string `json:"tax_rate,omitempty"`

    // 销售渠道
    SaleChannel   string `json:"sale_channel,omitempty"`

    // 销售来源
    SaleSource   string `json:"sale_source,omitempty"`

    // 交易类型
    TradeType   string `json:"trade_type,omitempty"`

    // 商品名称
    SkuName   string `json:"sku_name,omitempty"`

    // 商品编码
    SkuCode   string `json:"sku_code,omitempty"`

    // 业务主订单id
    PTradeId   string `json:"p_trade_id,omitempty"`

    // 门店名称
    ShopName   string `json:"shop_name,omitempty"`

    // 门店编码
    ShopCode   string `json:"shop_code,omitempty"`

    // 交易时间，用户实际下单时间，格式：HH:mm:ss
    TradeTime   string `json:"trade_time,omitempty"`

    // 业务日期，用户实际下单日期，格式：yyyyMMdd
    BizDate   string `json:"biz_date,omitempty"`

    // 业务主键
    BizUk   string `json:"biz_uk,omitempty"`

    // 交易类型编码      * 88 - 销售      * 99 - 退款
    TradeTypeCode   int64 `json:"trade_type_code,omitempty"`

}
