package wdk

// SyncActivityResultBo 结构体
type SyncActivityResultBo struct {
	// 门店ID
	StoreIds string `json:"store_ids,omitempty" xml:"store_ids,omitempty"`
	// 大润发活动ID
	PromotionId string `json:"promotion_id,omitempty" xml:"promotion_id,omitempty"`
	// 版本ID
	VersionId int64 `json:"version_id,omitempty" xml:"version_id,omitempty"`
}
