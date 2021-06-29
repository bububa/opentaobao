package traveltrade

// ScoreDetail 
type ScoreDetail struct {

    // 评价内容
    
    Content   string `json:"content,omitempty" xml:"content,omitempty"`
    

    // 评价分数
    
    Score   string `json:"score,omitempty" xml:"score,omitempty"`
    

    // 评价维度
    
    DimensionName   string `json:"dimension_name,omitempty" xml:"dimension_name,omitempty"`
    

}
