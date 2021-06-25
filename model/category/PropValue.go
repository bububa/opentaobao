package category

// PropValue 
type PropValue struct {

    // 类目ID
    Cid   int64 `json:"cid,omitempty"`

    // 属性 ID
    Pid   int64 `json:"pid,omitempty"`

    // 属性名
    PropName   string `json:"prop_name,omitempty"`

    // 属性值ID
    Vid   int64 `json:"vid,omitempty"`

    // 属性值
    Name   string `json:"name,omitempty"`

    // 属性值别名
    NameAlias   string `json:"name_alias,omitempty"`

    // 是否为父类目属性
    IsParent   bool `json:"is_parent,omitempty"`

    // 状态。可选值:normal(正常),deleted(删除)
    Status   string `json:"status,omitempty"`

    // 排列序号。取值范围:大于零的整数
    SortOrder   int64 `json:"sort_order,omitempty"`

    // 属性值feature
    Features   []Feature `json:"features,omitempty"`

}
