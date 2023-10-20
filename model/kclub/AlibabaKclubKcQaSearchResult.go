package kclub

// AlibabakclubkcqasearchResult 结构体
type AlibabakclubkcqasearchResult struct {
	// 返回数据列表
	DataList []KcSearchQuestion `json:"data_list,omitempty" xml:"data_list>kc_search_question,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误编码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
