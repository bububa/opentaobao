package user

// RecommendInfo 
/* model for simplify = false
type RecommendInfo struct {

    // 推荐具体数据list
    
    RecommendGuides  struct {
        RecommendGuide  []RecommendGuide `json:"recommend_guide,omitempty"`
    } `json:"recommend_guides,omitempty"`
    

}
*/

// RecommendInfo 
type RecommendInfo struct {

    // 推荐具体数据list
    RecommendGuides   []RecommendGuide `json:"recommend_guides,omitempty"`

}
