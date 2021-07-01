package tvupadmin

// PaginationDo 
type PaginationDo struct {
    // 总数
    Total   int64 `json:"total,omitempty" xml:"total,omitempty"`
    // 页码
    PageNo   int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
    // 单页数量
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    // 内容列表
    List   []AdvertScheduleDo `json:"list,omitempty" xml:"list>advert_schedule_do,omitempty"`
}
