package hotel

// ScoreInfo 结构体
type ScoreInfo struct {
	// desc
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// label
	Label string `json:"label,omitempty" xml:"label,omitempty"`
	// score
	Score string `json:"score,omitempty" xml:"score,omitempty"`
	// count
	Count int64 `json:"count,omitempty" xml:"count,omitempty"`
}
