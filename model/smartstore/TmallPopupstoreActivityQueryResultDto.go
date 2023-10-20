package smartstore

// TmallpopupstoreactivityqueryResultDto 结构体
type TmallpopupstoreactivityqueryResultDto struct {
	// 返回结果
	ResultList []TmallpopupstoreactivityqueryResult `json:"result_list,omitempty" xml:"result_list>tmallpopupstoreactivityquery_result,omitempty"`
	// 返回code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误msg
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 返回数据条数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
}
