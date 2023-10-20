package vaccin

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *AlibabaHealthVaccinNoticeTimebucketRemindAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHealthVaccinNoticeTimebucketRemindAPIResponseModel).Reset()
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

// Reset 清空结构体
func (m *AlibabaHealthVaccinNoticeTimebucketRemindAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MsgCode = ""
	m.MsgInfo = ""
	m.IsSuccess = false
	m.Model = false
}

var poolAlibabaHealthVaccinNoticeTimebucketRemindAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHealthVaccinNoticeTimebucketRemindAPIResponse)
	},
}

// GetAlibabaHealthVaccinNoticeTimebucketRemindAPIResponse 从 sync.Pool 获取 AlibabaHealthVaccinNoticeTimebucketRemindAPIResponse
func GetAlibabaHealthVaccinNoticeTimebucketRemindAPIResponse() *AlibabaHealthVaccinNoticeTimebucketRemindAPIResponse {
	return poolAlibabaHealthVaccinNoticeTimebucketRemindAPIResponse.Get().(*AlibabaHealthVaccinNoticeTimebucketRemindAPIResponse)
}

// ReleaseAlibabaHealthVaccinNoticeTimebucketRemindAPIResponse 将 AlibabaHealthVaccinNoticeTimebucketRemindAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHealthVaccinNoticeTimebucketRemindAPIResponse(v *AlibabaHealthVaccinNoticeTimebucketRemindAPIResponse) {
	v.Reset()
	poolAlibabaHealthVaccinNoticeTimebucketRemindAPIResponse.Put(v)
}
