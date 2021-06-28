package tmallhk

// CtsInfo 
/* model for simplify = false
type CtsInfo struct {

    // 托运
    
    Carriage  *struct {
        CtsCarriage  *CtsCarriage `json:"cts_carriage,omitempty"`
    } `json:"carriage,omitempty"`
    

    // 成品证书
    
    CompletedNgtc  *struct {
        CtsNgtc  *CtsNgtc `json:"cts_ngtc,omitempty"`
    } `json:"completed_ngtc,omitempty"`
    

    // 国内物流
    
    Delivery  *struct {
        CtsDelivery  *CtsDelivery `json:"cts_delivery,omitempty"`
    } `json:"delivery,omitempty"`
    

    // 裸钻证书
    
    DiamondNgtc  *struct {
        CtsNgtc  *CtsNgtc `json:"cts_ngtc,omitempty"`
    } `json:"diamond_ngtc,omitempty"`
    

    // 货品ID
    
    ItemId   string `json:"item_id,omitempty"`
    

    // 订单号
    
    OrderNo   string `json:"order_no,omitempty"`
    

    // 商品ID
    
    ProductId   string `json:"product_id,omitempty"`
    

    // 戒托信息
    
    Ring  *struct {
        CtsRing  *CtsRing `json:"cts_ring,omitempty"`
    } `json:"ring,omitempty"`
    

    // 国内报关
    
    Shipment  *struct {
        CtsShipment  *CtsShipment `json:"cts_shipment,omitempty"`
    } `json:"shipment,omitempty"`
    

    // 子订单号
    
    SubOrderNo   string `json:"sub_order_no,omitempty"`
    

    // 溯源码
    
    TraceCode   string `json:"trace_code,omitempty"`
    

}
*/

// CtsInfo 
type CtsInfo struct {

    // 托运
    Carriage   *CtsCarriage `json:"carriage,omitempty"`

    // 成品证书
    CompletedNgtc   *CtsNgtc `json:"completed_ngtc,omitempty"`

    // 国内物流
    Delivery   *CtsDelivery `json:"delivery,omitempty"`

    // 裸钻证书
    DiamondNgtc   *CtsNgtc `json:"diamond_ngtc,omitempty"`

    // 货品ID
    ItemId   string `json:"item_id,omitempty"`

    // 订单号
    OrderNo   string `json:"order_no,omitempty"`

    // 商品ID
    ProductId   string `json:"product_id,omitempty"`

    // 戒托信息
    Ring   *CtsRing `json:"ring,omitempty"`

    // 国内报关
    Shipment   *CtsShipment `json:"shipment,omitempty"`

    // 子订单号
    SubOrderNo   string `json:"sub_order_no,omitempty"`

    // 溯源码
    TraceCode   string `json:"trace_code,omitempty"`

}
