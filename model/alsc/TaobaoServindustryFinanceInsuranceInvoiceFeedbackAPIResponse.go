package alsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoServindustryFinanceInsuranceInvoiceFeedbackAPIResponse 保险-开票结果反馈 API返回值
// taobao.servindustry.finance.insurance.invoice.feedback
//
// 保险-开票结果反馈
type TaobaoServindustryFinanceInsuranceInvoiceFeedbackAPIResponse struct {
	model.CommonResponse
	TaobaoServindustryFinanceInsuranceInvoiceFeedbackAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoServindustryFinanceInsuranceInvoiceFeedbackAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoServindustryFinanceInsuranceInvoiceFeedbackAPIResponseModel).Reset()
}

// TaobaoServindustryFinanceInsuranceInvoiceFeedbackAPIResponseModel is 保险-开票结果反馈 成功返回结果
type TaobaoServindustryFinanceInsuranceInvoiceFeedbackAPIResponseModel struct {
	XMLName xml.Name `xml:"servindustry_finance_insurance_invoice_feedback_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *TaobaoServindustryFinanceInsuranceInvoiceFeedbackResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoServindustryFinanceInsuranceInvoiceFeedbackAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoServindustryFinanceInsuranceInvoiceFeedbackAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoServindustryFinanceInsuranceInvoiceFeedbackAPIResponse)
	},
}

// GetTaobaoServindustryFinanceInsuranceInvoiceFeedbackAPIResponse 从 sync.Pool 获取 TaobaoServindustryFinanceInsuranceInvoiceFeedbackAPIResponse
func GetTaobaoServindustryFinanceInsuranceInvoiceFeedbackAPIResponse() *TaobaoServindustryFinanceInsuranceInvoiceFeedbackAPIResponse {
	return poolTaobaoServindustryFinanceInsuranceInvoiceFeedbackAPIResponse.Get().(*TaobaoServindustryFinanceInsuranceInvoiceFeedbackAPIResponse)
}

// ReleaseTaobaoServindustryFinanceInsuranceInvoiceFeedbackAPIResponse 将 TaobaoServindustryFinanceInsuranceInvoiceFeedbackAPIResponse 保存到 sync.Pool
func ReleaseTaobaoServindustryFinanceInsuranceInvoiceFeedbackAPIResponse(v *TaobaoServindustryFinanceInsuranceInvoiceFeedbackAPIResponse) {
	v.Reset()
	poolTaobaoServindustryFinanceInsuranceInvoiceFeedbackAPIResponse.Put(v)
}
