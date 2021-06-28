package wdk

// ReasonVo 
/* model for simplify = false
type ReasonVo struct {

    // 原因id
    
    ReasonId   int64 `json:"reason_id,omitempty"`
    

    // 原因描述
    
    ReasonText   string `json:"reason_text,omitempty"`
    

    // tip
    
    ReasonTip   string `json:"reason_tip,omitempty"`
    

    // 标签
    
    Tags  struct {
        String  []string `json:"string,omitempty"`
    } `json:"tags,omitempty"`
    

}
*/

// ReasonVo 
type ReasonVo struct {

    // 原因id
    ReasonId   int64 `json:"reason_id,omitempty"`

    // 原因描述
    ReasonText   string `json:"reason_text,omitempty"`

    // tip
    ReasonTip   string `json:"reason_tip,omitempty"`

    // 标签
    Tags   []string `json:"tags,omitempty"`

}
