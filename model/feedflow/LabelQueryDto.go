package feedflow

// LabelQueryDTO 
type LabelQueryDTO struct {
    // 选项值
    OptionName   string `json:"option_name,omitempty" xml:"option_name,omitempty"`
    // 分页条件
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    // 定向id
    TargetId   int64 `json:"target_id,omitempty" xml:"target_id,omitempty"`
    // 分页条件
    Offset   int64 `json:"offset,omitempty" xml:"offset,omitempty"`
    // 定向类型
    TargetType   string `json:"target_type,omitempty" xml:"target_type,omitempty"`
    // 宝贝id列表
    ItemIdList   []int64 `json:"item_id_list,omitempty" xml:"item_id_list>int64,omitempty"`
}
