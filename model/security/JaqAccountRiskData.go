package security

import (
	"sync"
)

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
	// 触发验证所需的信息
	CaptchaCheckData *CaptchaCheckData `json:"captcha_check_data,omitempty" xml:"captcha_check_data,omitempty"`
	// 最终决定。0：可以接受；1：应该拒绝；2：需要人工审核
	FinalDecision int64 `json:"final_decision,omitempty" xml:"final_decision,omitempty"`
	// 最终得分
	FinalScore int64 `json:"final_score,omitempty" xml:"final_score,omitempty"`
}

var poolJaqAccountRiskData = sync.Pool{
	New: func() any {
		return new(JaqAccountRiskData)
	},
}

// GetJaqAccountRiskData() 从对象池中获取JaqAccountRiskData
func GetJaqAccountRiskData() *JaqAccountRiskData {
	return poolJaqAccountRiskData.Get().(*JaqAccountRiskData)
}

// ReleaseJaqAccountRiskData 释放JaqAccountRiskData
func ReleaseJaqAccountRiskData(v *JaqAccountRiskData) {
	v.Detail = v.Detail[:0]
	v.DetailList = v.DetailList[:0]
	v.EventId = ""
	v.FinalDesc = ""
	v.UserId = ""
	v.Umid = ""
	v.CaptchaCheckData = nil
	v.FinalDecision = 0
	v.FinalScore = 0
	poolJaqAccountRiskData.Put(v)
}
