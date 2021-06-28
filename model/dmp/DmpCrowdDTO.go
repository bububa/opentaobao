package dmp

// DmpCrowdDTO 
/* model for simplify = false
type DmpCrowdDTO struct {

    // 过期时间
    
    ValidDate   string `json:"valid_date,omitempty"`
    

    // 用户ID
    
    UserId   int64 `json:"user_id,omitempty"`
    

    // 人群名称
    
    CrowdName   string `json:"crowd_name,omitempty"`
    

    // 选项
    
    Selects  struct {
        DmpSelectTagOptionDTO  []DmpSelectTagOptionDTO `json:"dmp_select_tag_option_dto,omitempty"`
    } `json:"selects,omitempty"`
    

    // 覆盖人数
    
    Coverage   int64 `json:"coverage,omitempty"`
    

    // 人群ID
    
    CrowdId   int64 `json:"crowd_id,omitempty"`
    

    // 放大倍数
    
    LookalikeMultiple   int64 `json:"lookalike_multiple,omitempty"`
    

}
*/

// DmpCrowdDTO 
type DmpCrowdDTO struct {

    // 过期时间
    ValidDate   string `json:"valid_date,omitempty"`

    // 用户ID
    UserId   int64 `json:"user_id,omitempty"`

    // 人群名称
    CrowdName   string `json:"crowd_name,omitempty"`

    // 选项
    Selects   []DmpSelectTagOptionDTO `json:"selects,omitempty"`

    // 覆盖人数
    Coverage   int64 `json:"coverage,omitempty"`

    // 人群ID
    CrowdId   int64 `json:"crowd_id,omitempty"`

    // 放大倍数
    LookalikeMultiple   int64 `json:"lookalike_multiple,omitempty"`

}
