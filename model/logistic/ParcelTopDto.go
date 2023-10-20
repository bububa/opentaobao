package logistic

import (
	"sync"
)

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
	// number of the parcel and there&#39;s only one parcel
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
}

var poolParcelTopDto = sync.Pool{
	New: func() any {
		return new(ParcelTopDto)
	},
}

// GetParcelTopDto() 从对象池中获取ParcelTopDto
func GetParcelTopDto() *ParcelTopDto {
	return poolParcelTopDto.Get().(*ParcelTopDto)
}

// ReleaseParcelTopDto 释放ParcelTopDto
func ReleaseParcelTopDto(v *ParcelTopDto) {
	v.ParcelTypeCode = ""
	v.Length = ""
	v.Width = ""
	v.Weight = ""
	v.Height = ""
	v.Quantity = 0
	poolParcelTopDto.Put(v)
}
