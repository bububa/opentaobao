package user

// RecommendInfo 结构体
type RecommendInfo struct {
	// 推荐具体数据list
	RecommendGuides []RecommendGuide `json:"recommend_guides,omitempty" xml:"recommend_guides>recommend_guide,omitempty"`
}
