package simba

// CampaignGenderAgeFilterVo 结构体
type CampaignGenderAgeFilterVo struct {
	// 年龄段,0:,1,2,3,4,5,6
	AgeList []string `json:"age_list,omitempty" xml:"age_list>string,omitempty"`
	// 性别,man:男性人群,woman:女性人群
	Gender string `json:"gender,omitempty" xml:"gender,omitempty"`
}
