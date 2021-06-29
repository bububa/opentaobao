package alihealthalgo

// SolutionResultTopSupport 
type SolutionResultTopSupport struct {
    // 结果
    DataList   []string `json:"data_list,omitempty" xml:"data_list>string,omitempty"`
    // 状态码
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
    // 错误信息
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
}
