package security

// JaqAccountRiskDetailItem 
/* model for simplify = false
type JaqAccountRiskDetailItem struct {

    // 决定，0：可以接受；1：应该拒绝；2：需要人工审查
    
    Decision   int64 `json:"decision,omitempty"`
    

    // rule id
    
    Id   string `json:"id,omitempty"`
    

    // rule name
    
    Name   string `json:"name,omitempty"`
    

    // 分数
    
    Score   int64 `json:"score,omitempty"`
    

}
*/

// JaqAccountRiskDetailItem 
type JaqAccountRiskDetailItem struct {

    // 决定，0：可以接受；1：应该拒绝；2：需要人工审查
    Decision   int64 `json:"decision,omitempty"`

    // rule id
    Id   string `json:"id,omitempty"`

    // rule name
    Name   string `json:"name,omitempty"`

    // 分数
    Score   int64 `json:"score,omitempty"`

}
