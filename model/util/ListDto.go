package util

// ListDto 
type ListDto struct {

    // 申请单总数
    Count   int64 `json:"count,omitempty"`

    // 申请单列表
    List   []Json `json:"list,omitempty"`

}
