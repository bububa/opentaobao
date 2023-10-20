package einvoice

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaEinvoiceTaxOptSalaryaccountQueryAPIRequest) Reset() {
	r._paramTaxOptimizationQueryPaySalaryAccountDTO = nil
	r.Params.ToZero()
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

var poolAlibabaEinvoiceTaxOptSalaryaccountQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaEinvoiceTaxOptSalaryaccountQueryRequest()
	},
}

// GetAlibabaEinvoiceTaxOptSalaryaccountQueryRequest 从 sync.Pool 获取 AlibabaEinvoiceTaxOptSalaryaccountQueryAPIRequest
func GetAlibabaEinvoiceTaxOptSalaryaccountQueryAPIRequest() *AlibabaEinvoiceTaxOptSalaryaccountQueryAPIRequest {
	return poolAlibabaEinvoiceTaxOptSalaryaccountQueryAPIRequest.Get().(*AlibabaEinvoiceTaxOptSalaryaccountQueryAPIRequest)
}

// ReleaseAlibabaEinvoiceTaxOptSalaryaccountQueryAPIRequest 将 AlibabaEinvoiceTaxOptSalaryaccountQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaEinvoiceTaxOptSalaryaccountQueryAPIRequest(v *AlibabaEinvoiceTaxOptSalaryaccountQueryAPIRequest) {
	v.Reset()
	poolAlibabaEinvoiceTaxOptSalaryaccountQueryAPIRequest.Put(v)
}
