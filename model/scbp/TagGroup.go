package scbp

// TagGroup 
type TagGroup struct {

    // 分组名称
    Name   string `json:"name,omitempty"`

    // 分组ID，不对外暴露
    Id   int64 `json:"id,omitempty"`

    // 关键词数
    Count   int64 `json:"count,omitempty"`

}
