package icbu

// ProductScoreInfoResult 
type ProductScoreInfoResult struct {
    // 精品标，，返回字段中 boutique_tag 含义： 1 精品 2 普通品 3 低质品 4 实力优品
    BoutiqueTag   int64 `json:"boutique_tag,omitempty" xml:"boutique_tag,omitempty"`
    // 质量分
    FinalScore   string `json:"final_score,omitempty" xml:"final_score,omitempty"`
}
