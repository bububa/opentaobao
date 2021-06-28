package promotion

// CommonActivityParam 
/* model for simplify = false
type CommonActivityParam struct {

    // 五道口优惠券活动id
    
    ActivityId   int64 `json:"activity_id,omitempty"`
    

    // 商家优惠券活动id
    
    OutActId   string `json:"out_act_id,omitempty"`
    

}
*/

// CommonActivityParam 
type CommonActivityParam struct {

    // 五道口优惠券活动id
    ActivityId   int64 `json:"activity_id,omitempty"`

    // 商家优惠券活动id
    OutActId   string `json:"out_act_id,omitempty"`

}
