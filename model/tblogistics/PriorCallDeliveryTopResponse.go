package tblogistics

import (
	"sync"
)

// PriorCallDeliveryTopResponse 结构体
type PriorCallDeliveryTopResponse struct {
	// 资源列表
	ResourceList []ResourceDto `json:"resource_list,omitempty" xml:"resource_list>resource_dto,omitempty"`
}

var poolPriorCallDeliveryTopResponse = sync.Pool{
	New: func() any {
		return new(PriorCallDeliveryTopResponse)
	},
}

// GetPriorCallDeliveryTopResponse() 从对象池中获取PriorCallDeliveryTopResponse
func GetPriorCallDeliveryTopResponse() *PriorCallDeliveryTopResponse {
	return poolPriorCallDeliveryTopResponse.Get().(*PriorCallDeliveryTopResponse)
}

// ReleasePriorCallDeliveryTopResponse 释放PriorCallDeliveryTopResponse
func ReleasePriorCallDeliveryTopResponse(v *PriorCallDeliveryTopResponse) {
	v.ResourceList = v.ResourceList[:0]
	poolPriorCallDeliveryTopResponse.Put(v)
}
