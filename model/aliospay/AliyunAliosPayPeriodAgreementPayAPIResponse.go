package aliospay

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AliyunaliospayperiodagreementpayAPIResponse 周期扣款支付接口 API返回值
// aliyun.alios.pay.period.agreement.pay
//
// 周期扣款支付接口，商户服务端通过该接口完成后续期数的支付
type AliyunaliospayperiodagreementpayAPIResponse struct {
	model.CommonResponse
	AliyunaliospayperiodagreementpayAPIResponseModel
}

// AliyunaliospayperiodagreementpayAPIResponseModel is 周期扣款支付接口 成功返回结果
type AliyunaliospayperiodagreementpayAPIResponseModel struct {
	XMLName xml.Name `xml:"aliyun_alios_pay_period_agreement_pay_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应参数
	AliospayResponse *AliOspayResponse `json:"aliospay_response,omitempty" xml:"aliospay_response,omitempty"`
}
