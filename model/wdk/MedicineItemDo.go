package wdk

// MedicineItemDo 
type MedicineItemDo struct {

    // 数量
    Count   int64 `json:"count,omitempty"`

    // sku名称
    SkuTitle   string `json:"sku_title,omitempty"`

    // sku编码
    SkuCode   string `json:"sku_code,omitempty"`

}
