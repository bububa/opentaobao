package hotel

// BedInfoGroupVo 
type BedInfoGroupVo struct {

    // 和关系床型集合
    
    BedInfos   []BedInfoVo `json:"bed_infos,omitempty" xml:"bed_infos,omitempty"`
    

    // 床数量
    
    BedNum   int64 `json:"bed_num,omitempty" xml:"bed_num,omitempty"`
    

    // 简短描述，用于详情页报价前面的床型展示
    
    BriefDesc   string `json:"brief_desc,omitempty" xml:"brief_desc,omitempty"`
    

    // 分类，大类，用于搜索的筛选项。1:大床,2:双床,3:单人床,4:多床
    
    Classification   string `json:"classification,omitempty" xml:"classification,omitempty"`
    

    // 分类描述
    
    ClassificationDesc   string `json:"classification_desc,omitempty" xml:"classification_desc,omitempty"`
    

    // 描述，长描述，描述为最详细的内容，用于详情页房型详情Dialog，下单页，订单页展示
    
    Desc   string `json:"desc,omitempty" xml:"desc,omitempty"`
    

    // 模糊描述，最精简的描述信息，大床/双床/单人床/多床/床位简单描述，较长描述省略床宽
    
    FuzzyDesc   string `json:"fuzzy_desc,omitempty" xml:"fuzzy_desc,omitempty"`
    

    // 简单描述，较长描述省略床宽，但依然会描述具体的床型信息，用于详情页标准房型床型展示
    
    SimpleDesc   string `json:"simple_desc,omitempty" xml:"simple_desc,omitempty"`
    

}
