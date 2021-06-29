package wdk

// MedicineItemDO 
type MedicineItemDO struct {
    // 数量
    Count   int64 `json:"count,omitempty" xml:"count,omitempty"`
    // sku名称
    SkuTitle   string `json:"sku_title,omitempty" xml:"sku_title,omitempty"`
    // sku编码
    SkuCode   string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
}
