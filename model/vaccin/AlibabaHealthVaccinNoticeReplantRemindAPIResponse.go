package vaccin

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *AlibabaHealthVaccinNoticeReplantRemindAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHealthVaccinNoticeReplantRemindAPIResponseModel).Reset()
}

// AlibabaHealthVaccinNoticeReplantRemindAPIResponseModel is 支付宝疫苗补种提醒信息 成功返回结果
type AlibabaHealthVaccinNoticeReplantRemindAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_health_vaccin_notice_replant_remind_response"`
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
func (m *AlibabaHealthVaccinNoticeReplantRemindAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MsgCode = ""
	m.MsgInfo = ""
	m.IsSuccess = false
	m.Model = false
}

var poolAlibabaHealthVaccinNoticeReplantRemindAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHealthVaccinNoticeReplantRemindAPIResponse)
	},
}

// GetAlibabaHealthVaccinNoticeReplantRemindAPIResponse 从 sync.Pool 获取 AlibabaHealthVaccinNoticeReplantRemindAPIResponse
func GetAlibabaHealthVaccinNoticeReplantRemindAPIResponse() *AlibabaHealthVaccinNoticeReplantRemindAPIResponse {
	return poolAlibabaHealthVaccinNoticeReplantRemindAPIResponse.Get().(*AlibabaHealthVaccinNoticeReplantRemindAPIResponse)
}

// ReleaseAlibabaHealthVaccinNoticeReplantRemindAPIResponse 将 AlibabaHealthVaccinNoticeReplantRemindAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHealthVaccinNoticeReplantRemindAPIResponse(v *AlibabaHealthVaccinNoticeReplantRemindAPIResponse) {
	v.Reset()
	poolAlibabaHealthVaccinNoticeReplantRemindAPIResponse.Put(v)
}
