package scbp

// TopP4pCampaignTargetingTagView 
/* model for simplify = false
type TopP4pCampaignTargetingTagView struct {

    // 国家溢价列表
    
    CountryTagList  struct {
        CountryTagView  []CountryTagView `json:"country_tag_view,omitempty"`
    } `json:"country_tag_list,omitempty"`
    

    // 人群溢价列表
    
    CrowdTagList  struct {
        CrowdTagView  []CrowdTagView `json:"crowd_tag_view,omitempty"`
    } `json:"crowd_tag_list,omitempty"`
    

}
*/

// TopP4pCampaignTargetingTagView 
type TopP4pCampaignTargetingTagView struct {

    // 国家溢价列表
    CountryTagList   []CountryTagView `json:"country_tag_list,omitempty"`

    // 人群溢价列表
    CrowdTagList   []CrowdTagView `json:"crowd_tag_list,omitempty"`

}
