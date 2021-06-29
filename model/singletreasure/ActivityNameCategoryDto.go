package singletreasure

// ActivityNameCategoryDto 
type ActivityNameCategoryDto struct {

    // 活动描述   1:日常活动  2:官方活动
    
    Text   string `json:"text,omitempty" xml:"text,omitempty"`
    

    // 活动value
    
    Value   int64 `json:"value,omitempty" xml:"value,omitempty"`
    

    // 名称列表
    
    List   []ActivityNameInfoDto `json:"list,omitempty" xml:"list,omitempty"`
    

}
