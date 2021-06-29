package einvoice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询用户的发薪账号 APIResponse
alibaba.einvoice.tax.opt.salaryaccount.query

查询用户的发薪账号状态
*/
type AlibabaEinvoiceTaxOptSalaryaccountQueryAPIResponse struct {
    model.CommonResponse
    AlibabaEinvoiceTaxOptSalaryaccountQueryResponse
}

type AlibabaEinvoiceTaxOptSalaryaccountQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_einvoice_tax_opt_salaryaccount_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 出参
    
    Result   *TaxOptimizationQueryAlipayAccountResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
