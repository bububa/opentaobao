package user

// RecommendGuide 
/* model for simplify = false
type RecommendGuide struct {

    // 推荐类型
    
    GuideType   string `json:"guide_type,omitempty"`
    

    // 推荐引导语
    
    GuideUtterance   string `json:"guide_utterance,omitempty"`
    

}
*/

// RecommendGuide 
type RecommendGuide struct {

    // 推荐类型
    GuideType   string `json:"guide_type,omitempty"`

    // 推荐引导语
    GuideUtterance   string `json:"guide_utterance,omitempty"`

}
