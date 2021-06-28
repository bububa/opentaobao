package wms

// Orderitemwlbwmsstockinordernotifywl 
type Orderitemwlbwmsstockinordernotifywl struct {

    // 生产批号
    
    ProduceCode   string `json:"produce_code,omitempty" xml:"produce_code,omitempty"`
    

    // 商品数量
    
    ItemQuantity   int64 `json:"item_quantity,omitempty" xml:"item_quantity,omitempty"`
    

    // 批次编号
    
    BatchCode   string `json:"batch_code,omitempty" xml:"batch_code,omitempty"`
    

    // 单据明细ID
    
    OrderItemId   string `json:"order_item_id,omitempty" xml:"order_item_id,omitempty"`
    

    // 订单商品拓展属性
    
    ExtendFields   string `json:"extend_fields,omitempty" xml:"extend_fields,omitempty"`
    

    // 库存类型 1 正品 101 残次 102 机损 103 箱损 201 冻结库存 301 在途库存
    
    InventoryType   int64 `json:"inventory_type,omitempty" xml:"inventory_type,omitempty"`
    

    // 后端商品ID
    
    ItemId   string `json:"item_id,omitempty" xml:"item_id,omitempty"`
    

    // 失效日期
    
    DueDate   string `json:"due_date,omitempty" xml:"due_date,omitempty"`
    

    // 采购价格
    
    PurchasePrice   int64 `json:"purchase_price,omitempty" xml:"purchase_price,omitempty"`
    

    // 生产日期
    
    ProduceDate   string `json:"produce_date,omitempty" xml:"produce_date,omitempty"`
    

}
