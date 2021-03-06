package simba

// QScoreSplitDto 结构体
type QScoreSplitDto struct {
	// 关键词新质量得分列表，包含PC和移动的质量分
	WordScoreList []Wordscorelist `json:"word_score_list,omitempty" xml:"word_score_list>wordscorelist,omitempty"`
	// 类目质量得分
	CatMatchScore string `json:"cat_match_score,omitempty" xml:"cat_match_score,omitempty"`
	// 推广组id
	AdgroupId int64 `json:"adgroup_id,omitempty" xml:"adgroup_id,omitempty"`
}
