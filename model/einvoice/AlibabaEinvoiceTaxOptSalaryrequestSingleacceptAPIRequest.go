package einvoice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceTaxOptSalaryrequestSingleacceptAPIRequest 单明细发薪受理 API请求
// alibaba.einvoice.tax.opt.salaryrequest.singleaccept
//
// 单明细发薪受理
type AlibabaEinvoiceTaxOptSalaryrequestSingleacceptAPIRequest struct {
	model.Params
	// 入参
	_paramTaxOptimizationSingleDetailSalaryPaymentAccessDTO *TaxOptimizationSingleDetailSalaryPaymentAccessDto
}

// NewAlibabaEinvoiceTaxOptSalaryrequestSingleacceptRequest 初始化AlibabaEinvoiceTaxOptSalaryrequestSingleacceptAPIRequest对象
func NewAlibabaEinvoiceTaxOptSalaryrequestSingleacceptRequest() *AlibabaEinvoiceTaxOptSalaryrequestSingleacceptAPIRequest {
	return &AlibabaEinvoiceTaxOptSalaryrequestSingleacceptAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaEinvoiceTaxOptSalaryrequestSingleacceptAPIRequest) Reset() {
	r._paramTaxOptimizationSingleDetailSalaryPaymentAccessDTO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceTaxOptSalaryrequestSingleacceptAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.tax.opt.salaryrequest.singleaccept"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceTaxOptSalaryrequestSingleacceptAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaEinvoiceTaxOptSalaryrequestSingleacceptAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamTaxOptimizationSingleDetailSalaryPaymentAccessDTO is ParamTaxOptimizationSingleDetailSalaryPaymentAccessDTO Setter
// 入参
func (r *AlibabaEinvoiceTaxOptSalaryrequestSingleacceptAPIRequest) SetParamTaxOptimizationSingleDetailSalaryPaymentAccessDTO(_paramTaxOptimizationSingleDetailSalaryPaymentAccessDTO *TaxOptimizationSingleDetailSalaryPaymentAccessDto) error {
	r._paramTaxOptimizationSingleDetailSalaryPaymentAccessDTO = _paramTaxOptimizationSingleDetailSalaryPaymentAccessDTO
	r.Set("param_tax_optimization_single_detail_salary_payment_access_d_t_o", _paramTaxOptimizationSingleDetailSalaryPaymentAccessDTO)
	return nil
}

// GetParamTaxOptimizationSingleDetailSalaryPaymentAccessDTO ParamTaxOptimizationSingleDetailSalaryPaymentAccessDTO Getter
func (r AlibabaEinvoiceTaxOptSalaryrequestSingleacceptAPIRequest) GetParamTaxOptimizationSingleDetailSalaryPaymentAccessDTO() *TaxOptimizationSingleDetailSalaryPaymentAccessDto {
	return r._paramTaxOptimizationSingleDetailSalaryPaymentAccessDTO
}

var poolAlibabaEinvoiceTaxOptSalaryrequestSingleacceptAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaEinvoiceTaxOptSalaryrequestSingleacceptRequest()
	},
}

// GetAlibabaEinvoiceTaxOptSalaryrequestSingleacceptRequest 从 sync.Pool 获取 AlibabaEinvoiceTaxOptSalaryrequestSingleacceptAPIRequest
func GetAlibabaEinvoiceTaxOptSalaryrequestSingleacceptAPIRequest() *AlibabaEinvoiceTaxOptSalaryrequestSingleacceptAPIRequest {
	return poolAlibabaEinvoiceTaxOptSalaryrequestSingleacceptAPIRequest.Get().(*AlibabaEinvoiceTaxOptSalaryrequestSingleacceptAPIRequest)
}

// ReleaseAlibabaEinvoiceTaxOptSalaryrequestSingleacceptAPIRequest 将 AlibabaEinvoiceTaxOptSalaryrequestSingleacceptAPIRequest 放入 sync.Pool
func ReleaseAlibabaEinvoiceTaxOptSalaryrequestSingleacceptAPIRequest(v *AlibabaEinvoiceTaxOptSalaryrequestSingleacceptAPIRequest) {
	v.Reset()
	poolAlibabaEinvoiceTaxOptSalaryrequestSingleacceptAPIRequest.Put(v)
}
