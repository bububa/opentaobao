package tmallhk

// AwdcInfo 
/* model for simplify = false
type AwdcInfo struct {

    // hrd info
    
    Hrd  *struct {
        AwdcHrd  *AwdcHrd `json:"awdc_hrd,omitempty"`
    } `json:"hrd,omitempty"`
    

    // 货品ID
    
    ItemId   string `json:"item_id,omitempty"`
    

    // ngtc info
    
    Ngtc  *struct {
        AwdcNgtc  *AwdcNgtc `json:"awdc_ngtc,omitempty"`
    } `json:"ngtc,omitempty"`
    

    // 订单号
    
    OrderNo   string `json:"order_no,omitempty"`
    

    // 商品ID
    
    ProductId   string `json:"product_id,omitempty"`
    

    // shipment
    
    Shipment  *struct {
        AwdcShipment  *AwdcShipment `json:"awdc_shipment,omitempty"`
    } `json:"shipment,omitempty"`
    

    // 子订单号
    
    SubOrderNo   string `json:"sub_order_no,omitempty"`
    

    // 溯源码
    
    TraceCode   string `json:"trace_code,omitempty"`
    

    // 溯源码镭射时间
    
    TraceCodeTime   string `json:"trace_code_time,omitempty"`
    

}
*/

// AwdcInfo 
type AwdcInfo struct {

    // hrd info
    Hrd   *AwdcHrd `json:"hrd,omitempty"`

    // 货品ID
    ItemId   string `json:"item_id,omitempty"`

    // ngtc info
    Ngtc   *AwdcNgtc `json:"ngtc,omitempty"`

    // 订单号
    OrderNo   string `json:"order_no,omitempty"`

    // 商品ID
    ProductId   string `json:"product_id,omitempty"`

    // shipment
    Shipment   *AwdcShipment `json:"shipment,omitempty"`

    // 子订单号
    SubOrderNo   string `json:"sub_order_no,omitempty"`

    // 溯源码
    TraceCode   string `json:"trace_code,omitempty"`

    // 溯源码镭射时间
    TraceCodeTime   string `json:"trace_code_time,omitempty"`

}
