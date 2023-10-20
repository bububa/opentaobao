package icbudropshipping

import (
	"sync"
)

// MultiFreightTemplateRequest 结构体
type MultiFreightTemplateRequest struct {
	// Product List
	LogisticsProductList []LogisticsProduct `json:"logistics_product_list,omitempty" xml:"logistics_product_list>logistics_product,omitempty"`
	// Get from alibaba.dropshipping.product.get
	ECompanyId string `json:"e_company_id,omitempty" xml:"e_company_id,omitempty"`
	// Destination Country
	DestinationCountry string `json:"destination_country,omitempty" xml:"destination_country,omitempty"`
	// Shipping address
	Address *AddressInfoDto `json:"address,omitempty" xml:"address,omitempty"`
}

var poolMultiFreightTemplateRequest = sync.Pool{
	New: func() any {
		return new(MultiFreightTemplateRequest)
	},
}

// GetMultiFreightTemplateRequest() 从对象池中获取MultiFreightTemplateRequest
func GetMultiFreightTemplateRequest() *MultiFreightTemplateRequest {
	return poolMultiFreightTemplateRequest.Get().(*MultiFreightTemplateRequest)
}

// ReleaseMultiFreightTemplateRequest 释放MultiFreightTemplateRequest
func ReleaseMultiFreightTemplateRequest(v *MultiFreightTemplateRequest) {
	v.LogisticsProductList = v.LogisticsProductList[:0]
	v.ECompanyId = ""
	v.DestinationCountry = ""
	v.Address = nil
	poolMultiFreightTemplateRequest.Put(v)
}
