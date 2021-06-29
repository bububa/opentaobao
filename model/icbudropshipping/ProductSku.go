package icbudropshipping

// ProductSku 
type ProductSku struct {
    // sku image url
    ImageUrl   string `json:"image_url,omitempty" xml:"image_url,omitempty"`
    // inventory list
    InventoryList   []Inventory `json:"inventory_list,omitempty" xml:"inventory_list>inventory,omitempty"`
    // sku id
    SkuId   int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
    // sku name value list
    SkuNameValueList   []ProductSkuNameValue `json:"sku_name_value_list,omitempty" xml:"sku_name_value_list>product_sku_name_value,omitempty"`
    // ladder price list
    LadderPriceList   []LadderPrice `json:"ladder_price_list,omitempty" xml:"ladder_price_list>ladder_price,omitempty"`
}
