package icbulogistics

// Division 
type Division struct {
    // 节点名称拼音
    Pinyin   string `json:"pinyin,omitempty" xml:"pinyin,omitempty"`
    // 上级节点ID
    ParentId   int64 `json:"parent_id,omitempty" xml:"parent_id,omitempty"`
    // 中文名称
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    // 层级
    Level   int64 `json:"level,omitempty" xml:"level,omitempty"`
    // 节点id
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
}
