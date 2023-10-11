package tblogistics

// PriorCallDeliveryTopResponse 结构体
type PriorCallDeliveryTopResponse struct {
	// 资源列表
	ResourceList []ResourceDto `json:"resource_list,omitempty" xml:"resource_list>resource_dto,omitempty"`
}
