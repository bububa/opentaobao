package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceTaxOptSalaryaccountQueryAPIRequest 查询用户的发薪账号 API请求
// alibaba.einvoice.tax.opt.salaryaccount.query
//
// 查询用户的发薪账号状态
type AlibabaEinvoiceTaxOptSalaryaccountQueryAPIRequest struct {
	model.Params
	// 入参
	_paramTaxOptimizationQueryPaySalaryAccountDTO *TaxOptimizationQueryPaySalaryAccountDto
}

// NewAlibabaEinvoiceTaxOptSalaryaccountQueryRequest 初始化AlibabaEinvoiceTaxOptSalaryaccountQueryAPIRequest对象
func NewAlibabaEinvoiceTaxOptSalaryaccountQueryRequest() *AlibabaEinvoiceTaxOptSalaryaccountQueryAPIRequest {
	return &AlibabaEinvoiceTaxOptSalaryaccountQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceTaxOptSalaryaccountQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.tax.opt.salaryaccount.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceTaxOptSalaryaccountQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaEinvoiceTaxOptSalaryaccountQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamTaxOptimizationQueryPaySalaryAccountDTO is ParamTaxOptimizationQueryPaySalaryAccountDTO Setter
// 入参
func (r *AlibabaEinvoiceTaxOptSalaryaccountQueryAPIRequest) SetParamTaxOptimizationQueryPaySalaryAccountDTO(_paramTaxOptimizationQueryPaySalaryAccountDTO *TaxOptimizationQueryPaySalaryAccountDto) error {
	r._paramTaxOptimizationQueryPaySalaryAccountDTO = _paramTaxOptimizationQueryPaySalaryAccountDTO
	r.Set("param_tax_optimization_query_pay_salary_account_d_t_o", _paramTaxOptimizationQueryPaySalaryAccountDTO)
	return nil
}

// GetParamTaxOptimizationQueryPaySalaryAccountDTO ParamTaxOptimizationQueryPaySalaryAccountDTO Getter
func (r AlibabaEinvoiceTaxOptSalaryaccountQueryAPIRequest) GetParamTaxOptimizationQueryPaySalaryAccountDTO() *TaxOptimizationQueryPaySalaryAccountDto {
	return r._paramTaxOptimizationQueryPaySalaryAccountDTO
}
