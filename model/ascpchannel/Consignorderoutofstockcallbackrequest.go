package ascpchannel

// Consignorderoutofstockcallbackrequest 
type Consignorderoutofstockcallbackrequest struct {

    // 供应商id
    
    SupplierId   string `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
    

    // 外部业务号
    
    OutBizId   string `json:"out_biz_id,omitempty" xml:"out_biz_id,omitempty"`
    

    // 履约单号
    
    BizOrderCode   string `json:"biz_order_code,omitempty" xml:"biz_order_code,omitempty"`
    

    // 缺货明细
    
    OutOfStockItems   []Outofstockitems `json:"out_of_stock_items,omitempty" xml:"out_of_stock_items,omitempty"`
    

    // 缺货原因
    
    OutOfStockReason   string `json:"out_of_stock_reason,omitempty" xml:"out_of_stock_reason,omitempty"`
    

}
