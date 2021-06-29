package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
单明细发薪受理 APIRequest
alibaba.einvoice.tax.opt.salaryrequest.singleaccept

单明细发薪受理
*/
type AlibabaEinvoiceTaxOptSalaryrequestSingleacceptRequest struct {
    model.Params

    // 入参
    paramTaxOptimizationSingleDetailSalaryPaymentAccessDTO   *TaxOptimizationSingleDetailSalaryPaymentAccessDto 

}

func NewAlibabaEinvoiceTaxOptSalaryrequestSingleacceptRequest() *AlibabaEinvoiceTaxOptSalaryrequestSingleacceptRequest{
    return &AlibabaEinvoiceTaxOptSalaryrequestSingleacceptRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaEinvoiceTaxOptSalaryrequestSingleacceptRequest) GetApiMethodName() string {
    return "alibaba.einvoice.tax.opt.salaryrequest.singleaccept"
}

func (r AlibabaEinvoiceTaxOptSalaryrequestSingleacceptRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaEinvoiceTaxOptSalaryrequestSingleacceptRequest) SetParamTaxOptimizationSingleDetailSalaryPaymentAccessDTO(paramTaxOptimizationSingleDetailSalaryPaymentAccessDTO *TaxOptimizationSingleDetailSalaryPaymentAccessDto) error {
    r.paramTaxOptimizationSingleDetailSalaryPaymentAccessDTO = paramTaxOptimizationSingleDetailSalaryPaymentAccessDTO
    r.Set("param_tax_optimization_single_detail_salary_payment_access_d_t_o", paramTaxOptimizationSingleDetailSalaryPaymentAccessDTO)
    return nil
}

func (r AlibabaEinvoiceTaxOptSalaryrequestSingleacceptRequest) GetParamTaxOptimizationSingleDetailSalaryPaymentAccessDTO() *TaxOptimizationSingleDetailSalaryPaymentAccessDto {
    return r.paramTaxOptimizationSingleDetailSalaryPaymentAccessDTO
}

