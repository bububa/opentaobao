package aliospay

// SearchRecordRequest 
type SearchRecordRequest struct {
    // 请求唯一id，不可重复，服务端会根据此参数防重放
    TraceId   string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
    // 语言,en表示英文，zh表示中文
    Lang   string `json:"lang,omitempty" xml:"lang,omitempty"`
    // 发送请求的时间戳
    Time   string `json:"time,omitempty" xml:"time,omitempty"`
    // 起始时间，时间戳
    BeginTime   int64 `json:"begin_time,omitempty" xml:"begin_time,omitempty"`
    // 结束时间，时间戳
    EndTime   int64 `json:"end_time,omitempty" xml:"end_time,omitempty"`
    // 分页偏移量，默认0
    Offset   int64 `json:"offset,omitempty" xml:"offset,omitempty"`
    // 每页数量，默认100
    Size   int64 `json:"size,omitempty" xml:"size,omitempty"`
}
