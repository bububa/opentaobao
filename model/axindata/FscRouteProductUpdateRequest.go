package axindata

import (
	"sync"
)

// FscRouteProductUpdateRequest 结构体
type FscRouteProductUpdateRequest struct {
	// 线路产品信息
	RouteInfo *FscRouteInfoApiDto `json:"route_info,omitempty" xml:"route_info,omitempty"`
}

var poolFscRouteProductUpdateRequest = sync.Pool{
	New: func() any {
		return new(FscRouteProductUpdateRequest)
	},
}

// GetFscRouteProductUpdateRequest() 从对象池中获取FscRouteProductUpdateRequest
func GetFscRouteProductUpdateRequest() *FscRouteProductUpdateRequest {
	return poolFscRouteProductUpdateRequest.Get().(*FscRouteProductUpdateRequest)
}

// ReleaseFscRouteProductUpdateRequest 释放FscRouteProductUpdateRequest
func ReleaseFscRouteProductUpdateRequest(v *FscRouteProductUpdateRequest) {
	v.RouteInfo = nil
	poolFscRouteProductUpdateRequest.Put(v)
}
