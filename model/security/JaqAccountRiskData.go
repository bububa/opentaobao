package security

// JaqAccountRiskData 
/* model for simplify = false
type JaqAccountRiskData struct {

    // 详情列表
    
    Detail  struct {
        JaqAccountRiskDetailItem  []JaqAccountRiskDetailItem `json:"jaq_account_risk_detail_item,omitempty"`
    } `json:"detail,omitempty"`
    

    // 事件id
    
    EventId   string `json:"event_id,omitempty"`
    

    // 最终决定。0：可以接受；1：应该拒绝；2：需要人工审核
    
    FinalDecision   int64 `json:"final_decision,omitempty"`
    

    // 综合风险描述
    
    FinalDesc   string `json:"final_desc,omitempty"`
    

    // 最终得分
    
    FinalScore   int64 `json:"final_score,omitempty"`
    

    // 注册用户id
    
    UserId   string `json:"user_id,omitempty"`
    

    // 触发验证所需的信息
    
    CaptchaCheckData  *struct {
        CaptchaCheckData  *CaptchaCheckData `json:"captcha_check_data,omitempty"`
    } `json:"captcha_check_data,omitempty"`
    

    // umid
    
    Umid   string `json:"umid,omitempty"`
    

    // 详情列表
    
    DetailList  struct {
        JaqAccountRiskDetailItem  []JaqAccountRiskDetailItem `json:"jaq_account_risk_detail_item,omitempty"`
    } `json:"detail_list,omitempty"`
    

}
*/

// JaqAccountRiskData 
type JaqAccountRiskData struct {

    // 详情列表
    Detail   []JaqAccountRiskDetailItem `json:"detail,omitempty"`

    // 事件id
    EventId   string `json:"event_id,omitempty"`

    // 最终决定。0：可以接受；1：应该拒绝；2：需要人工审核
    FinalDecision   int64 `json:"final_decision,omitempty"`

    // 综合风险描述
    FinalDesc   string `json:"final_desc,omitempty"`

    // 最终得分
    FinalScore   int64 `json:"final_score,omitempty"`

    // 注册用户id
    UserId   string `json:"user_id,omitempty"`

    // 触发验证所需的信息
    CaptchaCheckData   *CaptchaCheckData `json:"captcha_check_data,omitempty"`

    // umid
    Umid   string `json:"umid,omitempty"`

    // 详情列表
    DetailList   []JaqAccountRiskDetailItem `json:"detail_list,omitempty"`

}
