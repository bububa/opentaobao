package einvoice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询发薪结果 APIResponse
alibaba.einvoice.tax.opt.salaryresult.query

查询发薪结果
*/
type AlibabaEinvoiceTaxOptSalaryresultQueryAPIResponse struct {
    model.CommonResponse
    AlibabaEinvoiceTaxOptSalaryresultQueryResponse
}

type AlibabaEinvoiceTaxOptSalaryresultQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_einvoice_tax_opt_salaryresult_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 出参
    
    Result   *TaxOptimizationSalaryPayResultQueryResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
