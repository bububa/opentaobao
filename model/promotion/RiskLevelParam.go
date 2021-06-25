package promotion

// RiskLevelParam 
type RiskLevelParam struct {

    // 风险等级(可选值:higher-risk,middle-risk,low-risk)
    Key   string `json:"key,omitempty"`

    // 风险等级比例左边值
    LeftRange   int64 `json:"left_range,omitempty"`

    // 风险等级比例右边值
    RightRange   int64 `json:"right_range,omitempty"`

}
