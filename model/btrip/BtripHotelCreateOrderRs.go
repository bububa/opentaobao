package btrip

// BtripHotelCreateOrderRs 
type BtripHotelCreateOrderRs struct {
    // 商旅订单id
    BtripOrderId   int64 `json:"btrip_order_id,omitempty" xml:"btrip_order_id,omitempty"`
    // 供应商订单id
    SupplierOrderId   string `json:"supplier_order_id,omitempty" xml:"supplier_order_id,omitempty"`
}
