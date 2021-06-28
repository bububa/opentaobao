package tmallservice

// ResultBase 
/* model for simplify = false
type ResultBase struct {

    // 单个一键求助单对象json字符串
    
    Value   string `json:"value,omitempty"`
    

    // errorCode
    
    ErrorCode   int64 `json:"error_code,omitempty"`
    

    // errorMsg
    
    ErrorMsg   string `json:"error_msg,omitempty"`
    

    // success
    
    Success   bool `json:"success,omitempty"`
    

    // gmtCreate
    
    GmtCreate   string `json:"gmt_create,omitempty"`
    

    // gmtModified
    
    GmtModified   string `json:"gmt_modified,omitempty"`
    

    // 任务id列表
    
    Values  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"values,omitempty"`
    

    // 服务预警消息列表
    
    ValueList  struct {
        ServiceMonitorMessage  []ServiceMonitorMessage `json:"service_monitor_message,omitempty"`
    } `json:"value_list,omitempty"`
    

}
*/

// ResultBase 
type ResultBase struct {

    // 单个一键求助单对象json字符串
    Value   string `json:"value,omitempty"`

    // errorCode
    ErrorCode   int64 `json:"error_code,omitempty"`

    // errorMsg
    ErrorMsg   string `json:"error_msg,omitempty"`

    // success
    Success   bool `json:"success,omitempty"`

    // gmtCreate
    GmtCreate   string `json:"gmt_create,omitempty"`

    // gmtModified
    GmtModified   string `json:"gmt_modified,omitempty"`

    // 任务id列表
    Values   []int64 `json:"values,omitempty"`

    // 服务预警消息列表
    ValueList   []ServiceMonitorMessage `json:"value_list,omitempty"`

}
