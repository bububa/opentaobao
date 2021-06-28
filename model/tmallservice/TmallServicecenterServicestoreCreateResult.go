package tmallservice

// TmallServicecenterServicestoreCreateResult 
/* model for simplify = false
type TmallServicecenterServicestoreCreateResult struct {

    // 错误码
    
    MsgCode   string `json:"msg_code,omitempty"`
    

    // 错误信息
    
    MsgInfo   string `json:"msg_info,omitempty"`
    

    // 门店数据
    
    ResultData  *struct {
        ServiceStoreCreateResp  *ServiceStoreCreateResp `json:"service_store_create_resp,omitempty"`
    } `json:"result_data,omitempty"`
    

    // 调用是否成功
    
    Success   bool `json:"success,omitempty"`
    

}
*/

// TmallServicecenterServicestoreCreateResult 
type TmallServicecenterServicestoreCreateResult struct {

    // 错误码
    MsgCode   string `json:"msg_code,omitempty"`

    // 错误信息
    MsgInfo   string `json:"msg_info,omitempty"`

    // 门店数据
    ResultData   *ServiceStoreCreateResp `json:"result_data,omitempty"`

    // 调用是否成功
    Success   bool `json:"success,omitempty"`

}
