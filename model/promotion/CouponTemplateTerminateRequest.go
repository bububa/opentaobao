package promotion

// CouponTemplateTerminateRequest 
/* model for simplify = false
type CouponTemplateTerminateRequest struct {

    // ump模板ID
    
    SourceId   int64 `json:"source_id,omitempty"`
    

    // 用户信息
    
    UserInfo  *struct {
        UserInfo  *UserInfo `json:"user_info,omitempty"`
    } `json:"user_info,omitempty"`
    

}
*/

// CouponTemplateTerminateRequest 
type CouponTemplateTerminateRequest struct {

    // ump模板ID
    SourceId   int64 `json:"source_id,omitempty"`

    // 用户信息
    UserInfo   *UserInfo `json:"user_info,omitempty"`

}
