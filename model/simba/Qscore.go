package simba

// Qscore 结构体
type Qscore struct {
	// 词质量得分列表
	KeywordQscoreList []KeywordQscore `json:"keyword_qscore_list,omitempty" xml:"keyword_qscore_list>keyword_qscore,omitempty"`
	// 类目出价质量得分
	CatmatchQscore string `json:"catmatch_qscore,omitempty" xml:"catmatch_qscore,omitempty"`
}
