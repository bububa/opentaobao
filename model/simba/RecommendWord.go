package simba

// RecommendWord 
/* model for simplify = false
type RecommendWord struct {

    // 推荐的关键词
    
    Word   string `json:"word,omitempty"`
    

    // 搜索量
    
    Pv   string `json:"pv,omitempty"`
    

    // 平均价格
    
    AveragePrice   string `json:"average_price,omitempty"`
    

    // 相关度
    
    Pertinence   string `json:"pertinence,omitempty"`
    

}
*/

// RecommendWord 
type RecommendWord struct {

    // 推荐的关键词
    Word   string `json:"word,omitempty"`

    // 搜索量
    Pv   string `json:"pv,omitempty"`

    // 平均价格
    AveragePrice   string `json:"average_price,omitempty"`

    // 相关度
    Pertinence   string `json:"pertinence,omitempty"`

}
