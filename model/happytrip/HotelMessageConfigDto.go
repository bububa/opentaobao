package happytrip

// HotelMessageConfigDto 结构体
type HotelMessageConfigDto struct {
	// 配置信息
	Segments []string `json:"segments,omitempty" xml:"segments>string,omitempty"`
}
