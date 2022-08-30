package logistic

// AELogisticsShippingMethodDto 结构体
type AELogisticsShippingMethodDto struct {
	// Logistics provider Id of the shipping order such as 1-Cainiao, 2-Pegaki, 3-Frenet, 4-Delivery Hub, etc.
	LogisticsChannelId string `json:"logistics_channel_id,omitempty" xml:"logistics_channel_id,omitempty"`
	// Logistics provider of the shipping order such as Cainiao, Pegaki, Frenet, Delivery Hub, etc.
	LogisticsChannelName string `json:"logistics_channel_name,omitempty" xml:"logistics_channel_name,omitempty"`
	// shipping method id of Logistics service such as Loggi Express, Correios Sedex, etc.
	DeliveryMethodId string `json:"delivery_method_id,omitempty" xml:"delivery_method_id,omitempty"`
	// Devlivery provider type such as express/standard, pac/sedex.
	DeliveryMethodName string `json:"delivery_method_name,omitempty" xml:"delivery_method_name,omitempty"`
	// Delivery provider name
	DeliveryProviderName string `json:"delivery_provider_name,omitempty" xml:"delivery_provider_name,omitempty"`
	// estimated shipping cost based on dimension and weight of a parcel
	DeliveryCost int64 `json:"delivery_cost,omitempty" xml:"delivery_cost,omitempty"`
	// delivery time on working days
	DeliveryEstimateDays int64 `json:"delivery_estimate_days,omitempty" xml:"delivery_estimate_days,omitempty"`
}
