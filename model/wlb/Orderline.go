package wlb

// OrderLine 
type OrderLine struct {
    // 库存类型，ZP=正品、CC=残次
    InventoryType   string `json:"inventory_type,omitempty" xml:"inventory_type,omitempty"`
    // 原交易单，供销平台交易单号、分销平台单号
    SourceOrderCode   string `json:"source_order_code,omitempty" xml:"source_order_code,omitempty"`
    // 子交易单号
    SubSourceOrderCode   string `json:"sub_source_order_code,omitempty" xml:"sub_source_order_code,omitempty"`
    // 批次编码
    BatchCode   string `json:"batch_code,omitempty" xml:"batch_code,omitempty"`
    // 生产日期，生产日期(YYYY-MM-DD)
    ProductDate   string `json:"product_date,omitempty" xml:"product_date,omitempty"`
    // 过期日期，生产日期(YYYY-MM-DD)
    ExpireDate   string `json:"expire_date,omitempty" xml:"expire_date,omitempty"`
    // 生产批号
    ProduceCode   string `json:"produce_code,omitempty" xml:"produce_code,omitempty"`
    // 商品数量
    ItemQuantity   int64 `json:"item_quantity,omitempty" xml:"item_quantity,omitempty"`
    // 商品编码
    ItemCode   string `json:"item_code,omitempty" xml:"item_code,omitempty"`
    // 商品名称
    ItemName   string `json:"item_name,omitempty" xml:"item_name,omitempty"`
    // 商品ID
    ItemId   string `json:"item_id,omitempty" xml:"item_id,omitempty"`
    // 订单行号
    OrderLineNo   string `json:"order_line_no,omitempty" xml:"order_line_no,omitempty"`
}
