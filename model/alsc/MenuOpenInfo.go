package alsc

// MenuOpenInfo 
type MenuOpenInfo struct {

    // 生效时段结束
    EffectEnd   string `json:"effect_end,omitempty"`

    // 生效时段起始
    EffectStart   string `json:"effect_start,omitempty"`

    // 特价菜单id
    MenuId   string `json:"menu_id,omitempty"`

    // 特价菜单名称
    Name   string `json:"name,omitempty"`

    // 统一折扣价
    ProDiscount   string `json:"pro_discount,omitempty"`

    // 特价模式：统一折扣、不同菜不同折扣
    ProMode   string `json:"pro_mode,omitempty"`

    // 是否逻辑删除
    Deleted   bool `json:"deleted,omitempty"`

    // 菜品集合
    MenuDetailOpenInfoList   []MenuDetailOpenInfo `json:"menu_detail_open_info_list,omitempty"`

    // 更新时间
    GmtModified   string `json:"gmt_modified,omitempty"`

    // 创建时间
    GmtCreate   string `json:"gmt_create,omitempty"`

}
