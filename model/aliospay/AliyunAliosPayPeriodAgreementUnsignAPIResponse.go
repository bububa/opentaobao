package aliospay

import (
	"encoding/xml"

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

// AliyunAliosPayPeriodAgreementUnsignAPIResponseModel is 周期扣款协议解约接口 成功返回结果
type AliyunAliosPayPeriodAgreementUnsignAPIResponseModel struct {
	XMLName xml.Name `xml:"aliyun_alios_pay_period_agreement_unsign_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应参数
	AliospayResponse *AliOspayResponse `json:"aliospay_response,omitempty" xml:"aliospay_response,omitempty"`
}
