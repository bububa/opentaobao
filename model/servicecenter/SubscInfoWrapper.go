package servicecenter

// SubscInfoWrapper 
/* model for simplify = false
type SubscInfoWrapper struct {

    // 总量
    
    TotalCount   int64 `json:"total_count,omitempty"`
    

    // 需求订购信息
    
    SubscInfoList  struct {
        SubscInfo  []SubscInfo `json:"subsc_info,omitempty"`
    } `json:"subsc_info_list,omitempty"`
    

}
*/

// SubscInfoWrapper 
type SubscInfoWrapper struct {

    // 总量
    TotalCount   int64 `json:"total_count,omitempty"`

    // 需求订购信息
    SubscInfoList   []SubscInfo `json:"subsc_info_list,omitempty"`

}
