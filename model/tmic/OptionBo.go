package tmic

// OptionBo 
type OptionBo struct {
    // 是否还有其他选项
    HasOther   bool `json:"has_other,omitempty" xml:"has_other,omitempty"`
    // 选项具体值
    Items   []ItemBo `json:"items,omitempty" xml:"items>item_bo,omitempty"`
    // optionItemBOList
    OptionItemBoList   []OptionItemBo `json:"option_item_bo_list,omitempty" xml:"option_item_bo_list>option_item_bo,omitempty"`
}
