package simba

// BidStrategyVo 结构体
type BidStrategyVo struct {
	// 名称
	RankName string `json:"rank_name,omitempty" xml:"rank_name,omitempty"`
	// 溢价
	Dicount int64 `json:"dicount,omitempty" xml:"dicount,omitempty"`
	// 0停止，1生效
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 抢位排名
	Rank int64 `json:"rank,omitempty" xml:"rank,omitempty"`
}
