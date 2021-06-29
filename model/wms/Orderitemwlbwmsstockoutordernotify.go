package wms

// Orderitemwlbwmsstockoutordernotify 
type Orderitemwlbwmsstockoutordernotify struct {
    // 生产日期
    ProduceDate   string `json:"produce_date,omitempty" xml:"produce_date,omitempty"`
    // 商品数量
    ItemQuantity   int64 `json:"item_quantity,omitempty" xml:"item_quantity,omitempty"`
    // 批次号
    BatchCode   string `json:"batch_code,omitempty" xml:"batch_code,omitempty"`
    // ERP主键ID
    OrderItemId   string `json:"order_item_id,omitempty" xml:"order_item_id,omitempty"`
    // ERP商品ID
    ItemId   string `json:"item_id,omitempty" xml:"item_id,omitempty"`
    // 生产编码，同一商品可能因商家不同有不同编码
    ProduceCode   string `json:"produce_code,omitempty" xml:"produce_code,omitempty"`
    // 订单商品拓展属性数据
    ExtendFields   string `json:"extend_fields,omitempty" xml:"extend_fields,omitempty"`
    // 库存类型
    InventoryType   int64 `json:"inventory_type,omitempty" xml:"inventory_type,omitempty"`
    // 到货日期
    DueDate   string `json:"due_date,omitempty" xml:"due_date,omitempty"`
}
