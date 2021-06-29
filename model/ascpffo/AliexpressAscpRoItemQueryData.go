package ascpffo

// AliexpressAscpRoItemQueryData 
type AliexpressAscpRoItemQueryData struct {

    // 退供单号
    
    ReturnOrderNo   string `json:"return_order_no,omitempty" xml:"return_order_no,omitempty"`
    

    // 货品Id
    
    ScItemId   int64 `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
    

    // 货品名称
    
    Title   string `json:"title,omitempty" xml:"title,omitempty"`
    

    // 退供数量
    
    Quantity   int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
    

    // 实际退供数量
    
    ReturnQuantity   int64 `json:"return_quantity,omitempty" xml:"return_quantity,omitempty"`
    

    // 库存类型
    
    InventoryTypeDesc   string `json:"inventory_type_desc,omitempty" xml:"inventory_type_desc,omitempty"`
    

    // 税率
    
    TaxRate   string `json:"tax_rate,omitempty" xml:"tax_rate,omitempty"`
    

    // 含税退供价，单位分
    
    ReturnPrice   string `json:"return_price,omitempty" xml:"return_price,omitempty"`
    

    // 扩展字段
    
    ExtendFields   string `json:"extend_fields,omitempty" xml:"extend_fields,omitempty"`
    

}
