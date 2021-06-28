package legalcase

// RefStandpointModel 
/* model for simplify = false
type RefStandpointModel struct {

    // 案件id
    
    SuitId   int64 `json:"suit_id,omitempty"`
    

    // 委托id
    
    EntrustId   int64 `json:"entrust_id,omitempty"`
    

    // 口径id
    
    StandpointId   int64 `json:"standpoint_id,omitempty"`
    

    // 口径描述
    
    StandpointDesc   string `json:"standpoint_desc,omitempty"`
    

    // 答辩口径
    
    DefenseCaliber   string `json:"defense_caliber,omitempty"`
    

    // 口径版本
    
    Version   int64 `json:"version,omitempty"`
    

}
*/

// RefStandpointModel 
type RefStandpointModel struct {

    // 案件id
    SuitId   int64 `json:"suit_id,omitempty"`

    // 委托id
    EntrustId   int64 `json:"entrust_id,omitempty"`

    // 口径id
    StandpointId   int64 `json:"standpoint_id,omitempty"`

    // 口径描述
    StandpointDesc   string `json:"standpoint_desc,omitempty"`

    // 答辩口径
    DefenseCaliber   string `json:"defense_caliber,omitempty"`

    // 口径版本
    Version   int64 `json:"version,omitempty"`

}
