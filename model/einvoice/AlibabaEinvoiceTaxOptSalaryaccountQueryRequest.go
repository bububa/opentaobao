package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询用户的发薪账号 APIRequest
alibaba.einvoice.tax.opt.salaryaccount.query

查询用户的发薪账号状态
*/
type AlibabaEinvoiceTaxOptSalaryaccountQueryRequest struct {
    model.Params

    // 入参
    paramTaxOptimizationQueryPaySalaryAccountDTO   *TaxOptimizationQueryPaySalaryAccountDto 

}

func NewAlibabaEinvoiceTaxOptSalaryaccountQueryRequest() *AlibabaEinvoiceTaxOptSalaryaccountQueryRequest{
    return &AlibabaEinvoiceTaxOptSalaryaccountQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaEinvoiceTaxOptSalaryaccountQueryRequest) GetApiMethodName() string {
    return "alibaba.einvoice.tax.opt.salaryaccount.query"
}

func (r AlibabaEinvoiceTaxOptSalaryaccountQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaEinvoiceTaxOptSalaryaccountQueryRequest) SetParamTaxOptimizationQueryPaySalaryAccountDTO(paramTaxOptimizationQueryPaySalaryAccountDTO *TaxOptimizationQueryPaySalaryAccountDto) error {
    r.paramTaxOptimizationQueryPaySalaryAccountDTO = paramTaxOptimizationQueryPaySalaryAccountDTO
    r.Set("param_tax_optimization_query_pay_salary_account_d_t_o", paramTaxOptimizationQueryPaySalaryAccountDTO)
    return nil
}

func (r AlibabaEinvoiceTaxOptSalaryaccountQueryRequest) GetParamTaxOptimizationQueryPaySalaryAccountDTO() *TaxOptimizationQueryPaySalaryAccountDto {
    return r.paramTaxOptimizationQueryPaySalaryAccountDTO
}

