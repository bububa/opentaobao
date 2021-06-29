package ott

// Notebtns 
type Notebtns struct {
    // 便签类型
    Type   string `json:"type,omitempty" xml:"type,omitempty"`
    // 便签名称
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    // 排序
    Sort   int64 `json:"sort,omitempty" xml:"sort,omitempty"`
}
