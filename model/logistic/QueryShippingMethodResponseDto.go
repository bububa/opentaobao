package logistic

import (
	"sync"
)

// QueryShippingMethodResponseDto 结构体
type QueryShippingMethodResponseDto struct {
	// delivery option list
	DeliveryOptions []AELogisticsShippingMethodDto `json:"delivery_options,omitempty" xml:"delivery_options>ae_logistics_shipping_method_dto,omitempty"`
	// query id, if the value is not empty, please fill it in create order api
	QueryId string `json:"query_id,omitempty" xml:"query_id,omitempty"`
	// parcel
	Parcel *ParcelDto `json:"parcel,omitempty" xml:"parcel,omitempty"`
}

var poolQueryShippingMethodResponseDto = sync.Pool{
	New: func() any {
		return new(QueryShippingMethodResponseDto)
	},
}

// GetQueryShippingMethodResponseDto() 从对象池中获取QueryShippingMethodResponseDto
func GetQueryShippingMethodResponseDto() *QueryShippingMethodResponseDto {
	return poolQueryShippingMethodResponseDto.Get().(*QueryShippingMethodResponseDto)
}

// ReleaseQueryShippingMethodResponseDto 释放QueryShippingMethodResponseDto
func ReleaseQueryShippingMethodResponseDto(v *QueryShippingMethodResponseDto) {
	v.DeliveryOptions = v.DeliveryOptions[:0]
	v.QueryId = ""
	v.Parcel = nil
	poolQueryShippingMethodResponseDto.Put(v)
}
