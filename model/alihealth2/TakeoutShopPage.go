package alihealth2

// TakeoutShopPage 
type TakeoutShopPage struct {
    // 总条数
    TotalCount   int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
    // 当前页码
    Page   int64 `json:"page,omitempty" xml:"page,omitempty"`
    // 每页条数
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    // 总页数
    TotalPage   int64 `json:"total_page,omitempty" xml:"total_page,omitempty"`
    // 上一页页码
    PrevPage   int64 `json:"prev_page,omitempty" xml:"prev_page,omitempty"`
    // 下一页页码
    NextPage   int64 `json:"next_page,omitempty" xml:"next_page,omitempty"`
    // 店铺信息列表
    TakeoutSummaryInfos   []TakeoutShopSummaryInfo `json:"takeout_summary_infos,omitempty" xml:"takeout_summary_infos>takeout_shop_summary_info,omitempty"`
}
