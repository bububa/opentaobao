package einvoice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
更新用户发薪资产 APIResponse
alibaba.einvoice.tax.opt.salaryaccount.update

更新用户的发薪账号
*/
type AlibabaEinvoiceTaxOptSalaryaccountUpdateAPIResponse struct {
    model.CommonResponse
    AlibabaEinvoiceTaxOptSalaryaccountUpdateResponse
}

type AlibabaEinvoiceTaxOptSalaryaccountUpdateResponse struct {
    XMLName xml.Name `xml:"alibaba_einvoice_tax_opt_salaryaccount_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 出参
    
    Result   *TaxOptimizationEmployeeAssetUpdateResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
