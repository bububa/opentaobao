package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaeinvoicetaxoptsalaryrequestsingleacceptAPIRequest 单明细发薪受理 API请求
// alibaba.einvoice.tax.opt.salaryrequest.singleaccept
//
// 单明细发薪受理
type AlibabaeinvoicetaxoptsalaryrequestsingleacceptAPIRequest struct {
	model.Params
	// 入参
	_paramTaxOptimizationSingleDetailSalaryPaymentAccessDTO *TaxOptimizationSingleDetailSalaryPaymentAccessDto
}

// NewAlibabaeinvoicetaxoptsalaryrequestsingleacceptRequest 初始化AlibabaeinvoicetaxoptsalaryrequestsingleacceptAPIRequest对象
func NewAlibabaeinvoicetaxoptsalaryrequestsingleacceptRequest() *AlibabaeinvoicetaxoptsalaryrequestsingleacceptAPIRequest {
	return &AlibabaeinvoicetaxoptsalaryrequestsingleacceptAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaeinvoicetaxoptsalaryrequestsingleacceptAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.tax.opt.salaryrequest.singleaccept"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaeinvoicetaxoptsalaryrequestsingleacceptAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaeinvoicetaxoptsalaryrequestsingleacceptAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamTaxOptimizationSingleDetailSalaryPaymentAccessDTO is ParamTaxOptimizationSingleDetailSalaryPaymentAccessDTO Setter
// 入参
func (r *AlibabaeinvoicetaxoptsalaryrequestsingleacceptAPIRequest) SetParamTaxOptimizationSingleDetailSalaryPaymentAccessDTO(_paramTaxOptimizationSingleDetailSalaryPaymentAccessDTO *TaxOptimizationSingleDetailSalaryPaymentAccessDto) error {
	r._paramTaxOptimizationSingleDetailSalaryPaymentAccessDTO = _paramTaxOptimizationSingleDetailSalaryPaymentAccessDTO
	r.Set("param_tax_optimization_single_detail_salary_payment_access_d_t_o", _paramTaxOptimizationSingleDetailSalaryPaymentAccessDTO)
	return nil
}

// GetParamTaxOptimizationSingleDetailSalaryPaymentAccessDTO ParamTaxOptimizationSingleDetailSalaryPaymentAccessDTO Getter
func (r AlibabaeinvoicetaxoptsalaryrequestsingleacceptAPIRequest) GetParamTaxOptimizationSingleDetailSalaryPaymentAccessDTO() *TaxOptimizationSingleDetailSalaryPaymentAccessDto {
	return r._paramTaxOptimizationSingleDetailSalaryPaymentAccessDTO
}
