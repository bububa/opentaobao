package wdk

// ChannelSkuDo 
/* model for simplify = false
type ChannelSkuDo struct {

    // 状态（用来判断是否可售；1-正常）
    
    LifeStatus   string `json:"life_status,omitempty"`
    

    // 商品编码
    
    SkuCode   string `json:"sku_code,omitempty"`
    

    // 销售价格，单位元
    
    SalePrice   string `json:"sale_price,omitempty"`
    

    // 是否可退：0不可退，1可退，默认不可退
    
    ReturnFlag   string `json:"return_flag,omitempty"`
    

    // 最小销售量
    
    SaleMinimum   int64 `json:"sale_minimum,omitempty"`
    

    // 销售单位
    
    SaleUnit   string `json:"sale_unit,omitempty"`
    

    // 销售规格（如果不填，默认为sale_minimum字段值）
    
    SaleSpec   string `json:"sale_spec,omitempty"`
    

    // 门店或DC编码
    
    OuCode   string `json:"ou_code,omitempty"`
    

    // 出货仓（默认为店仓一体，仓编码为店编码）
    
    DeliverWarehouse   string `json:"deliver_warehouse,omitempty"`
    

    // 是否支持先销后采，默认false
    
    SaleBeforePurchase   bool `json:"sale_before_purchase,omitempty"`
    

    // 原始供应商编码
    
    OriginalSupplierNo   string `json:"original_supplier_no,omitempty"`
    

    // 渠道（默认-1）
    
    ChannelCode   int64 `json:"channel_code,omitempty"`
    

    // 是否toB渠道（默认true）
    
    ToBChannel   bool `json:"to_b_channel,omitempty"`
    

    // 时间戳，时间的毫秒数
    
    TimeStamp   int64 `json:"time_stamp,omitempty"`
    

    // 客户商家编码
    
    CustomerMerchantCode   string `json:"customer_merchant_code,omitempty"`
    

    // 渠道店id
    
    ShopId   string `json:"shop_id,omitempty"`
    

    // 是否在线上露出，0不露出，1露出
    
    OnlineSaleFlag   int64 `json:"online_sale_flag,omitempty"`
    

}
*/

// ChannelSkuDo 
type ChannelSkuDo struct {

    // 状态（用来判断是否可售；1-正常）
    LifeStatus   string `json:"life_status,omitempty"`

    // 商品编码
    SkuCode   string `json:"sku_code,omitempty"`

    // 销售价格，单位元
    SalePrice   string `json:"sale_price,omitempty"`

    // 是否可退：0不可退，1可退，默认不可退
    ReturnFlag   string `json:"return_flag,omitempty"`

    // 最小销售量
    SaleMinimum   int64 `json:"sale_minimum,omitempty"`

    // 销售单位
    SaleUnit   string `json:"sale_unit,omitempty"`

    // 销售规格（如果不填，默认为sale_minimum字段值）
    SaleSpec   string `json:"sale_spec,omitempty"`

    // 门店或DC编码
    OuCode   string `json:"ou_code,omitempty"`

    // 出货仓（默认为店仓一体，仓编码为店编码）
    DeliverWarehouse   string `json:"deliver_warehouse,omitempty"`

    // 是否支持先销后采，默认false
    SaleBeforePurchase   bool `json:"sale_before_purchase,omitempty"`

    // 原始供应商编码
    OriginalSupplierNo   string `json:"original_supplier_no,omitempty"`

    // 渠道（默认-1）
    ChannelCode   int64 `json:"channel_code,omitempty"`

    // 是否toB渠道（默认true）
    ToBChannel   bool `json:"to_b_channel,omitempty"`

    // 时间戳，时间的毫秒数
    TimeStamp   int64 `json:"time_stamp,omitempty"`

    // 客户商家编码
    CustomerMerchantCode   string `json:"customer_merchant_code,omitempty"`

    // 渠道店id
    ShopId   string `json:"shop_id,omitempty"`

    // 是否在线上露出，0不露出，1露出
    OnlineSaleFlag   int64 `json:"online_sale_flag,omitempty"`

}
