package maitix

// ProjectPageParam 
type ProjectPageParam struct {
    // 查询页码
    PageNo   int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
    // 每页数据大小，可以稍微大一点
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}
