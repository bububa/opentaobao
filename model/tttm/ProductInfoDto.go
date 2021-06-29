package tttm

// ProductInfoDTO 
type ProductInfoDTO struct {
    // 货品编码
    ProductCode   string `json:"product_code,omitempty" xml:"product_code,omitempty"`
    // 货品名称
    ProductName   string `json:"product_name,omitempty" xml:"product_name,omitempty"`
    // 货品状态
    ProductStatus   int64 `json:"product_status,omitempty" xml:"product_status,omitempty"`
    // 套餐数量
    SetAmount   int64 `json:"set_amount,omitempty" xml:"set_amount,omitempty"`
    // 总库存
    TotalAmount   int64 `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
    // 增量库存
    IncrementAmount   int64 `json:"increment_amount,omitempty" xml:"increment_amount,omitempty"`
    // 出入库
    IncrementType   int64 `json:"increment_type,omitempty" xml:"increment_type,omitempty"`
}
