package category

// AlibabaImapFixedmappingQueryResult 结构体
type AlibabaImapFixedmappingQueryResult struct {
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// list参数
	TopImapUnionCategoryPathDoList []TopImapUnionCategoryPathDo `json:"top_imap_union_category_path_do_list,omitempty" xml:"top_imap_union_category_path_do_list>top_imap_union_category_path_do,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
}
