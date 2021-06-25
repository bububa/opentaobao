package promotion

// DataServiceResponse 
type DataServiceResponse struct {

    // 扩展字段
    ExtMap   *Extmap `json:"ext_map,omitempty"`

    // 结果对象
    Result   bool `json:"result,omitempty"`

    // 报错编码
    ResultCode   string `json:"result_code,omitempty"`

    // 报错描述
    ResultMsg   string `json:"result_msg,omitempty"`

    // 调用是否成功
    Success   bool `json:"success,omitempty"`

    // 鹰眼id
    TraceId   string `json:"trace_id,omitempty"`

}
