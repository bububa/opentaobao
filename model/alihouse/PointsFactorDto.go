package alihouse

// PointsFactorDto 结构体
type PointsFactorDto struct {
	// 经纪人积分因素
	PointFactor string `json:"point_factor,omitempty" xml:"point_factor,omitempty"`
	// 经纪人积分分值
	PointScore int64 `json:"point_score,omitempty" xml:"point_score,omitempty"`
}
