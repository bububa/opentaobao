package qimen

// TransferItems 
type TransferItems struct {

    // 库存类型(1:可销售库存.101:残次),HZ1234,string(500),必填,
    
    ScItemCode   string `json:"scItemCode,omitempty" xml:"scItemCode,omitempty"`
    

    // 数量,Item1234,string(50),必填,
    
    Count   string `json:"count,omitempty" xml:"count,omitempty"`
    

    // 货品编码,HZ1234,string(50),必填,
    
    InventoryType   string `json:"inventoryType,omitempty" xml:"inventoryType,omitempty"`
    

}
