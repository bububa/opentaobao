package alihouse

// SyncBrokerPointsDto 结构体
type SyncBrokerPointsDto struct {
	// 经纪人积分因素列表
	PointsFactors []PointsFactorDto `json:"points_factors,omitempty" xml:"points_factors>points_factor_dto,omitempty"`
	// 外部经纪人ID
	OuterConsultantId string `json:"outer_consultant_id,omitempty" xml:"outer_consultant_id,omitempty"`
	// 业务数据归属时间
	BusinessTime string `json:"business_time,omitempty" xml:"business_time,omitempty"`
	// 0默认不是测试，1是测试数据
	IsTest int64 `json:"is_test,omitempty" xml:"is_test,omitempty"`
}
