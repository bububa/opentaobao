package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
单明细发薪受理 API请求
alibaba.einvoice.tax.opt.salaryrequest.singleaccept

单明细发薪受理
*/
type AlibabaEinvoiceTaxOptSalaryrequestSingleacceptRequest struct {
    model.Params
    // 入参
    _paramTaxOptimizationSingleDetailSalaryPaymentAccessDTO   *TaxOptimizationSingleDetailSalaryPaymentAccessDto
}

// 初始化AlibabaEinvoiceTaxOptSalaryrequestSingleacceptRequest对象
func NewAlibabaEinvoiceTaxOptSalaryrequestSingleacceptRequest() *AlibabaEinvoiceTaxOptSalaryrequestSingleacceptRequest{
    return &AlibabaEinvoiceTaxOptSalaryrequestSingleacceptRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceTaxOptSalaryrequestSingleacceptRequest) GetApiMethodName() string {
    return "alibaba.einvoice.tax.opt.salaryrequest.singleaccept"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceTaxOptSalaryrequestSingleacceptRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamTaxOptimizationSingleDetailSalaryPaymentAccessDTO Setter
// 入参
func (r *AlibabaEinvoiceTaxOptSalaryrequestSingleacceptRequest) SetParamTaxOptimizationSingleDetailSalaryPaymentAccessDTO(_paramTaxOptimizationSingleDetailSalaryPaymentAccessDTO *TaxOptimizationSingleDetailSalaryPaymentAccessDto) error {
    r._paramTaxOptimizationSingleDetailSalaryPaymentAccessDTO = _paramTaxOptimizationSingleDetailSalaryPaymentAccessDTO
    r.Set("param_tax_optimization_single_detail_salary_payment_access_d_t_o", _paramTaxOptimizationSingleDetailSalaryPaymentAccessDTO)
    return nil
}

// ParamTaxOptimizationSingleDetailSalaryPaymentAccessDTO Getter
func (r AlibabaEinvoiceTaxOptSalaryrequestSingleacceptRequest) GetParamTaxOptimizationSingleDetailSalaryPaymentAccessDTO() *TaxOptimizationSingleDetailSalaryPaymentAccessDto {
    return r._paramTaxOptimizationSingleDetailSalaryPaymentAccessDTO
}
