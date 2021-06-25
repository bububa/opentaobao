package tbk

// PageResult 
type PageResult struct {

    // 翻页的pageno
    PageNo   int64 `json:"page_no,omitempty"`

    // 翻页的pagesie
    PageSize   int64 `json:"page_size,omitempty"`

    // 一共能查询出来的结果总数
    TotalCount   int64 `json:"total_count,omitempty"`

    // 处罚订单列表
    Results   []TaobaoTbkDgPunishOrderGetResult `json:"results,omitempty"`

}
