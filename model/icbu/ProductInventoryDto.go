package icbu

// ProductInventoryDto 
type ProductInventoryDto struct {

    // 库存编码，为空时表示默认国内仓
    
    StoreCode   string `json:"store_code,omitempty" xml:"store_code,omitempty"`
    

    // 库存值
    
    Inventory   int64 `json:"inventory,omitempty" xml:"inventory,omitempty"`
    

}
