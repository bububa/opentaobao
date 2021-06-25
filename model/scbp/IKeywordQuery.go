package scbp

// IKeywordQuery 
type IKeywordQuery struct {

    // 每页行数
    PerPageSize   int64 `json:"per_page_size,omitempty"`

    // 第几页
    ToPage   int64 `json:"to_page,omitempty"`

    // 结束时间 当inteval=7或者30的时候不需要填写
    EndDate   string `json:"end_date,omitempty"`

    // 区间 只能为1 7 30
    Inteval   int64 `json:"inteval,omitempty"`

    // 关键词
    Keyword   string `json:"keyword,omitempty"`

    // 开始时间 当inteval=7或者30的时候不需要填写
    BeginDate   string `json:"begin_date,omitempty"`

}
