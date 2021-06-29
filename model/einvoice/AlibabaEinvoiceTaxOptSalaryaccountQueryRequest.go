package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询用户的发薪账号 API请求
alibaba.einvoice.tax.opt.salaryaccount.query

查询用户的发薪账号状态
*/
type AlibabaEinvoiceTaxOptSalaryaccountQueryRequest struct {
    model.Params
    // 入参
    _paramTaxOptimizationQueryPaySalaryAccountDTO   *TaxOptimizationQueryPaySalaryAccountDto
}

// 初始化AlibabaEinvoiceTaxOptSalaryaccountQueryRequest对象
func NewAlibabaEinvoiceTaxOptSalaryaccountQueryRequest() *AlibabaEinvoiceTaxOptSalaryaccountQueryRequest{
    return &AlibabaEinvoiceTaxOptSalaryaccountQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceTaxOptSalaryaccountQueryRequest) GetApiMethodName() string {
    return "alibaba.einvoice.tax.opt.salaryaccount.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceTaxOptSalaryaccountQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamTaxOptimizationQueryPaySalaryAccountDTO Setter
// 入参
func (r *AlibabaEinvoiceTaxOptSalaryaccountQueryRequest) SetParamTaxOptimizationQueryPaySalaryAccountDTO(_paramTaxOptimizationQueryPaySalaryAccountDTO *TaxOptimizationQueryPaySalaryAccountDto) error {
    r._paramTaxOptimizationQueryPaySalaryAccountDTO = _paramTaxOptimizationQueryPaySalaryAccountDTO
    r.Set("param_tax_optimization_query_pay_salary_account_d_t_o", _paramTaxOptimizationQueryPaySalaryAccountDTO)
    return nil
}

// ParamTaxOptimizationQueryPaySalaryAccountDTO Getter
func (r AlibabaEinvoiceTaxOptSalaryaccountQueryRequest) GetParamTaxOptimizationQueryPaySalaryAccountDTO() *TaxOptimizationQueryPaySalaryAccountDto {
    return r._paramTaxOptimizationQueryPaySalaryAccountDTO
}
