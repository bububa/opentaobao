package logistic

// ReachableServiceWaybillForTopResponseDto 
/* model for simplify = false
type ReachableServiceWaybillForTopResponseDto struct {

    // 结果列表
    
    ResultList  struct {
        ReachableServiceWaybillResponseDto  []ReachableServiceWaybillResponseDto `json:"reachable_service_waybill_response_dto,omitempty"`
    } `json:"result_list,omitempty"`
    

}
*/

// ReachableServiceWaybillForTopResponseDto 
type ReachableServiceWaybillForTopResponseDto struct {

    // 结果列表
    ResultList   []ReachableServiceWaybillResponseDto `json:"result_list,omitempty"`

}
