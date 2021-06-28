package waybill

// WaybillCloudPrintWithResultDescResponse 
/* model for simplify = false
type WaybillCloudPrintWithResultDescResponse struct {

    // 单查询错误code
    
    ErrorCode   string `json:"error_code,omitempty"`
    

    // 单查询错误message
    
    ErrorMessage   string `json:"error_message,omitempty"`
    

    // 请求id
    
    ObjectId   string `json:"object_id,omitempty"`
    

    // 是否查询
    
    Success   bool `json:"success,omitempty"`
    

    // 面单信息
    
    WaybillCloudPrintResponse  *struct {
        WaybillCloudPrintResponse  *WaybillCloudPrintResponse `json:"waybill_cloud_print_response,omitempty"`
    } `json:"waybill_cloud_print_response,omitempty"`
    

}
*/

// WaybillCloudPrintWithResultDescResponse 
type WaybillCloudPrintWithResultDescResponse struct {

    // 单查询错误code
    ErrorCode   string `json:"error_code,omitempty"`

    // 单查询错误message
    ErrorMessage   string `json:"error_message,omitempty"`

    // 请求id
    ObjectId   string `json:"object_id,omitempty"`

    // 是否查询
    Success   bool `json:"success,omitempty"`

    // 面单信息
    WaybillCloudPrintResponse   *WaybillCloudPrintResponse `json:"waybill_cloud_print_response,omitempty"`

}
