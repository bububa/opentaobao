package logistic

import (
	"sync"
)

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
	// number of the parcel and there&#39;s only one parcel
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
}

var poolParcelDto = sync.Pool{
	New: func() any {
		return new(ParcelDto)
	},
}

// GetParcelDto() 从对象池中获取ParcelDto
func GetParcelDto() *ParcelDto {
	return poolParcelDto.Get().(*ParcelDto)
}

// ReleaseParcelDto 释放ParcelDto
func ReleaseParcelDto(v *ParcelDto) {
	v.ParcelTypeCode = ""
	v.Width = ""
	v.Length = ""
	v.Weight = ""
	v.Height = ""
	v.ProductTotalPrice = ""
	v.ParcelType = ""
	v.Quantity = 0
	poolParcelDto.Put(v)
}
