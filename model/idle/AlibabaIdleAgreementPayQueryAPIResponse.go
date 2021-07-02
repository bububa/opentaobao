package idle

import (
	"encoding/xml"

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

// AlibabaIdleAgreementPayQueryAPIResponseModel is 代扣详情查询 成功返回结果
type AlibabaIdleAgreementPayQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_agreement_pay_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回
	Result *AlibabaIdleAgreementPayQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
