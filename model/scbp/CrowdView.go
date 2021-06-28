package scbp

// CrowdView 
/* model for simplify = false
type CrowdView struct {

    // 人群list
    
    CrowdList  struct {
        CrowdInfo  []CrowdInfo `json:"crowd_info,omitempty"`
    } `json:"crowd_list,omitempty"`
    

    // 人群类型(店铺老客、优选人群)
    
    CrowdType   string `json:"crowd_type,omitempty"`
    

}
*/

// CrowdView 
type CrowdView struct {

    // 人群list
    CrowdList   []CrowdInfo `json:"crowd_list,omitempty"`

    // 人群类型(店铺老客、优选人群)
    CrowdType   string `json:"crowd_type,omitempty"`

}
