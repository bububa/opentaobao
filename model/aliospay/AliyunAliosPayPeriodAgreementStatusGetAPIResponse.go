package aliospay

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliyunAliosPayPeriodAgreementStatusGetAPIResponse 查询周期扣款签约状态 API返回值
// aliyun.alios.pay.period.agreement.status.get
//
// 查询周期扣款签约状态
type AliyunAliosPayPeriodAgreementStatusGetAPIResponse struct {
	model.CommonResponse
	AliyunAliosPayPeriodAgreementStatusGetAPIResponseModel
}

// Reset 清空结构体
func (m *AliyunAliosPayPeriodAgreementStatusGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliyunAliosPayPeriodAgreementStatusGetAPIResponseModel).Reset()
}

// AliyunAliosPayPeriodAgreementStatusGetAPIResponseModel is 查询周期扣款签约状态 成功返回结果
type AliyunAliosPayPeriodAgreementStatusGetAPIResponseModel struct {
	XMLName xml.Name `xml:"aliyun_alios_pay_period_agreement_status_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应参数
	AliospayResponse *AliOSPayResponse `json:"aliospay_response,omitempty" xml:"aliospay_response,omitempty"`
}

// Reset 清空结构体
func (m *AliyunAliosPayPeriodAgreementStatusGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.AliospayResponse = nil
}

var poolAliyunAliosPayPeriodAgreementStatusGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AliyunAliosPayPeriodAgreementStatusGetAPIResponse)
	},
}

// GetAliyunAliosPayPeriodAgreementStatusGetAPIResponse 从 sync.Pool 获取 AliyunAliosPayPeriodAgreementStatusGetAPIResponse
func GetAliyunAliosPayPeriodAgreementStatusGetAPIResponse() *AliyunAliosPayPeriodAgreementStatusGetAPIResponse {
	return poolAliyunAliosPayPeriodAgreementStatusGetAPIResponse.Get().(*AliyunAliosPayPeriodAgreementStatusGetAPIResponse)
}

// ReleaseAliyunAliosPayPeriodAgreementStatusGetAPIResponse 将 AliyunAliosPayPeriodAgreementStatusGetAPIResponse 保存到 sync.Pool
func ReleaseAliyunAliosPayPeriodAgreementStatusGetAPIResponse(v *AliyunAliosPayPeriodAgreementStatusGetAPIResponse) {
	v.Reset()
	poolAliyunAliosPayPeriodAgreementStatusGetAPIResponse.Put(v)
}
