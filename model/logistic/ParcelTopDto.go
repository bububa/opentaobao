package logistic

// ParcelTopDto 结构体
type ParcelTopDto struct {
	// Type: ENVELOPE, BOX, BAG, TUBE, PALLET
	ParcelTypeCode string `json:"parcel_type_code,omitempty" xml:"parcel_type_code,omitempty"`
	// Length of the parcel in the shipment order. Default unit: centimeters
	Length string `json:"length,omitempty" xml:"length,omitempty"`
	// Width of the parcel in the shipment order. Default unit: centimeters
	Width string `json:"width,omitempty" xml:"width,omitempty"`
	// Actual gross weight of the parcel instead of the weight of the product, nor the added weight of the invoice volumes.  Default unit: kilograms
	Weight string `json:"weight,omitempty" xml:"weight,omitempty"`
	// Height of the parcel in the shipment order. Default unit: centimeters
	Height string `json:"height,omitempty" xml:"height,omitempty"`
	// number of the parcel and there's only one parcel
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
}
