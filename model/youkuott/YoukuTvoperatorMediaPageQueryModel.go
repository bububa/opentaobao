package youkuott

// YoukuTvoperatorMediaPageQueryModel 
type YoukuTvoperatorMediaPageQueryModel struct {
    // 总条数
    TotalCount   int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
    // 是否有下一页
    HasNext   bool `json:"has_next,omitempty" xml:"has_next,omitempty"`
    // 页号
    PageNo   int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
    // 数据列表
    DataList   []YoukuTvoperatorMediaPageQueryData `json:"data_list,omitempty" xml:"data_list>youku_tvoperator_media_page_query_data,omitempty"`
    // 分页大小
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    // 总页数
    TotalPage   int64 `json:"total_page,omitempty" xml:"total_page,omitempty"`
}
