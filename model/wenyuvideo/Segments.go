package wenyuvideo

// Segments 结构体
type Segments struct {
	// 开始时间点
	From int64 `json:"from,omitempty" xml:"from,omitempty"`
	// 结束时间点
	To int64 `json:"to,omitempty" xml:"to,omitempty"`
}
