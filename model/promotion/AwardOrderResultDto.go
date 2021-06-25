package promotion

// AwardOrderResultDto 
type AwardOrderResultDto struct {

    // 订单列表
    Orders   []AwardOrder `json:"orders,omitempty"`

    // 页码
    PageNo   int64 `json:"page_no,omitempty"`

    // 页数
    PageSize   int64 `json:"page_size,omitempty"`

    // 总页数
    TotalCount   int64 `json:"total_count,omitempty"`

}
