package einvoice

import (
	"encoding/xml"

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

// AlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentAPIResponseModel is 受理发薪 成功返回结果
type AlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_einvoice_tax_opt_salaryrequest_acceptpayment_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 服务出参
	Result *TaxOptimizationSalaryPaymentAccessResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
