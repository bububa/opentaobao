package wdk

// DrfB2CCallbackContent 
/* model for simplify = false
type DrfB2CCallbackContent struct {

    // 是否缺货出
    
    IsShortage   bool `json:"is_shortage,omitempty"`
    

    // 商品编码
    
    SkuCode   string `json:"sku_code,omitempty"`
    

    // 缺货出库存数量
    
    OutOfStockStockQuantity   string `json:"out_of_stock_stock_quantity,omitempty"`
    

    // 缺货出销售数量
    
    OutOfStockSaleQuantity   string `json:"out_of_stock_sale_quantity,omitempty"`
    

    // 实际库存拣货数量
    
    ActualStockQuantity   string `json:"actual_stock_quantity,omitempty"`
    

    // 实际销售拣货数量
    
    ActualSaleQuantity   string `json:"actual_sale_quantity,omitempty"`
    

    // 作业内容单号
    
    WorkUnitContentId   string `json:"work_unit_content_id,omitempty"`
    

}
*/

// DrfB2CCallbackContent 
type DrfB2CCallbackContent struct {

    // 是否缺货出
    IsShortage   bool `json:"is_shortage,omitempty"`

    // 商品编码
    SkuCode   string `json:"sku_code,omitempty"`

    // 缺货出库存数量
    OutOfStockStockQuantity   string `json:"out_of_stock_stock_quantity,omitempty"`

    // 缺货出销售数量
    OutOfStockSaleQuantity   string `json:"out_of_stock_sale_quantity,omitempty"`

    // 实际库存拣货数量
    ActualStockQuantity   string `json:"actual_stock_quantity,omitempty"`

    // 实际销售拣货数量
    ActualSaleQuantity   string `json:"actual_sale_quantity,omitempty"`

    // 作业内容单号
    WorkUnitContentId   string `json:"work_unit_content_id,omitempty"`

}
