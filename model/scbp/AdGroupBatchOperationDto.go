package scbp

// AdGroupBatchOperationDto 结构体
type AdGroupBatchOperationDto struct {
	// 入参
	AdGroupOperationList []AdGroupOperationDto `json:"ad_group_operation_list,omitempty" xml:"ad_group_operation_list>ad_group_operation_dto,omitempty"`
}
