package vaccin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHealthVaccinNoticeUserCreateAPIResponse 支付宝医疗健康疫苗用户创建 API返回值
// alibaba.health.vaccin.notice.user.create
//
// 支付宝医疗健康疫苗用户创建
type AlibabaHealthVaccinNoticeUserCreateAPIResponse struct {
	model.CommonResponse
	AlibabaHealthVaccinNoticeUserCreateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaHealthVaccinNoticeUserCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHealthVaccinNoticeUserCreateAPIResponseModel).Reset()
}

// AlibabaHealthVaccinNoticeUserCreateAPIResponseModel is 支付宝医疗健康疫苗用户创建 成功返回结果
type AlibabaHealthVaccinNoticeUserCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_health_vaccin_notice_user_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回描述
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 结果集
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaHealthVaccinNoticeUserCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MsgInfo = ""
	m.MsgCode = ""
	m.Model = false
	m.IsSuccess = false
}

var poolAlibabaHealthVaccinNoticeUserCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHealthVaccinNoticeUserCreateAPIResponse)
	},
}

// GetAlibabaHealthVaccinNoticeUserCreateAPIResponse 从 sync.Pool 获取 AlibabaHealthVaccinNoticeUserCreateAPIResponse
func GetAlibabaHealthVaccinNoticeUserCreateAPIResponse() *AlibabaHealthVaccinNoticeUserCreateAPIResponse {
	return poolAlibabaHealthVaccinNoticeUserCreateAPIResponse.Get().(*AlibabaHealthVaccinNoticeUserCreateAPIResponse)
}

// ReleaseAlibabaHealthVaccinNoticeUserCreateAPIResponse 将 AlibabaHealthVaccinNoticeUserCreateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHealthVaccinNoticeUserCreateAPIResponse(v *AlibabaHealthVaccinNoticeUserCreateAPIResponse) {
	v.Reset()
	poolAlibabaHealthVaccinNoticeUserCreateAPIResponse.Put(v)
}
