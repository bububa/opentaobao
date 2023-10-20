package einvoice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceTaxOptSalarybillCommitbillAPIRequest 提交发薪账单 API请求
// alibaba.einvoice.tax.opt.salarybill.commitbill
//
// 提交发薪账单
type AlibabaEinvoiceTaxOptSalarybillCommitbillAPIRequest struct {
	model.Params
	// 入参
	_paramTaxOptimizationSalaryBillCommitReqDTO *TaxOptimizationSalaryBillCommitReqDto
}

// NewAlibabaEinvoiceTaxOptSalarybillCommitbillRequest 初始化AlibabaEinvoiceTaxOptSalarybillCommitbillAPIRequest对象
func NewAlibabaEinvoiceTaxOptSalarybillCommitbillRequest() *AlibabaEinvoiceTaxOptSalarybillCommitbillAPIRequest {
	return &AlibabaEinvoiceTaxOptSalarybillCommitbillAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaEinvoiceTaxOptSalarybillCommitbillAPIRequest) Reset() {
	r._paramTaxOptimizationSalaryBillCommitReqDTO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceTaxOptSalarybillCommitbillAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.tax.opt.salarybill.commitbill"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceTaxOptSalarybillCommitbillAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaEinvoiceTaxOptSalarybillCommitbillAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamTaxOptimizationSalaryBillCommitReqDTO is ParamTaxOptimizationSalaryBillCommitReqDTO Setter
// 入参
func (r *AlibabaEinvoiceTaxOptSalarybillCommitbillAPIRequest) SetParamTaxOptimizationSalaryBillCommitReqDTO(_paramTaxOptimizationSalaryBillCommitReqDTO *TaxOptimizationSalaryBillCommitReqDto) error {
	r._paramTaxOptimizationSalaryBillCommitReqDTO = _paramTaxOptimizationSalaryBillCommitReqDTO
	r.Set("param_tax_optimization_salary_bill_commit_req_d_t_o", _paramTaxOptimizationSalaryBillCommitReqDTO)
	return nil
}

// GetParamTaxOptimizationSalaryBillCommitReqDTO ParamTaxOptimizationSalaryBillCommitReqDTO Getter
func (r AlibabaEinvoiceTaxOptSalarybillCommitbillAPIRequest) GetParamTaxOptimizationSalaryBillCommitReqDTO() *TaxOptimizationSalaryBillCommitReqDto {
	return r._paramTaxOptimizationSalaryBillCommitReqDTO
}

var poolAlibabaEinvoiceTaxOptSalarybillCommitbillAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaEinvoiceTaxOptSalarybillCommitbillRequest()
	},
}

// GetAlibabaEinvoiceTaxOptSalarybillCommitbillRequest 从 sync.Pool 获取 AlibabaEinvoiceTaxOptSalarybillCommitbillAPIRequest
func GetAlibabaEinvoiceTaxOptSalarybillCommitbillAPIRequest() *AlibabaEinvoiceTaxOptSalarybillCommitbillAPIRequest {
	return poolAlibabaEinvoiceTaxOptSalarybillCommitbillAPIRequest.Get().(*AlibabaEinvoiceTaxOptSalarybillCommitbillAPIRequest)
}

// ReleaseAlibabaEinvoiceTaxOptSalarybillCommitbillAPIRequest 将 AlibabaEinvoiceTaxOptSalarybillCommitbillAPIRequest 放入 sync.Pool
func ReleaseAlibabaEinvoiceTaxOptSalarybillCommitbillAPIRequest(v *AlibabaEinvoiceTaxOptSalarybillCommitbillAPIRequest) {
	v.Reset()
	poolAlibabaEinvoiceTaxOptSalarybillCommitbillAPIRequest.Put(v)
}
