package icbudropshipping

// PaymentDetail 
type PaymentDetail struct {
    // shipment fee
    ShipmentFee   string `json:"shipment_fee,omitempty" xml:"shipment_fee,omitempty"`
    // total amount
    TotalAmount   string `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
}
