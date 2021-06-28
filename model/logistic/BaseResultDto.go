package logistic

// BaseResultDto 
/* model for simplify = false
type BaseResultDto struct {

    // 请求是否成功
    
    Success   bool `json:"success,omitempty"`
    

    // 返回信息
    
    Module  *struct {
        ReachableServiceWaybillForTopResponseDto  *ReachableServiceWaybillForTopResponseDto `json:"reachable_service_waybill_for_top_response_dto,omitempty"`
    } `json:"module,omitempty"`
    

    // 请求错误信息
    
    OneErrorInfo  *struct {
        ErrorInfo  *ErrorInfo `json:"error_info,omitempty"`
    } `json:"one_error_info,omitempty"`
    

}
*/

// BaseResultDto 
type BaseResultDto struct {

    // 请求是否成功
    Success   bool `json:"success,omitempty"`

    // 返回信息
    Module   *ReachableServiceWaybillForTopResponseDto `json:"module,omitempty"`

    // 请求错误信息
    OneErrorInfo   *ErrorInfo `json:"one_error_info,omitempty"`

}
