package idle

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIdleAgreementPayAPIResponse
闲鱼平台商户代扣 API返回值
alibaba.idle.agreement.pay

闲鱼平台代扣能力：用户和闲鱼签约代扣协议 服务商通过直付通产品挂载成为闲鱼二级商户 来完成用户和服务商的结算 */
type AlibabaIdleAgreementPayAPIResponse struct {
	model.CommonResponse
	AlibabaIdleAgreementPayAPIResponseModel
}

// AlibabaIdleAgreementPayAPIResponseModel is 闲鱼平台商户代扣 成功返回结果
type AlibabaIdleAgreementPayAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_agreement_pay_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaIdleAgreementPayResult `json:"result,omitempty" xml:"result,omitempty"`
}
