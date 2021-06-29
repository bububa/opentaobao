package scbp

// CrowdInfo 
type CrowdInfo struct {
    // 人群中文名称
    CrowdName   string `json:"crowd_name,omitempty" xml:"crowd_name,omitempty"`
    // 人群id
    CrowdId   string `json:"crowd_id,omitempty" xml:"crowd_id,omitempty"`
    // 人群预估数量
    EstimateCountList   []int64 `json:"estimate_count_list,omitempty" xml:"estimate_count_list>int64,omitempty"`
}
