package legalcase

// FeedbackRequestModel 
type FeedbackRequestModel struct {

    // 实际答辩口径
    DefenseCaliber   string `json:"defense_caliber,omitempty"`

    // 不采纳理由
    Reason   string `json:"reason,omitempty"`

    // 是否采纳标志 true 为采纳 , false 为不采纳
    AcceptFlag   bool `json:"accept_flag,omitempty"`

    // 案件id
    SuitId   int64 `json:"suit_id,omitempty"`

    // 委托id
    EntrustId   int64 `json:"entrust_id,omitempty"`

    // 观点id
    StandpointId   int64 `json:"standpoint_id,omitempty"`

    // 观点版本
    Version   int64 `json:"version,omitempty"`

    // 提交人
    SubmitPeople   string `json:"submit_people,omitempty"`

}
