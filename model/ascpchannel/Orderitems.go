package ascpchannel

// Orderitems 
type Orderitems struct {

    // 履约子单号
    
    SubOrderCode   string `json:"sub_order_code,omitempty" xml:"sub_order_code,omitempty"`
    

    // 货品ID
    
    ScItemId   string `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
    

    // 货品实发数量
    
    ItemQuantity   int64 `json:"item_quantity,omitempty" xml:"item_quantity,omitempty"`
    

    // 货品缺发数量
    
    LackQuantity   int64 `json:"lack_quantity,omitempty" xml:"lack_quantity,omitempty"`
    

    // erp订单明细行号
    
    ErpOrderLine   string `json:"erp_order_line,omitempty" xml:"erp_order_line,omitempty"`
    

    // 货品计划退回数量
    
    PlanReturnQuantity   int64 `json:"plan_return_quantity,omitempty" xml:"plan_return_quantity,omitempty"`
    

    // 货品实际收货总数量
    
    ActualReceivedQuantity   int64 `json:"actual_received_quantity,omitempty" xml:"actual_received_quantity,omitempty"`
    

    // 货品未收货总数量
    
    ActualLackQuantity   int64 `json:"actual_lack_quantity,omitempty" xml:"actual_lack_quantity,omitempty"`
    

    // 销退回告明细列表
    
    InstorageDetails   []Instoragedetails `json:"instorage_details,omitempty" xml:"instorage_details,omitempty"`
    

}
