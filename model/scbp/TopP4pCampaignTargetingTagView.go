package scbp

// TopP4pCampaignTargetingTagView 
type TopP4pCampaignTargetingTagView struct {

    // 国家溢价列表
    
    CountryTagList   []CountryTagView `json:"country_tag_list,omitempty" xml:"country_tag_list,omitempty"`
    

    // 人群溢价列表
    
    CrowdTagList   []CrowdTagView `json:"crowd_tag_list,omitempty" xml:"crowd_tag_list,omitempty"`
    

}
