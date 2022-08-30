package trade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoServindustryFinanceInsuranceAgreementSignAPIResponse 保司合同签约后回调接口 API返回值
// taobao.servindustry.finance.insurance.agreement.sign
//
// 保司合同签约后回调接口
type TaobaoServindustryFinanceInsuranceAgreementSignAPIResponse struct {
	model.CommonResponse
	TaobaoServindustryFinanceInsuranceAgreementSignAPIResponseModel
}

// TaobaoServindustryFinanceInsuranceAgreementSignAPIResponseModel is 保司合同签约后回调接口 成功返回结果
type TaobaoServindustryFinanceInsuranceAgreementSignAPIResponseModel struct {
	XMLName xml.Name `xml:"servindustry_finance_insurance_agreement_sign_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回参数
	Result *RetryResult `json:"result,omitempty" xml:"result,omitempty"`
}
