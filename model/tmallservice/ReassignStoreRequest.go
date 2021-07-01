package tmallservice

// ReassignStoreRequest 结构体
type ReassignStoreRequest struct {
	// 工单id
	WorkcardId int64 `json:"workcard_id,omitempty" xml:"workcard_id,omitempty"`
	// 目标门店id
	TargetServiceStoreId int64 `json:"target_service_store_id,omitempty" xml:"target_service_store_id,omitempty"`
	// 原门店id
	OriginServiceStoreId int64 `json:"origin_service_store_id,omitempty" xml:"origin_service_store_id,omitempty"`
}
