package einvoice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询用户的发薪账号 API返回值 
alibaba.einvoice.tax.opt.salaryaccount.query

查询用户的发薪账号状态
*/
type AlibabaEinvoiceTaxOptSalaryaccountQueryAPIResponse struct {
    model.CommonResponse
    AlibabaEinvoiceTaxOptSalaryaccountQueryResponse
}

// 查询用户的发薪账号 成功返回结果
type AlibabaEinvoiceTaxOptSalaryaccountQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_einvoice_tax_opt_salaryaccount_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 出参
    Result   *TaxOptimizationQueryAlipayAccountResultDTO `json:"result,omitempty" xml:"result,omitempty"`
}
