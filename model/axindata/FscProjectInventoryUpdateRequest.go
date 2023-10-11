package axindata

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
