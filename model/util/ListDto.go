package util

// ListDto 
type ListDto struct {
    // 申请单总数
    Count   int64 `json:"count,omitempty" xml:"count,omitempty"`
    // 申请单列表
    List   []string `json:"list,omitempty" xml:"list>string,omitempty"`
}
