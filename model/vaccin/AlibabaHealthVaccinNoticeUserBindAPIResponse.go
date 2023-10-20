package vaccin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHealthVaccinNoticeUserBindAPIResponse 支付宝疫苗绑定接种人 API返回值
// alibaba.health.vaccin.notice.user.bind
//
// 支付宝疫苗绑定接种人
type AlibabaHealthVaccinNoticeUserBindAPIResponse struct {
	model.CommonResponse
	AlibabaHealthVaccinNoticeUserBindAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaHealthVaccinNoticeUserBindAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHealthVaccinNoticeUserBindAPIResponseModel).Reset()
}

// AlibabaHealthVaccinNoticeUserBindAPIResponseModel is 支付宝疫苗绑定接种人 成功返回结果
type AlibabaHealthVaccinNoticeUserBindAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_health_vaccin_notice_user_bind_response"`
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
func (m *AlibabaHealthVaccinNoticeUserBindAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MsgCode = ""
	m.MsgInfo = ""
	m.IsSuccess = false
	m.Model = false
}

var poolAlibabaHealthVaccinNoticeUserBindAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHealthVaccinNoticeUserBindAPIResponse)
	},
}

// GetAlibabaHealthVaccinNoticeUserBindAPIResponse 从 sync.Pool 获取 AlibabaHealthVaccinNoticeUserBindAPIResponse
func GetAlibabaHealthVaccinNoticeUserBindAPIResponse() *AlibabaHealthVaccinNoticeUserBindAPIResponse {
	return poolAlibabaHealthVaccinNoticeUserBindAPIResponse.Get().(*AlibabaHealthVaccinNoticeUserBindAPIResponse)
}

// ReleaseAlibabaHealthVaccinNoticeUserBindAPIResponse 将 AlibabaHealthVaccinNoticeUserBindAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHealthVaccinNoticeUserBindAPIResponse(v *AlibabaHealthVaccinNoticeUserBindAPIResponse) {
	v.Reset()
	poolAlibabaHealthVaccinNoticeUserBindAPIResponse.Put(v)
}
