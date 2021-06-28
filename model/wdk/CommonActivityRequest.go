package wdk

// CommonActivityRequest 
/* model for simplify = false
type CommonActivityRequest struct {

    // 五道口活动id
    
    ActivityId   int64 `json:"activity_id,omitempty"`
    

    // 商家活动id
    
    OutActId   string `json:"out_act_id,omitempty"`
    

}
*/

// CommonActivityRequest 
type CommonActivityRequest struct {

    // 五道口活动id
    ActivityId   int64 `json:"activity_id,omitempty"`

    // 商家活动id
    OutActId   string `json:"out_act_id,omitempty"`

}
