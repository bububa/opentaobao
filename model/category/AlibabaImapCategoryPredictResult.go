package category

// AlibabaImapCategoryPredictResult 结构体
type AlibabaImapCategoryPredictResult struct {
	// 1
	TopImapUnionCategoryPathDoList []TopImapUnionCategoryPathDo `json:"top_imap_union_category_path_do_list,omitempty" xml:"top_imap_union_category_path_do_list>top_imap_union_category_path_do,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
