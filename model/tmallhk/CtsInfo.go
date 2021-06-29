package tmallhk

// CtsInfo 
type CtsInfo struct {
    // 托运
    Carriage   *CtsCarriage `json:"carriage,omitempty" xml:"carriage,omitempty"`
    // 成品证书
    CompletedNgtc   *CtsNgtc `json:"completed_ngtc,omitempty" xml:"completed_ngtc,omitempty"`
    // 国内物流
    Delivery   *CtsDelivery `json:"delivery,omitempty" xml:"delivery,omitempty"`
    // 裸钻证书
    DiamondNgtc   *CtsNgtc `json:"diamond_ngtc,omitempty" xml:"diamond_ngtc,omitempty"`
    // 货品ID
    ItemId   string `json:"item_id,omitempty" xml:"item_id,omitempty"`
    // 订单号
    OrderNo   string `json:"order_no,omitempty" xml:"order_no,omitempty"`
    // 商品ID
    ProductId   string `json:"product_id,omitempty" xml:"product_id,omitempty"`
    // 戒托信息
    Ring   *CtsRing `json:"ring,omitempty" xml:"ring,omitempty"`
    // 国内报关
    Shipment   *CtsShipment `json:"shipment,omitempty" xml:"shipment,omitempty"`
    // 子订单号
    SubOrderNo   string `json:"sub_order_no,omitempty" xml:"sub_order_no,omitempty"`
    // 溯源码
    TraceCode   string `json:"trace_code,omitempty" xml:"trace_code,omitempty"`
}
