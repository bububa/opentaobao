package examination

// ServiceResult 
type ServiceResult struct {
    // 错误信息
    ErrMessage   string `json:"err_message,omitempty" xml:"err_message,omitempty"`
    // 返回数据,一般为空
    Data   string `json:"data,omitempty" xml:"data,omitempty"`
    // 错误码
    ErrCode   string `json:"err_code,omitempty" xml:"err_code,omitempty"`
    // 执行是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // eagleEyeTraceId
    EagleEyeTraceId   string `json:"eagle_eye_trace_id,omitempty" xml:"eagle_eye_trace_id,omitempty"`
    // 返回数据对象
    DataList   []StoreSpecialTagsResponse `json:"data_list,omitempty" xml:"data_list>store_special_tags_response,omitempty"`
}
