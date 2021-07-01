package westcrm

// ThirdTagsVo 结构体
type ThirdTagsVo struct {
	// 目标客群年龄段
	TagsAgeList []string `json:"tags_age_list,omitempty" xml:"tags_age_list>string,omitempty"`
	// 目前客群收入
	TagsMoneyList []string `json:"tags_money_list,omitempty" xml:"tags_money_list>string,omitempty"`
	// 目标客群等级
	TagsGraderList []string `json:"tags_grader_list,omitempty" xml:"tags_grader_list>string,omitempty"`
}
