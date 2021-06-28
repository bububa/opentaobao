package scbp

// ErrorKeyword 
/* model for simplify = false
type ErrorKeyword struct {

    // 关键词
    
    Keyword   string `json:"keyword,omitempty"`
    

    // 类型
    
    Type   string `json:"type,omitempty"`
    

    // 重复关键词
    
    RepeatKeyword   string `json:"repeat_keyword,omitempty"`
    

}
*/

// ErrorKeyword 
type ErrorKeyword struct {

    // 关键词
    Keyword   string `json:"keyword,omitempty"`

    // 类型
    Type   string `json:"type,omitempty"`

    // 重复关键词
    RepeatKeyword   string `json:"repeat_keyword,omitempty"`

}
