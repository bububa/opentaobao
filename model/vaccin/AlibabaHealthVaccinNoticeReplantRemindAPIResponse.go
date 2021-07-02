package vaccin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHealthVaccinNoticeReplantRemindAPIResponse 支付宝疫苗补种提醒信息 API返回值
// alibaba.health.vaccin.notice.replant.remind
//
// 支付宝疫苗补种提醒
type AlibabaHealthVaccinNoticeReplantRemindAPIResponse struct {
	model.CommonResponse
	AlibabaHealthVaccinNoticeReplantRemindAPIResponseModel
}

// AlibabaHealthVaccinNoticeReplantRemindAPIResponseModel is 支付宝疫苗补种提醒信息 成功返回结果
type AlibabaHealthVaccinNoticeReplantRemindAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_health_vaccin_notice_replant_remind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
	// 执行是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// 错误code
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 错误描述
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
}
