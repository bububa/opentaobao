package axindata

import (
	"sync"
)

// FscProjectModifyRequest 结构体
type FscProjectModifyRequest struct {
	// 团期列表
	ProjectList []FscRouteProjectApiDto `json:"project_list,omitempty" xml:"project_list>fsc_route_project_api_dto,omitempty"`
	// 产品编号
	ProductCode string `json:"product_code,omitempty" xml:"product_code,omitempty"`
	// 行程编号
	JourneyCode string `json:"journey_code,omitempty" xml:"journey_code,omitempty"`
	// 供应商id
	SupplierId string `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
}

var poolFscProjectModifyRequest = sync.Pool{
	New: func() any {
		return new(FscProjectModifyRequest)
	},
}

// GetFscProjectModifyRequest() 从对象池中获取FscProjectModifyRequest
func GetFscProjectModifyRequest() *FscProjectModifyRequest {
	return poolFscProjectModifyRequest.Get().(*FscProjectModifyRequest)
}

// ReleaseFscProjectModifyRequest 释放FscProjectModifyRequest
func ReleaseFscProjectModifyRequest(v *FscProjectModifyRequest) {
	v.ProjectList = v.ProjectList[:0]
	v.ProductCode = ""
	v.JourneyCode = ""
	v.SupplierId = ""
	poolFscProjectModifyRequest.Put(v)
}
