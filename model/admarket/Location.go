package admarket

// Location 结构体
type Location struct {
	// 定位类型(WGS84/GCJ02/BD09)
	CoordinateType string `json:"coordinate_type,omitempty" xml:"coordinate_type,omitempty"`
	// 经度
	Longitude string `json:"longitude,omitempty" xml:"longitude,omitempty"`
	// 纬度
	Latitude string `json:"latitude,omitempty" xml:"latitude,omitempty"`
}
