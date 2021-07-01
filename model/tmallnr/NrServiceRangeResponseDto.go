package tmallnr

// NrServiceRangeResponseDto 结构体
type NrServiceRangeResponseDto struct {
	// 围栏数据
	Points []Points `json:"points,omitempty" xml:"points>points,omitempty"`
}
