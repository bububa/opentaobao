package btrip

// OpenCostCenterTransferRq 结构体
type OpenCostCenterTransferRq struct {
	// 第三方成本中心id
	ThirdpartId string `json:"thirdpart_id,omitempty" xml:"thirdpart_id,omitempty"`
	// 商旅成本中心id
	CostCenterId int64 `json:"cost_center_id,omitempty" xml:"cost_center_id,omitempty"`
	// 企业id
	CorpId string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	// 商旅开放平台传2
	Version int64 `json:"version,omitempty" xml:"version,omitempty"`
}
