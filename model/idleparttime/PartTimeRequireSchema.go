package idleparttime

// PartTimeRequireSchema 
type PartTimeRequireSchema struct {
    // 要求
    Description   string `json:"description,omitempty" xml:"description,omitempty"`
    // 类型, 1:文本 2: 图片
    Type   int64 `json:"type,omitempty" xml:"type,omitempty"`
    // 要求描述
    Requirement   string `json:"requirement,omitempty" xml:"requirement,omitempty"`
}
