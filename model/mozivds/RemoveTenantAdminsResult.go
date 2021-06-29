package mozivds

// RemoveTenantAdminsResult 
type RemoveTenantAdminsResult struct {
    // 是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 请求id
    RequestId   string `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回消息
    ResponseMessage   string `json:"response_message,omitempty" xml:"response_message,omitempty"`
    // 返回元数据
    ResponseMetaData   string `json:"response_meta_data,omitempty" xml:"response_meta_data,omitempty"`
    // 返回Code
    ResponseCode   string `json:"response_code,omitempty" xml:"response_code,omitempty"`
}
