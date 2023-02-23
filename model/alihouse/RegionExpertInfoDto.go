package alihouse

// RegionExpertInfoDto 结构体
type RegionExpertInfoDto struct {
	// 外部经纪人ID
	OuterConsultantId string `json:"outer_consultant_id,omitempty" xml:"outer_consultant_id,omitempty"`
	// 外部门店ID
	OuterStoreId string `json:"outer_store_id,omitempty" xml:"outer_store_id,omitempty"`
	// 经纪人状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}
