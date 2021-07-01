package dmp

// DmpCrowdDto 
type DmpCrowdDto struct {
    // 过期时间
    ValidDate   string `json:"valid_date,omitempty" xml:"valid_date,omitempty"`
    // 用户ID
    UserId   int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
    // 人群名称
    CrowdName   string `json:"crowd_name,omitempty" xml:"crowd_name,omitempty"`
    // 选项
    Selects   []DmpSelectTagOptionDto `json:"selects,omitempty" xml:"selects>dmp_select_tag_option_dto,omitempty"`
    // 覆盖人数
    Coverage   int64 `json:"coverage,omitempty" xml:"coverage,omitempty"`
    // 人群ID
    CrowdId   int64 `json:"crowd_id,omitempty" xml:"crowd_id,omitempty"`
    // 放大倍数
    LookalikeMultiple   int64 `json:"lookalike_multiple,omitempty" xml:"lookalike_multiple,omitempty"`
}
