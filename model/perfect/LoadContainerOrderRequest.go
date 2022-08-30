package perfect

// LoadContainerOrderRequest 结构体
type LoadContainerOrderRequest struct {
	// 包裹
	Packages []LoadPackageOrderRequest `json:"packages,omitempty" xml:"packages>load_package_order_request,omitempty"`
	// 装箱单
	ContainerOrderCode string `json:"container_order_code,omitempty" xml:"container_order_code,omitempty"`
	// 箱单类型，NORMAL_BOX(正常箱)，OUTER_BOX(箱外件)，IRREGULAR_BOX(异形件)
	ContainerOrderType string `json:"container_order_type,omitempty" xml:"container_order_type,omitempty"`
	// 扩展
	Attributes string `json:"attributes,omitempty" xml:"attributes,omitempty"`
	// 周转箱号
	ContainerCode string `json:"container_code,omitempty" xml:"container_code,omitempty"`
	// 站点
	StationCode string `json:"station_code,omitempty" xml:"station_code,omitempty"`
	// 站点名称
	StationName string `json:"station_name,omitempty" xml:"station_name,omitempty"`
}
