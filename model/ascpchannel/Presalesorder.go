package ascpchannel

// Presalesorder 
type Presalesorder struct {

    // 包裹信息
    
    Packages   []Packages `json:"packages,omitempty" xml:"packages,omitempty"`
    

    // 订单完成时间
    
    OrderConfirmTime   string `json:"order_confirm_time,omitempty" xml:"order_confirm_time,omitempty"`
    

    // 菜鸟订单号
    
    PresalesOrderId   string `json:"presales_order_id,omitempty" xml:"presales_order_id,omitempty"`
    

    // 出库单号
    
    PresalesOrderCode   string `json:"presales_order_code,omitempty" xml:"presales_order_code,omitempty"`
    

    // 发货仓code
    
    StoreCode   string `json:"store_code,omitempty" xml:"store_code,omitempty"`
    

    // 扩展属性(JSON格式)
    
    ExtendFields   string `json:"extend_fields,omitempty" xml:"extend_fields,omitempty"`
    

}
