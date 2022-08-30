package axindata

// MatchedRoomDataDto 结构体
type MatchedRoomDataDto struct {
	// 分值,越高可信度越高
	Score string `json:"score,omitempty" xml:"score,omitempty"`
	// 匹配结果标准房型id
	Srid int64 `json:"srid,omitempty" xml:"srid,omitempty"`
}
