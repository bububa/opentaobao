package nlife

// ItemCodeConvertResponse 
type ItemCodeConvertResponse struct {
    // 转码后的结果
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    // itemId
    ItemId   int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
    // skuId
    SkuId   int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
    // storeId
    StoreId   int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
    // BARCODE / UNIQUECODE
    CodeType   string `json:"code_type,omitempty" xml:"code_type,omitempty"`
}
