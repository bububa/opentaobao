package hotel

// SearchPoi 结构体
type SearchPoi struct {
	// 经度
	Latitude string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	// 纬度
	Longitude string `json:"longitude,omitempty" xml:"longitude,omitempty"`
	// 名字
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}
