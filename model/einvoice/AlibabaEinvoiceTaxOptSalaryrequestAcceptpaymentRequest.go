package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
受理发薪 API请求
alibaba.einvoice.tax.opt.salaryrequest.acceptpayment

发薪受理接口
*/
type AlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentRequest struct {
    model.Params
    // 请求入参
    _paramTaxOptimizationSalaryPaymentAccessReqDTO   *TaxOptimizationSalaryPaymentAccessReqDTO
}

// 初始化AlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentRequest对象
func NewAlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentRequest() *AlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentRequest{
    return &AlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentRequest) GetApiMethodName() string {
    return "alibaba.einvoice.tax.opt.salaryrequest.acceptpayment"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamTaxOptimizationSalaryPaymentAccessReqDTO Setter
// 请求入参
func (r *AlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentRequest) SetParamTaxOptimizationSalaryPaymentAccessReqDTO(_paramTaxOptimizationSalaryPaymentAccessReqDTO *TaxOptimizationSalaryPaymentAccessReqDTO) error {
    r._paramTaxOptimizationSalaryPaymentAccessReqDTO = _paramTaxOptimizationSalaryPaymentAccessReqDTO
    r.Set("param_tax_optimization_salary_payment_access_req_d_t_o", _paramTaxOptimizationSalaryPaymentAccessReqDTO)
    return nil
}

// ParamTaxOptimizationSalaryPaymentAccessReqDTO Getter
func (r AlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentRequest) GetParamTaxOptimizationSalaryPaymentAccessReqDTO() *TaxOptimizationSalaryPaymentAccessReqDTO {
    return r._paramTaxOptimizationSalaryPaymentAccessReqDTO
}
