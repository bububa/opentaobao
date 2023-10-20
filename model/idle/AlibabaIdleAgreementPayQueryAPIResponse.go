package idle

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleAgreementPayQueryAPIResponse 代扣详情查询 API返回值
// alibaba.idle.agreement.pay.query
//
// 查询代扣结果
type AlibabaIdleAgreementPayQueryAPIResponse struct {
	model.CommonResponse
	AlibabaIdleAgreementPayQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIdleAgreementPayQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIdleAgreementPayQueryAPIResponseModel).Reset()
}

// AlibabaIdleAgreementPayQueryAPIResponseModel is 代扣详情查询 成功返回结果
type AlibabaIdleAgreementPayQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_agreement_pay_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回
	Result *AlibabaIdleAgreementPayQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIdleAgreementPayQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaIdleAgreementPayQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIdleAgreementPayQueryAPIResponse)
	},
}

// GetAlibabaIdleAgreementPayQueryAPIResponse 从 sync.Pool 获取 AlibabaIdleAgreementPayQueryAPIResponse
func GetAlibabaIdleAgreementPayQueryAPIResponse() *AlibabaIdleAgreementPayQueryAPIResponse {
	return poolAlibabaIdleAgreementPayQueryAPIResponse.Get().(*AlibabaIdleAgreementPayQueryAPIResponse)
}

// ReleaseAlibabaIdleAgreementPayQueryAPIResponse 将 AlibabaIdleAgreementPayQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIdleAgreementPayQueryAPIResponse(v *AlibabaIdleAgreementPayQueryAPIResponse) {
	v.Reset()
	poolAlibabaIdleAgreementPayQueryAPIResponse.Put(v)
}
