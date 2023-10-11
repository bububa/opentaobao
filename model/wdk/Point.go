package wdk

// Point 结构体
type Point struct {
	// 经度(高德坐标)
	Lng string `json:"lng,omitempty" xml:"lng,omitempty"`
	// 纬度(高德坐标)
	Lat string `json:"lat,omitempty" xml:"lat,omitempty"`
}
