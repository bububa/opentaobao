package tvupadmin

// PageVo 结构体
type PageVo struct {
	// 总数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// 内容列表
	RecordList []ChildNodeContentVo `json:"record_list,omitempty" xml:"record_list>child_node_content_vo,omitempty"`
	// 页码
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 单页数量
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}
