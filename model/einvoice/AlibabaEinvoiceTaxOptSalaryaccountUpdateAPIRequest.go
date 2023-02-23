package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceTaxOptSalaryaccountUpdateAPIRequest 更新用户发薪资产 API请求
// alibaba.einvoice.tax.opt.salaryaccount.update
//
// 更新用户的发薪账号
type AlibabaEinvoiceTaxOptSalaryaccountUpdateAPIRequest struct {
	model.Params
	// 入参
	_paramTaxOptimizationEmployeeAssetUpdateDTO *TaxOptimizationEmployeeAssetUpdateDto
}

// NewAlibabaEinvoiceTaxOptSalaryaccountUpdateRequest 初始化AlibabaEinvoiceTaxOptSalaryaccountUpdateAPIRequest对象
func NewAlibabaEinvoiceTaxOptSalaryaccountUpdateRequest() *AlibabaEinvoiceTaxOptSalaryaccountUpdateAPIRequest {
	return &AlibabaEinvoiceTaxOptSalaryaccountUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceTaxOptSalaryaccountUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.tax.opt.salaryaccount.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceTaxOptSalaryaccountUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaEinvoiceTaxOptSalaryaccountUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamTaxOptimizationEmployeeAssetUpdateDTO is ParamTaxOptimizationEmployeeAssetUpdateDTO Setter
// 入参
func (r *AlibabaEinvoiceTaxOptSalaryaccountUpdateAPIRequest) SetParamTaxOptimizationEmployeeAssetUpdateDTO(_paramTaxOptimizationEmployeeAssetUpdateDTO *TaxOptimizationEmployeeAssetUpdateDto) error {
	r._paramTaxOptimizationEmployeeAssetUpdateDTO = _paramTaxOptimizationEmployeeAssetUpdateDTO
	r.Set("param_tax_optimization_employee_asset_update_d_t_o", _paramTaxOptimizationEmployeeAssetUpdateDTO)
	return nil
}

// GetParamTaxOptimizationEmployeeAssetUpdateDTO ParamTaxOptimizationEmployeeAssetUpdateDTO Getter
func (r AlibabaEinvoiceTaxOptSalaryaccountUpdateAPIRequest) GetParamTaxOptimizationEmployeeAssetUpdateDTO() *TaxOptimizationEmployeeAssetUpdateDto {
	return r._paramTaxOptimizationEmployeeAssetUpdateDTO
}
