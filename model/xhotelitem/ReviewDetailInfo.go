package xhotelitem

// ReviewDetailInfo 结构体
type ReviewDetailInfo struct {
	// 评分，10分制，Double类型得，2-非常差 4-差 6-一般 8-好 10-非常好
	Score string `json:"score,omitempty" xml:"score,omitempty"`
	// 纬度id，1-地理位置 2-清洁程度 3-服务比较 4-性价比
	DimensionId int64 `json:"dimension_id,omitempty" xml:"dimension_id,omitempty"`
}
