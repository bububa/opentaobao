package eleenterpriseordernew

// TrackingInfoDto 结构体
type TrackingInfoDto struct {
	// 纬度
	Latitude string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	// 经度
	Longitude string `json:"longitude,omitempty" xml:"longitude,omitempty"`
}
