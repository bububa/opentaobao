package ascp

// RegionIds 结构体
type RegionIds struct {
	// 经度（高德）
	Longitude string `json:"longitude,omitempty" xml:"longitude,omitempty"`
	// 纬度（高德）
	Latitude string `json:"latitude,omitempty" xml:"latitude,omitempty"`
}
