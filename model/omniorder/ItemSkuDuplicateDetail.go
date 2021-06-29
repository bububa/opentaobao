package omniorder

// ItemSkuDuplicateDetail 
type ItemSkuDuplicateDetail struct {

    // barcode
    
    Barcode   string `json:"barcode,omitempty" xml:"barcode,omitempty"`
    

    // itemId
    
    ItemId   int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
    

    // outerId
    
    OuterId   string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
    

    // skuBarcode
    
    SkuBarcode   string `json:"sku_barcode,omitempty" xml:"sku_barcode,omitempty"`
    

    // skuId
    
    SkuId   int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
    

    // skuOuterId
    
    SkuOuterId   string `json:"sku_outer_id,omitempty" xml:"sku_outer_id,omitempty"`
    

}
