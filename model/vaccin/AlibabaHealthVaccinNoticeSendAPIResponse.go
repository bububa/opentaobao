package vaccin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHealthVaccinNoticeSendAPIResponse 发送消息提醒 API返回值
// alibaba.health.vaccin.notice.send
//
// ISV 通过免疫规划中心给用户发送短信或者支付宝 PUSH 提醒。
type AlibabaHealthVaccinNoticeSendAPIResponse struct {
	model.CommonResponse
	AlibabaHealthVaccinNoticeSendAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaHealthVaccinNoticeSendAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHealthVaccinNoticeSendAPIResponseModel).Reset()
}

// AlibabaHealthVaccinNoticeSendAPIResponseModel is 发送消息提醒 成功返回结果
type AlibabaHealthVaccinNoticeSendAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_health_vaccin_notice_send_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 找不到疫苗信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 200
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 是否成功执行
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// 有数据返回时的数据详情
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaHealthVaccinNoticeSendAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MsgInfo = ""
	m.MsgCode = ""
	m.IsSuccess = false
	m.Model = false
}

var poolAlibabaHealthVaccinNoticeSendAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHealthVaccinNoticeSendAPIResponse)
	},
}

// GetAlibabaHealthVaccinNoticeSendAPIResponse 从 sync.Pool 获取 AlibabaHealthVaccinNoticeSendAPIResponse
func GetAlibabaHealthVaccinNoticeSendAPIResponse() *AlibabaHealthVaccinNoticeSendAPIResponse {
	return poolAlibabaHealthVaccinNoticeSendAPIResponse.Get().(*AlibabaHealthVaccinNoticeSendAPIResponse)
}

// ReleaseAlibabaHealthVaccinNoticeSendAPIResponse 将 AlibabaHealthVaccinNoticeSendAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHealthVaccinNoticeSendAPIResponse(v *AlibabaHealthVaccinNoticeSendAPIResponse) {
	v.Reset()
	poolAlibabaHealthVaccinNoticeSendAPIResponse.Put(v)
}
