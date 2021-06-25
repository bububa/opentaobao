package yunos

// RdamResponse 
type RdamResponse struct {

    // message
    Message   string `json:"message,omitempty"`

    // more
    More   string `json:"more,omitempty"`

    // code
    Code   int64 `json:"code,omitempty"`

    // dataList
    DataList   []Json `json:"data_list,omitempty"`

    // traceId
    TraceId   string `json:"trace_id,omitempty"`

}
