package icbudropshipping

import (
	"sync"
)

// FreightTemplateRequest 结构体
type FreightTemplateRequest struct {
	// destination country  ISO 3166-2
	DestinationCountry string `json:"destination_country,omitempty" xml:"destination_country,omitempty"`
	// destination zip code
	ZipCode string `json:"zip_code,omitempty" xml:"zip_code,omitempty"`
	// product id
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// quantity
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
}

var poolFreightTemplateRequest = sync.Pool{
	New: func() any {
		return new(FreightTemplateRequest)
	},
}

// GetFreightTemplateRequest() 从对象池中获取FreightTemplateRequest
func GetFreightTemplateRequest() *FreightTemplateRequest {
	return poolFreightTemplateRequest.Get().(*FreightTemplateRequest)
}

// ReleaseFreightTemplateRequest 释放FreightTemplateRequest
func ReleaseFreightTemplateRequest(v *FreightTemplateRequest) {
	v.DestinationCountry = ""
	v.ZipCode = ""
	v.ProductId = 0
	v.Quantity = 0
	poolFreightTemplateRequest.Put(v)
}
