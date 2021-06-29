package omniorder

// OmniItemCategoryDetailDTO 
type OmniItemCategoryDetailDTO struct {
    // 内部
    InnerName   string `json:"inner_name,omitempty" xml:"inner_name,omitempty"`
    // 分类ID（添加时不太填，修改时要填）
    ClassifyId   int64 `json:"classify_id,omitempty" xml:"classify_id,omitempty"`
    // 展示名称
    DisplayName   string `json:"display_name,omitempty" xml:"display_name,omitempty"`
    // 是否在线上显示
    ShowOffline   bool `json:"show_offline,omitempty" xml:"show_offline,omitempty"`
    // 是否在线下显示
    ShowOnline   bool `json:"show_online,omitempty" xml:"show_online,omitempty"`
}
