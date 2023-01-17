package alihouse

// RegionExpertDto 结构体
type RegionExpertDto struct {
	// 经纪人信息
	OuterConsultantInfos []RegionExpertInfoDto `json:"outer_consultant_infos,omitempty" xml:"outer_consultant_infos>region_expert_info_dto,omitempty"`
	// 外部区域/商圈ID
	OuterRegionId string `json:"outer_region_id,omitempty" xml:"outer_region_id,omitempty"`
	// 类型400-商圈
	RegionType int64 `json:"region_type,omitempty" xml:"region_type,omitempty"`
	// 0-正式数据 1-测试数据
	IsTest int64 `json:"is_test,omitempty" xml:"is_test,omitempty"`
	// 版本号
	Version int64 `json:"version,omitempty" xml:"version,omitempty"`
	// 负责业务1-新房 2-二手房
	BusinessType int64 `json:"business_type,omitempty" xml:"business_type,omitempty"`
}
