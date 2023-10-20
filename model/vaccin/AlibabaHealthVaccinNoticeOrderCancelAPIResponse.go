package vaccin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHealthVaccinNoticeOrderCancelAPIResponse 福州疫苗取消预约 API返回值
// alibaba.health.vaccin.notice.order.cancel
//
// 福州疫苗用户取消预约接口
type AlibabaHealthVaccinNoticeOrderCancelAPIResponse struct {
	model.CommonResponse
	AlibabaHealthVaccinNoticeOrderCancelAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaHealthVaccinNoticeOrderCancelAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHealthVaccinNoticeOrderCancelAPIResponseModel).Reset()
}

// AlibabaHealthVaccinNoticeOrderCancelAPIResponseModel is 福州疫苗取消预约 成功返回结果
type AlibabaHealthVaccinNoticeOrderCancelAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_health_vaccin_notice_order_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回码
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回描述
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 结果集
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaHealthVaccinNoticeOrderCancelAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MsgInfo = ""
	m.MsgCode = ""
	m.Model = false
	m.IsSuccess = false
}

var poolAlibabaHealthVaccinNoticeOrderCancelAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHealthVaccinNoticeOrderCancelAPIResponse)
	},
}

// GetAlibabaHealthVaccinNoticeOrderCancelAPIResponse 从 sync.Pool 获取 AlibabaHealthVaccinNoticeOrderCancelAPIResponse
func GetAlibabaHealthVaccinNoticeOrderCancelAPIResponse() *AlibabaHealthVaccinNoticeOrderCancelAPIResponse {
	return poolAlibabaHealthVaccinNoticeOrderCancelAPIResponse.Get().(*AlibabaHealthVaccinNoticeOrderCancelAPIResponse)
}

// ReleaseAlibabaHealthVaccinNoticeOrderCancelAPIResponse 将 AlibabaHealthVaccinNoticeOrderCancelAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHealthVaccinNoticeOrderCancelAPIResponse(v *AlibabaHealthVaccinNoticeOrderCancelAPIResponse) {
	v.Reset()
	poolAlibabaHealthVaccinNoticeOrderCancelAPIResponse.Put(v)
}
