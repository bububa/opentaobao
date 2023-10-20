package einvoice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentAPIResponse 受理发薪 API返回值
// alibaba.einvoice.tax.opt.salaryrequest.acceptpayment
//
// 发薪受理接口
type AlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentAPIResponse struct {
	model.CommonResponse
	AlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentAPIResponseModel).Reset()
}

// AlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentAPIResponseModel is 受理发薪 成功返回结果
type AlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_einvoice_tax_opt_salaryrequest_acceptpayment_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 服务出参
	Result *TaxOptimizationSalaryPaymentAccessResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentAPIResponse)
	},
}

// GetAlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentAPIResponse 从 sync.Pool 获取 AlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentAPIResponse
func GetAlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentAPIResponse() *AlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentAPIResponse {
	return poolAlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentAPIResponse.Get().(*AlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentAPIResponse)
}

// ReleaseAlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentAPIResponse 将 AlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentAPIResponse 保存到 sync.Pool
func ReleaseAlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentAPIResponse(v *AlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentAPIResponse) {
	v.Reset()
	poolAlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentAPIResponse.Put(v)
}
