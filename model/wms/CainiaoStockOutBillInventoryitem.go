package wms

// CainiaoStockOutBillInventoryitem 
type CainiaoStockOutBillInventoryitem struct {

    // 库存类型1 可销售库存 101 类型用来定义残次品 201 冻结类型库存301 在途库存
    
    InventoryType   int64 `json:"inventory_type,omitempty" xml:"inventory_type,omitempty"`
    

    // 数量
    
    ItemQty   int64 `json:"item_qty,omitempty" xml:"item_qty,omitempty"`
    

    // 商品保质期信息，失效日期
    
    DueDate   string `json:"due_date,omitempty" xml:"due_date,omitempty"`
    

    // 商品保质期信息，生产日期
    
    ProduceDate   string `json:"produce_date,omitempty" xml:"produce_date,omitempty"`
    

    // 生产编码，同一商品可能因商家不同有不同编码
    
    ProduceCode   string `json:"produce_code,omitempty" xml:"produce_code,omitempty"`
    

    // 生产地区
    
    ProduceArea   string `json:"produce_area,omitempty" xml:"produce_area,omitempty"`
    

    // 批次号
    
    BatchCode   string `json:"batch_code,omitempty" xml:"batch_code,omitempty"`
    

}
