package icbudropshipping

// MoqAndPrice 
type MoqAndPrice struct {
    // min order quantity  delivery period
    MoqDeliveryPeriod   int64 `json:"moq_delivery_period,omitempty" xml:"moq_delivery_period,omitempty"`
    // min order quantity unit
    Unit   string `json:"unit,omitempty" xml:"unit,omitempty"`
    // min order quantity unit price
    MoqUnitPrice   float64 `json:"moq_unit_price,omitempty" xml:"moq_unit_price,omitempty"`
    // min order quantity
    MinOrderQuantity   string `json:"min_order_quantity,omitempty" xml:"min_order_quantity,omitempty"`
}
