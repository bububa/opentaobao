package kclub

// AlibabaKclubKcQaSearchResult 
type AlibabaKclubKcQaSearchResult struct {

    // 错误信息
    Message   string `json:"message,omitempty"`

    // 返回数据列表
    DataList   []KcSearchQuestion `json:"data_list,omitempty"`

    // 错误编码
    Code   string `json:"code,omitempty"`

    // 是否成功
    Success   bool `json:"success,omitempty"`

}