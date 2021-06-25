package hotel

// RoomTypeBedInfoVo 
type RoomTypeBedInfoVo struct {

    // 或关系床型集合
    BedInfoGroups   []BedInfoGroupVo `json:"bed_info_groups,omitempty"`

    // 简短描述，用于详情页报价前面的床型展示。
    BriefDesc   string `json:"brief_desc,omitempty"`

    // 分类，大类，用于搜索的筛选项
    Classifications   []String `json:"classifications,omitempty"`

    // 分类描述
    ClassificationDesc   string `json:"classification_desc,omitempty"`

    // 描述，长描述，描述为最详细的内容，用于详情页房型详情Dialog，下单页，订单页展示
    Desc   string `json:"desc,omitempty"`

    // 模糊描述，最精简的描述信息，大床/双床/单人床/多床/床位简单描述，较长描述省略床宽
    FuzzyDesc   string `json:"fuzzy_desc,omitempty"`

    // 简单描述，较长描述省略床宽，但依然会描述具体的床型信息，用于详情页标准房型床型展示
    SimpleDesc   string `json:"simple_desc,omitempty"`

}
