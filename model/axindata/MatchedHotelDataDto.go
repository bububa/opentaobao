package axindata

// MatchedHotelDataDto 结构体
type MatchedHotelDataDto struct {
	// 分值
	Score string `json:"score,omitempty" xml:"score,omitempty"`
	// 匹配的标准酒店id
	Shid int64 `json:"shid,omitempty" xml:"shid,omitempty"`
}
