package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaeinvoicetaxoptsalarybillcommitbillAPIRequest 提交发薪账单 API请求
// alibaba.einvoice.tax.opt.salarybill.commitbill
//
// 提交发薪账单
type AlibabaeinvoicetaxoptsalarybillcommitbillAPIRequest struct {
	model.Params
	// 入参
	_paramTaxOptimizationSalaryBillCommitReqDTO *TaxOptimizationSalaryBillCommitReqDto
}

// NewAlibabaeinvoicetaxoptsalarybillcommitbillRequest 初始化AlibabaeinvoicetaxoptsalarybillcommitbillAPIRequest对象
func NewAlibabaeinvoicetaxoptsalarybillcommitbillRequest() *AlibabaeinvoicetaxoptsalarybillcommitbillAPIRequest {
	return &AlibabaeinvoicetaxoptsalarybillcommitbillAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaeinvoicetaxoptsalarybillcommitbillAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.tax.opt.salarybill.commitbill"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaeinvoicetaxoptsalarybillcommitbillAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaeinvoicetaxoptsalarybillcommitbillAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamTaxOptimizationSalaryBillCommitReqDTO is ParamTaxOptimizationSalaryBillCommitReqDTO Setter
// 入参
func (r *AlibabaeinvoicetaxoptsalarybillcommitbillAPIRequest) SetParamTaxOptimizationSalaryBillCommitReqDTO(_paramTaxOptimizationSalaryBillCommitReqDTO *TaxOptimizationSalaryBillCommitReqDto) error {
	r._paramTaxOptimizationSalaryBillCommitReqDTO = _paramTaxOptimizationSalaryBillCommitReqDTO
	r.Set("param_tax_optimization_salary_bill_commit_req_d_t_o", _paramTaxOptimizationSalaryBillCommitReqDTO)
	return nil
}

// GetParamTaxOptimizationSalaryBillCommitReqDTO ParamTaxOptimizationSalaryBillCommitReqDTO Getter
func (r AlibabaeinvoicetaxoptsalarybillcommitbillAPIRequest) GetParamTaxOptimizationSalaryBillCommitReqDTO() *TaxOptimizationSalaryBillCommitReqDto {
	return r._paramTaxOptimizationSalaryBillCommitReqDTO
}
