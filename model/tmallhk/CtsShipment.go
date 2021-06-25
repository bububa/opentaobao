package tmallhk

// CtsShipment 
type CtsShipment struct {

    // 报关开始时间
    Begin   string `json:"begin,omitempty"`

    // 报关结束时间
    End   string `json:"end,omitempty"`

    // 报关单号
    ShipmentNo   string `json:"shipment_no,omitempty"`

}
