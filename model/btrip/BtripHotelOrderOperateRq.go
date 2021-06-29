package btrip

// BtripHotelOrderOperateRq 
type BtripHotelOrderOperateRq struct {
    // 分销商订单id
    DisOrderId   string `json:"dis_order_id,omitempty" xml:"dis_order_id,omitempty"`
    // 分销商子渠道，通常是商旅企业id
    SubChannel   string `json:"sub_channel,omitempty" xml:"sub_channel,omitempty"`
    // 供应商标识
    SupplierCode   string `json:"supplier_code,omitempty" xml:"supplier_code,omitempty"`
    // 商旅订单id
    BtripOrderId   int64 `json:"btrip_order_id,omitempty" xml:"btrip_order_id,omitempty"`
}
