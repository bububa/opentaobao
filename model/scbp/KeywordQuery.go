package scbp

// KeywordQuery 
type KeywordQuery struct {

    // 请求实体集合
    KeywordList   []KeywordInfo `json:"keyword_list,omitempty"`

    // 指定查询开始的具体日期，最多往前推4周
    FromDate   string `json:"from_date,omitempty"`

    // 按关键词状态查询，只能取值stopped和in_promotion
    Status   string `json:"status,omitempty"`

    // 按星级查询,可取0,1,2,3,4,5
    QsStar   string `json:"qs_star,omitempty"`

    // 每页个数
    PerPageSize   int64 `json:"per_page_size,omitempty"`

    // 按关键词分组名称查询
    TagName   string `json:"tag_name,omitempty"`

    // 按关键词查询
    Keyword   string `json:"keyword,omitempty"`

    // 指定第几页
    ToPage   int64 `json:"to_page,omitempty"`

    // 是否精确查询
    IsExact   string `json:"is_exact,omitempty"`

    // 按最近1天，7天，30天查询
    Inteval   int64 `json:"inteval,omitempty"`

}
