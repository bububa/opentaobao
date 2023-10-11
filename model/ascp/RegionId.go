package ascp

// RegionId 结构体
type RegionId struct {
	// 经度（高德）
	Latitude string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	// 纬度（高德）
	Longitude string `json:"longitude,omitempty" xml:"longitude,omitempty"`
}
