package wangwang

// ChatRelationRequest 
type ChatRelationRequest struct {
    // 查询账号
    Uid   string `json:"uid,omitempty" xml:"uid,omitempty"`
    // 查询起始时间（精度到天）
    Beg   string `json:"beg,omitempty" xml:"beg,omitempty"`
    // 查询结束时间（精度到天）
    End   string `json:"end,omitempty" xml:"end,omitempty"`
    // 查询条数
    Count   int64 `json:"count,omitempty" xml:"count,omitempty"`
    // 翻页查询起始key
    PageBeg   string `json:"page_beg,omitempty" xml:"page_beg,omitempty"`
    // 翻页查询结束key
    PageEnd   string `json:"page_end,omitempty" xml:"page_end,omitempty"`
}
