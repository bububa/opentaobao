package scbp

// CrowdInfo 
type CrowdInfo struct {

    // 人群中文名称
    CrowdName   string `json:"crowd_name,omitempty"`

    // 人群id
    CrowdId   string `json:"crowd_id,omitempty"`

    // 人群预估数量
    EstimateCountList   []Number `json:"estimate_count_list,omitempty"`

}
