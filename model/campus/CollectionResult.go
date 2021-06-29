package campus

// CollectionResult 
type CollectionResult struct {
    // 菜单内容
    Contents   []TreeNode `json:"contents,omitempty" xml:"contents>tree_node,omitempty"`
    // 错误码
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // 错误描述
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
    // 是否调用成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 错误级别
    ErrorLevel   string `json:"error_level,omitempty" xml:"error_level,omitempty"`
    // content
    ContentList   []AlibabaCampusCoreAppGetappusagesT `json:"content_list,omitempty" xml:"content_list>alibaba_campus_core_app_getappusages_t,omitempty"`
    // requestId
    RequestId   string `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // totalCount
    TotalCount   int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
    // errorExtInfo
    ErrorExtInfo   string `json:"error_ext_info,omitempty" xml:"error_ext_info,omitempty"`
}
