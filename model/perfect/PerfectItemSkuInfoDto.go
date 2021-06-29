package perfect

// PerfectItemSkuInfoDTO 
type PerfectItemSkuInfoDTO struct {
    // sku销售属性
    SaleProperties   []PerfectSalePropertyDTO `json:"sale_properties,omitempty" xml:"sale_properties>perfect_sale_property_dto,omitempty"`
    // sku关联货品
    ScProductInfo   *PerfectScProductInfoDTO `json:"sc_product_info,omitempty" xml:"sc_product_info,omitempty"`
    // sku条码
    SkuBarcode   string `json:"sku_barcode,omitempty" xml:"sku_barcode,omitempty"`
    // sku外部编码
    SkuOuterId   string `json:"sku_outer_id,omitempty" xml:"sku_outer_id,omitempty"`
    // sku吊牌价
    SkuPretium   string `json:"sku_pretium,omitempty" xml:"sku_pretium,omitempty"`
    // sku价格
    SkuPrice   string `json:"sku_price,omitempty" xml:"sku_price,omitempty"`
    // sku数量
    SkuQuantity   int64 `json:"sku_quantity,omitempty" xml:"sku_quantity,omitempty"`
}
