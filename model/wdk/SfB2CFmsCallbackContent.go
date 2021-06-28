package wdk

// SfB2CFmsCallbackContent 
/* model for simplify = false
type SfB2CFmsCallbackContent struct {

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
    

    // 快递包裹信息
    
    Packages  struct {
        ExpressPackage  []ExpressPackage `json:"express_package,omitempty"`
    } `json:"packages,omitempty"`
    

}
*/

// SfB2CFmsCallbackContent 
type SfB2CFmsCallbackContent struct {

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

    // 快递包裹信息
    Packages   []ExpressPackage `json:"packages,omitempty"`

}
