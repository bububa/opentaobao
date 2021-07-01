package logistic

// ReachableServiceWaybillForTopResponseDto 结构体
type ReachableServiceWaybillForTopResponseDto struct {
	// 结果列表
	ResultList []ReachableServiceWaybillResponseDto `json:"result_list,omitempty" xml:"result_list>reachable_service_waybill_response_dto,omitempty"`
}
