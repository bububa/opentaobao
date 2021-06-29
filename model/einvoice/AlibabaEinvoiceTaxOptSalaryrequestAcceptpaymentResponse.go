package einvoice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
受理发薪 APIResponse
alibaba.einvoice.tax.opt.salaryrequest.acceptpayment

发薪受理接口
*/
type AlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentAPIResponse struct {
    model.CommonResponse
    AlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentResponse
}

type AlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentResponse struct {
    XMLName xml.Name `xml:"alibaba_einvoice_tax_opt_salaryrequest_acceptpayment_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 服务出参
    
    Result   *TaxOptimizationSalaryPaymentAccessResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
