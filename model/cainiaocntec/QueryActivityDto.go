package cainiaocntec

// QueryActivityDto 
type QueryActivityDto struct {
    // 门店id
    StoreId   string `json:"store_id,omitempty" xml:"store_id,omitempty"`
    // 是否分页
    Page   bool `json:"page,omitempty" xml:"page,omitempty"`
    // 分页大小
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    // 当前页
    PageIndex   int64 `json:"page_index,omitempty" xml:"page_index,omitempty"`
}
