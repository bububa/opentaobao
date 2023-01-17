package logistic

// Location 结构体
type Location struct {
	// 经度
	Longitude string `json:"longitude,omitempty" xml:"longitude,omitempty"`
	// 纬度
	Latitude string `json:"latitude,omitempty" xml:"latitude,omitempty"`
}
