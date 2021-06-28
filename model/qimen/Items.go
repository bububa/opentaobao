package qimen

// Items 
type Items struct {

    // 明细
    
    Item   *Item `json:"item,omitempty" xml:"item,omitempty"`
    

    // 货品编码,HZ1234,string(50),,
    
    ScItemCode   string `json:"scItemCode,omitempty" xml:"scItemCode,omitempty"`
    

    // 库存类型(1:可销售库存.101:残次),HZ1234,string(500),,
    
    InventoryType   string `json:"inventoryType,omitempty" xml:"inventoryType,omitempty"`
    

    // 实际出库数量,Item1234,string(50),,
    
    OutCount   string `json:"outCount,omitempty" xml:"outCount,omitempty"`
    

    // 实际出库数量,Item1234,string(50),,
    
    InCount   string `json:"inCount,omitempty" xml:"inCount,omitempty"`
    

    // 计划调拨数量
    
    PlanCount   string `json:"planCount,omitempty" xml:"planCount,omitempty"`
    

}
