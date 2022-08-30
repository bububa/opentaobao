package logistic

// QueryShippingMethodResponseDto 结构体
type QueryShippingMethodResponseDto struct {
	// delivery option list
	DeliveryOptions []AELogisticsShippingMethodDto `json:"delivery_options,omitempty" xml:"delivery_options>ae_logistics_shipping_method_dto,omitempty"`
	// query id, if the value is not empty, please fill it in create order api
	QueryId string `json:"query_id,omitempty" xml:"query_id,omitempty"`
	// parcel
	Parcel *ParcelDto `json:"parcel,omitempty" xml:"parcel,omitempty"`
}
