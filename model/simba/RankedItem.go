package simba

// RankedItem 结构体
type RankedItem struct {
	// 客户昵称
	Nick string `json:"nick,omitempty" xml:"nick,omitempty"`
	// 排名
	Order int64 `json:"order,omitempty" xml:"order,omitempty"`
	// 最高出价
	MaxPrice float64 `json:"max_price,omitempty" xml:"max_price,omitempty"`
	// 创意标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 商品链接
	LinkUrl string `json:"link_url,omitempty" xml:"link_url,omitempty"`
	// 原始质量评分
	RankScore int64 `json:"rank_score,omitempty" xml:"rank_score,omitempty"`
}
