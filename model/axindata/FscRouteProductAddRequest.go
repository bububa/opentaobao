package axindata

import (
	"sync"
)

// FscRouteProductAddRequest 结构体
type FscRouteProductAddRequest struct {
	// 线路产品信息
	RouteInfo *FscRouteInfoApiDto `json:"route_info,omitempty" xml:"route_info,omitempty"`
}

var poolFscRouteProductAddRequest = sync.Pool{
	New: func() any {
		return new(FscRouteProductAddRequest)
	},
}

// GetFscRouteProductAddRequest() 从对象池中获取FscRouteProductAddRequest
func GetFscRouteProductAddRequest() *FscRouteProductAddRequest {
	return poolFscRouteProductAddRequest.Get().(*FscRouteProductAddRequest)
}

// ReleaseFscRouteProductAddRequest 释放FscRouteProductAddRequest
func ReleaseFscRouteProductAddRequest(v *FscRouteProductAddRequest) {
	v.RouteInfo = nil
	poolFscRouteProductAddRequest.Put(v)
}
