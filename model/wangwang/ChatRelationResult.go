package wangwang

// ChatRelationResult 
type ChatRelationResult struct {

    // 返回码
    Code   int64 `json:"code,omitempty"`

    // 关系列表
    Peers   []ChatPeer `json:"peers,omitempty"`

    // 错误原因
    Reason   string `json:"reason,omitempty"`

    // 起始key。如果要实现上翻页，可以将该值作为下次请求的page_end
    StartKey   string `json:"start_key,omitempty"`

    // 结束key。如果该值为空，则表示请求时间区间内的数据已经拿完。否则，表示区间内还有数据，可以将该值作为下次请求条件的page_beg传入进行迭代请求。
    EndKey   string `json:"end_key,omitempty"`

}
