package tmallchannel

// ProductSkuTopDTO 
type ProductSkuTopDTO struct {
    // skuId
    SkuId   int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
    // sku商家编码
    SkuNumber   string `json:"sku_number,omitempty" xml:"sku_number,omitempty"`
    // sku后端货品
    SkuScItemId   int64 `json:"sku_sc_item_id,omitempty" xml:"sku_sc_item_id,omitempty"`
    // 基准价
    StandardPrice   int64 `json:"standard_price,omitempty" xml:"standard_price,omitempty"`
    // 条形码
    BarCode   string `json:"bar_code,omitempty" xml:"bar_code,omitempty"`
    // 图片链接
    PictureUrl   string `json:"picture_url,omitempty" xml:"picture_url,omitempty"`
}
