package happytrip

// HotelMessageConfigDto 结构体
type HotelMessageConfigDto struct {
	// 配置信息
	Segments []Null `json:"segments,omitempty" xml:"segments>null,omitempty"`
}
