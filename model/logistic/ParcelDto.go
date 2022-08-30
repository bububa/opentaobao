package logistic

// ParcelDto 结构体
type ParcelDto struct {
	// Type: ENVELOPE, BOX, BAG, TUBE, PALLET
	ParcelTypeCode string `json:"parcel_type_code,omitempty" xml:"parcel_type_code,omitempty"`
	// Width of the parcel in the shipment order. Default unit: centimeters
	Width string `json:"width,omitempty" xml:"width,omitempty"`
	// Length of the parcel in the shipment order. Default unit: centimeters
	Length string `json:"length,omitempty" xml:"length,omitempty"`
	// Actual gross weight of the parcel instead of the weight of the product, nor the added weight of the invoice volumes.  Default unit: kilograms
	Weight string `json:"weight,omitempty" xml:"weight,omitempty"`
	// Height of the parcel in the shipment order. Default unit: centimeters
	Height string `json:"height,omitempty" xml:"height,omitempty"`
	// price of all SKUs in the parcel
	ProductTotalPrice string `json:"product_total_price,omitempty" xml:"product_total_price,omitempty"`
	// Parcel type including ENVELOPE, BOX, BAG, TUBE, PALLET and you should input it when creating a shipment order
	ParcelType string `json:"parcel_type,omitempty" xml:"parcel_type,omitempty"`
	// number of the parcel and there's only one parcel
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
}
