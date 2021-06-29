package wdk

// ChannelSkuQueryDO 
type ChannelSkuQueryDO struct {
    // skucode集合
    SkuCodes   []string `json:"sku_codes,omitempty" xml:"sku_codes>string,omitempty"`
    // 门店或DC编码
    OuCode   string `json:"ou_code,omitempty" xml:"ou_code,omitempty"`
    // 渠道编码（默认-1）
    ChannelCode   int64 `json:"channel_code,omitempty" xml:"channel_code,omitempty"`
    // 出货仓（默认为店仓一体，仓编码为店编码）
    DeliverWarehouse   string `json:"deliver_warehouse,omitempty" xml:"deliver_warehouse,omitempty"`
    // 客户商家编码
    CustomerMerchantCode   string `json:"customer_merchant_code,omitempty" xml:"customer_merchant_code,omitempty"`
}
