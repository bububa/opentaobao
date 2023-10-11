package axindata

// FscRouteProductUpdateRequest 结构体
type FscRouteProductUpdateRequest struct {
	// 线路产品信息
	RouteInfo *FscRouteInfoApiDto `json:"route_info,omitempty" xml:"route_info,omitempty"`
}
