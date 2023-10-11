package axindata

// FscRouteProductAddRequest 结构体
type FscRouteProductAddRequest struct {
	// 线路产品信息
	RouteInfo *FscRouteInfoApiDto `json:"route_info,omitempty" xml:"route_info,omitempty"`
}
