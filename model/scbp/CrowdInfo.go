package scbp

// CrowdInfo 
/* model for simplify = false
type CrowdInfo struct {

    // 人群中文名称
    
    CrowdName   string `json:"crowd_name,omitempty"`
    

    // 人群id
    
    CrowdId   string `json:"crowd_id,omitempty"`
    

    // 人群预估数量
    
    EstimateCountList  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"estimate_count_list,omitempty"`
    

}
*/

// CrowdInfo 
type CrowdInfo struct {

    // 人群中文名称
    CrowdName   string `json:"crowd_name,omitempty"`

    // 人群id
    CrowdId   string `json:"crowd_id,omitempty"`

    // 人群预估数量
    EstimateCountList   []int64 `json:"estimate_count_list,omitempty"`

}
