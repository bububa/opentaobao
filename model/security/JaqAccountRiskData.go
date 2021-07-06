package security

// JaqAccountRiskData 结构体
type JaqAccountRiskData struct {
	// 详情列表
	Detail []JaqAccountRiskDetailItem `json:"detail,omitempty" xml:"detail>jaq_account_risk_detail_item,omitempty"`
	// 详情列表
	DetailList []JaqAccountRiskDetailItem `json:"detail_list,omitempty" xml:"detail_list>jaq_account_risk_detail_item,omitempty"`
	// 事件id
	EventId string `json:"event_id,omitempty" xml:"event_id,omitempty"`
	// 综合风险描述
	FinalDesc string `json:"final_desc,omitempty" xml:"final_desc,omitempty"`
	// 注册用户id
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// umid
	Umid string `json:"umid,omitempty" xml:"umid,omitempty"`
	// 最终决定。0：可以接受；1：应该拒绝；2：需要人工审核
	FinalDecision int64 `json:"final_decision,omitempty" xml:"final_decision,omitempty"`
	// 最终得分
	FinalScore int64 `json:"final_score,omitempty" xml:"final_score,omitempty"`
	// 触发验证所需的信息
	CaptchaCheckData *CaptchaCheckData `json:"captcha_check_data,omitempty" xml:"captcha_check_data,omitempty"`
}
