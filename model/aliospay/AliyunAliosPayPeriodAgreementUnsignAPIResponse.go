package aliospay

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliyunAliosPayPeriodAgreementUnsignAPIResponse 周期扣款协议解约接口 API返回值
// aliyun.alios.pay.period.agreement.unsign
//
// 周期扣款协议解约接口
type AliyunAliosPayPeriodAgreementUnsignAPIResponse struct {
	model.CommonResponse
	AliyunAliosPayPeriodAgreementUnsignAPIResponseModel
}

// Reset 清空结构体
func (m *AliyunAliosPayPeriodAgreementUnsignAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliyunAliosPayPeriodAgreementUnsignAPIResponseModel).Reset()
}

// AliyunAliosPayPeriodAgreementUnsignAPIResponseModel is 周期扣款协议解约接口 成功返回结果
type AliyunAliosPayPeriodAgreementUnsignAPIResponseModel struct {
	XMLName xml.Name `xml:"aliyun_alios_pay_period_agreement_unsign_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应参数
	AliospayResponse *AliOSPayResponse `json:"aliospay_response,omitempty" xml:"aliospay_response,omitempty"`
}

// Reset 清空结构体
func (m *AliyunAliosPayPeriodAgreementUnsignAPIResponseModel) Reset() {
	m.RequestId = ""
	m.AliospayResponse = nil
}

var poolAliyunAliosPayPeriodAgreementUnsignAPIResponse = sync.Pool{
	New: func() any {
		return new(AliyunAliosPayPeriodAgreementUnsignAPIResponse)
	},
}

// GetAliyunAliosPayPeriodAgreementUnsignAPIResponse 从 sync.Pool 获取 AliyunAliosPayPeriodAgreementUnsignAPIResponse
func GetAliyunAliosPayPeriodAgreementUnsignAPIResponse() *AliyunAliosPayPeriodAgreementUnsignAPIResponse {
	return poolAliyunAliosPayPeriodAgreementUnsignAPIResponse.Get().(*AliyunAliosPayPeriodAgreementUnsignAPIResponse)
}

// ReleaseAliyunAliosPayPeriodAgreementUnsignAPIResponse 将 AliyunAliosPayPeriodAgreementUnsignAPIResponse 保存到 sync.Pool
func ReleaseAliyunAliosPayPeriodAgreementUnsignAPIResponse(v *AliyunAliosPayPeriodAgreementUnsignAPIResponse) {
	v.Reset()
	poolAliyunAliosPayPeriodAgreementUnsignAPIResponse.Put(v)
}
