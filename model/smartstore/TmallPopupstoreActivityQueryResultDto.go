package smartstore

// TmallPopupstoreActivityQueryResultDto 结构体
type TmallPopupstoreActivityQueryResultDto struct {
	// 返回数据条数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// 返回code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误msg
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 返回结果
	ResultList []TmallPopupstoreActivityQueryResult `json:"result_list,omitempty" xml:"result_list>tmall_popupstore_activity_query_result,omitempty"`
}
