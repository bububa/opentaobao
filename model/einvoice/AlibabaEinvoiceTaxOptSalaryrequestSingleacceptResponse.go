package einvoice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
单明细发薪受理 APIResponse
alibaba.einvoice.tax.opt.salaryrequest.singleaccept

单明细发薪受理
*/
type AlibabaEinvoiceTaxOptSalaryrequestSingleacceptAPIResponse struct {
    model.CommonResponse
    AlibabaEinvoiceTaxOptSalaryrequestSingleacceptResponse
}

type AlibabaEinvoiceTaxOptSalaryrequestSingleacceptResponse struct {
    XMLName xml.Name `xml:"alibaba_einvoice_tax_opt_salaryrequest_singleaccept_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果类
    
    Result   *TaxOptimizationSingleDetailPaymentAccessResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
