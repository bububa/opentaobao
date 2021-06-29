package icbudropshipping

// LogisticsDetail 
type LogisticsDetail struct {
    // For BuyNow orders, use the value shown in the vendorCode field from the shipping cost template API; non-BuyNow orders don’t need to provide this information.    alibaba.shipping.freight.calculate  's vender_code
    CarrierCode   string `json:"carrier_code,omitempty" xml:"carrier_code,omitempty"`
    // shipment address
    ShipmentAddress   *Address `json:"shipment_address,omitempty" xml:"shipment_address,omitempty"`
}
