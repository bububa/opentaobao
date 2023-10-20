package vaccin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHealthVaccinNoticeTimebucketRemindAPIResponse 疫苗预约时间段提醒 API返回值
// alibaba.health.vaccin.notice.timebucket.remind
//
// 疫苗预约时间段提醒
type AlibabaHealthVaccinNoticeTimebucketRemindAPIResponse struct {
	model.CommonResponse
	AlibabaHealthVaccinNoticeTimebucketRemindAPIResponseModel
}

// AlibabaHealthVaccinNoticeTimebucketRemindAPIResponseModel is 疫苗预约时间段提醒 成功返回结果
type AlibabaHealthVaccinNoticeTimebucketRemindAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_health_vaccin_notice_timebucket_remind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误code
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 错误描述
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 执行是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// 结果
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
}
