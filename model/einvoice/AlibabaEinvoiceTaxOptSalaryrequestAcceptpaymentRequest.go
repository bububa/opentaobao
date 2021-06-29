package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
受理发薪 APIRequest
alibaba.einvoice.tax.opt.salaryrequest.acceptpayment

发薪受理接口
*/
type AlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentRequest struct {
    model.Params

    // 请求入参
    paramTaxOptimizationSalaryPaymentAccessReqDTO   *TaxOptimizationSalaryPaymentAccessReqDto 

}

func NewAlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentRequest() *AlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentRequest{
    return &AlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentRequest) GetApiMethodName() string {
    return "alibaba.einvoice.tax.opt.salaryrequest.acceptpayment"
}

func (r AlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentRequest) SetParamTaxOptimizationSalaryPaymentAccessReqDTO(paramTaxOptimizationSalaryPaymentAccessReqDTO *TaxOptimizationSalaryPaymentAccessReqDto) error {
    r.paramTaxOptimizationSalaryPaymentAccessReqDTO = paramTaxOptimizationSalaryPaymentAccessReqDTO
    r.Set("param_tax_optimization_salary_payment_access_req_d_t_o", paramTaxOptimizationSalaryPaymentAccessReqDTO)
    return nil
}

func (r AlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentRequest) GetParamTaxOptimizationSalaryPaymentAccessReqDTO() *TaxOptimizationSalaryPaymentAccessReqDto {
    return r.paramTaxOptimizationSalaryPaymentAccessReqDTO
}

