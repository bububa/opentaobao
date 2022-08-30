package tmallservice

// ServiceCapacityAdjustReqDto 结构体
type ServiceCapacityAdjustReqDto struct {
	// 具体的调整值参数
	DayQuantityList []DayQuantity `json:"day_quantity_list,omitempty" xml:"day_quantity_list>day_quantity,omitempty"`
	// 服务提供者
	ServiceProviderDto *ServiceProviderDto `json:"service_provider_dto,omitempty" xml:"service_provider_dto,omitempty"`
}
