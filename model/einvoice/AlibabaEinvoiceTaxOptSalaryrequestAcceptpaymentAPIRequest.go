package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaeinvoicetaxoptsalaryrequestacceptpaymentAPIRequest 受理发薪 API请求
// alibaba.einvoice.tax.opt.salaryrequest.acceptpayment
//
// 发薪受理接口
type AlibabaeinvoicetaxoptsalaryrequestacceptpaymentAPIRequest struct {
	model.Params
	// 请求入参
	_paramTaxOptimizationSalaryPaymentAccessReqDTO *TaxOptimizationSalaryPaymentAccessReqDto
}

// NewAlibabaeinvoicetaxoptsalaryrequestacceptpaymentRequest 初始化AlibabaeinvoicetaxoptsalaryrequestacceptpaymentAPIRequest对象
func NewAlibabaeinvoicetaxoptsalaryrequestacceptpaymentRequest() *AlibabaeinvoicetaxoptsalaryrequestacceptpaymentAPIRequest {
	return &AlibabaeinvoicetaxoptsalaryrequestacceptpaymentAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaeinvoicetaxoptsalaryrequestacceptpaymentAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.tax.opt.salaryrequest.acceptpayment"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaeinvoicetaxoptsalaryrequestacceptpaymentAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaeinvoicetaxoptsalaryrequestacceptpaymentAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamTaxOptimizationSalaryPaymentAccessReqDTO is ParamTaxOptimizationSalaryPaymentAccessReqDTO Setter
// 请求入参
func (r *AlibabaeinvoicetaxoptsalaryrequestacceptpaymentAPIRequest) SetParamTaxOptimizationSalaryPaymentAccessReqDTO(_paramTaxOptimizationSalaryPaymentAccessReqDTO *TaxOptimizationSalaryPaymentAccessReqDto) error {
	r._paramTaxOptimizationSalaryPaymentAccessReqDTO = _paramTaxOptimizationSalaryPaymentAccessReqDTO
	r.Set("param_tax_optimization_salary_payment_access_req_d_t_o", _paramTaxOptimizationSalaryPaymentAccessReqDTO)
	return nil
}

// GetParamTaxOptimizationSalaryPaymentAccessReqDTO ParamTaxOptimizationSalaryPaymentAccessReqDTO Getter
func (r AlibabaeinvoicetaxoptsalaryrequestacceptpaymentAPIRequest) GetParamTaxOptimizationSalaryPaymentAccessReqDTO() *TaxOptimizationSalaryPaymentAccessReqDto {
	return r._paramTaxOptimizationSalaryPaymentAccessReqDTO
}
