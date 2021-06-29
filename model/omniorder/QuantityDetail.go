package omniorder

// QuantityDetail 
type QuantityDetail struct {

    // 库存类型
    
    InventoryType   string `json:"inventory_type,omitempty" xml:"inventory_type,omitempty"`
    

    // 库存数量
    
    Quantity   int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
    

}
