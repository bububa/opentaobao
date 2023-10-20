package axindata

import (
	"sync"
)

// FscProjectInventoryUpdateRequest 结构体
type FscProjectInventoryUpdateRequest struct {
	// 团期列表
	ProjectList []FscRouteProjectInventoryApiDto `json:"project_list,omitempty" xml:"project_list>fsc_route_project_inventory_api_dto,omitempty"`
	// 产品编号
	ProductCode string `json:"product_code,omitempty" xml:"product_code,omitempty"`
	// 行程编号
	JourneyCode string `json:"journey_code,omitempty" xml:"journey_code,omitempty"`
	// 供应商ID
	SupplierId string `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
}

var poolFscProjectInventoryUpdateRequest = sync.Pool{
	New: func() any {
		return new(FscProjectInventoryUpdateRequest)
	},
}

// GetFscProjectInventoryUpdateRequest() 从对象池中获取FscProjectInventoryUpdateRequest
func GetFscProjectInventoryUpdateRequest() *FscProjectInventoryUpdateRequest {
	return poolFscProjectInventoryUpdateRequest.Get().(*FscProjectInventoryUpdateRequest)
}

// ReleaseFscProjectInventoryUpdateRequest 释放FscProjectInventoryUpdateRequest
func ReleaseFscProjectInventoryUpdateRequest(v *FscProjectInventoryUpdateRequest) {
	v.ProjectList = v.ProjectList[:0]
	v.ProductCode = ""
	v.JourneyCode = ""
	v.SupplierId = ""
	poolFscProjectInventoryUpdateRequest.Put(v)
}
