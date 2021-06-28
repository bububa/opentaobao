package wdk

// ChannelSkuQueryDo 
/* model for simplify = false
type ChannelSkuQueryDo struct {

    // skucode集合
    
    SkuCodes  struct {
        String  []string `json:"string,omitempty"`
    } `json:"sku_codes,omitempty"`
    

    // 门店或DC编码
    
    OuCode   string `json:"ou_code,omitempty"`
    

    // 渠道编码（默认-1）
    
    ChannelCode   int64 `json:"channel_code,omitempty"`
    

    // 出货仓（默认为店仓一体，仓编码为店编码）
    
    DeliverWarehouse   string `json:"deliver_warehouse,omitempty"`
    

    // 客户商家编码
    
    CustomerMerchantCode   string `json:"customer_merchant_code,omitempty"`
    

}
*/

// ChannelSkuQueryDo 
type ChannelSkuQueryDo struct {

    // skucode集合
    SkuCodes   []string `json:"sku_codes,omitempty"`

    // 门店或DC编码
    OuCode   string `json:"ou_code,omitempty"`

    // 渠道编码（默认-1）
    ChannelCode   int64 `json:"channel_code,omitempty"`

    // 出货仓（默认为店仓一体，仓编码为店编码）
    DeliverWarehouse   string `json:"deliver_warehouse,omitempty"`

    // 客户商家编码
    CustomerMerchantCode   string `json:"customer_merchant_code,omitempty"`

}
