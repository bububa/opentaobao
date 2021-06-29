package omniorder

// ItemLightPublishSkuDTO 
type ItemLightPublishSkuDTO struct {
    // sku条形码
    Barcode   string `json:"barcode,omitempty" xml:"barcode,omitempty"`
    // sku扩展字段
    ExtendAttr   string `json:"extend_attr,omitempty" xml:"extend_attr,omitempty"`
    // sku销售价
    Pretium   string `json:"pretium,omitempty" xml:"pretium,omitempty"`
    // sku吊牌价
    Price   string `json:"price,omitempty" xml:"price,omitempty"`
    // salePropsDTO
    SalePropsDTOs   []ItemLightPublishSalePropDTO `json:"sale_props_d_t_os,omitempty" xml:"sale_props_d_t_os>item_light_publish_sale_prop_dto,omitempty"`
    // sku图片
    SkuImages   []string `json:"sku_images,omitempty" xml:"sku_images>string,omitempty"`
    // skuOuterId
    SkuOuterId   string `json:"sku_outer_id,omitempty" xml:"sku_outer_id,omitempty"`
    // sku条形码
    SkuBarcode   string `json:"sku_barcode,omitempty" xml:"sku_barcode,omitempty"`
    // sku销售属性
    SaleProps   []ItemLightPublishSalePropDTO `json:"sale_props,omitempty" xml:"sale_props>item_light_publish_sale_prop_dto,omitempty"`
    // skuId
    SkuId   int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
    // customCode
    CustomCode   string `json:"custom_code,omitempty" xml:"custom_code,omitempty"`
}
