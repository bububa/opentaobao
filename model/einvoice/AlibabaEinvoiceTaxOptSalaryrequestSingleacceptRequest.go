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
type AlibabaEinvoiceTaxOptSalaryrequestSingleacceptAPIRequest struct {
    model.Params
    // 入参
    _paramTaxOptimizationSingleDetailSalaryPaymentAccessDTO   *TaxOptimizationSingleDetailSalaryPaymentAccessDTO
}

// 初始化AlibabaEinvoiceTaxOptSalaryrequestSingleacceptAPIRequest对象
func NewAlibabaEinvoiceTaxOptSalaryrequestSingleacceptRequest() *AlibabaEinvoiceTaxOptSalaryrequestSingleacceptAPIRequest{
    return &AlibabaEinvoiceTaxOptSalaryrequestSingleacceptAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceTaxOptSalaryrequestSingleacceptAPIRequest) GetApiMethodName() string {
    return "alibaba.einvoice.tax.opt.salaryrequest.singleaccept"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceTaxOptSalaryrequestSingleacceptAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamTaxOptimizationSingleDetailSalaryPaymentAccessDTO Setter
// 入参
func (r *AlibabaEinvoiceTaxOptSalaryrequestSingleacceptAPIRequest) SetParamTaxOptimizationSingleDetailSalaryPaymentAccessDTO(_paramTaxOptimizationSingleDetailSalaryPaymentAccessDTO *TaxOptimizationSingleDetailSalaryPaymentAccessDTO) error {
    r._paramTaxOptimizationSingleDetailSalaryPaymentAccessDTO = _paramTaxOptimizationSingleDetailSalaryPaymentAccessDTO
    r.Set("param_tax_optimization_single_detail_salary_payment_access_d_t_o", _paramTaxOptimizationSingleDetailSalaryPaymentAccessDTO)
    return nil
}

// ParamTaxOptimizationSingleDetailSalaryPaymentAccessDTO Getter
func (r AlibabaEinvoiceTaxOptSalaryrequestSingleacceptAPIRequest) GetParamTaxOptimizationSingleDetailSalaryPaymentAccessDTO() *TaxOptimizationSingleDetailSalaryPaymentAccessDTO {
    return r._paramTaxOptimizationSingleDetailSalaryPaymentAccessDTO
}
