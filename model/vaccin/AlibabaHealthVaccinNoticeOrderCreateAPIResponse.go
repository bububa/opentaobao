package vaccin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHealthVaccinNoticeOrderCreateAPIResponse 支付宝医疗健康疫苗预约创建 API返回值
// alibaba.health.vaccin.notice.order.create
//
// 支付宝医疗健康疫苗预约创建
type AlibabaHealthVaccinNoticeOrderCreateAPIResponse struct {
	model.CommonResponse
	AlibabaHealthVaccinNoticeOrderCreateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaHealthVaccinNoticeOrderCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHealthVaccinNoticeOrderCreateAPIResponseModel).Reset()
}

// AlibabaHealthVaccinNoticeOrderCreateAPIResponseModel is 支付宝医疗健康疫苗预约创建 成功返回结果
type AlibabaHealthVaccinNoticeOrderCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_health_vaccin_notice_order_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回描述
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// 结果集
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaHealthVaccinNoticeOrderCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MsgInfo = ""
	m.MsgCode = ""
	m.IsSuccess = false
	m.Model = false
}

var poolAlibabaHealthVaccinNoticeOrderCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHealthVaccinNoticeOrderCreateAPIResponse)
	},
}

// GetAlibabaHealthVaccinNoticeOrderCreateAPIResponse 从 sync.Pool 获取 AlibabaHealthVaccinNoticeOrderCreateAPIResponse
func GetAlibabaHealthVaccinNoticeOrderCreateAPIResponse() *AlibabaHealthVaccinNoticeOrderCreateAPIResponse {
	return poolAlibabaHealthVaccinNoticeOrderCreateAPIResponse.Get().(*AlibabaHealthVaccinNoticeOrderCreateAPIResponse)
}

// ReleaseAlibabaHealthVaccinNoticeOrderCreateAPIResponse 将 AlibabaHealthVaccinNoticeOrderCreateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHealthVaccinNoticeOrderCreateAPIResponse(v *AlibabaHealthVaccinNoticeOrderCreateAPIResponse) {
	v.Reset()
	poolAlibabaHealthVaccinNoticeOrderCreateAPIResponse.Put(v)
}
