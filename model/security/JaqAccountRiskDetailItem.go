package security

// JaqAccountRiskDetailItem 
type JaqAccountRiskDetailItem struct {
    // 决定，0：可以接受；1：应该拒绝；2：需要人工审查
    Decision   int64 `json:"decision,omitempty" xml:"decision,omitempty"`
    // rule id
    Id   string `json:"id,omitempty" xml:"id,omitempty"`
    // rule name
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    // 分数
    Score   int64 `json:"score,omitempty" xml:"score,omitempty"`
}
