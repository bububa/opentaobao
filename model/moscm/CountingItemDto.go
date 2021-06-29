package moscm

// CountingItemDto 
type CountingItemDto struct {
    // 外部商品编码
    OutId   string `json:"out_id,omitempty" xml:"out_id,omitempty"`
    // 商品名称
    ProductName   string `json:"product_name,omitempty" xml:"product_name,omitempty"`
    // 库存数量
    Quantity   string `json:"quantity,omitempty" xml:"quantity,omitempty"`
    // 计量单位
    Unit   string `json:"unit,omitempty" xml:"unit,omitempty"`
}
