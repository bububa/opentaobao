package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaeinvoicetaxoptsalaryaccountupdateAPIRequest 更新用户发薪资产 API请求
// alibaba.einvoice.tax.opt.salaryaccount.update
//
// 更新用户的发薪账号
type AlibabaeinvoicetaxoptsalaryaccountupdateAPIRequest struct {
	model.Params
	// 入参
	_paramTaxOptimizationEmployeeAssetUpdateDTO *TaxOptimizationEmployeeAssetUpdateDto
}

// NewAlibabaeinvoicetaxoptsalaryaccountupdateRequest 初始化AlibabaeinvoicetaxoptsalaryaccountupdateAPIRequest对象
func NewAlibabaeinvoicetaxoptsalaryaccountupdateRequest() *AlibabaeinvoicetaxoptsalaryaccountupdateAPIRequest {
	return &AlibabaeinvoicetaxoptsalaryaccountupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaeinvoicetaxoptsalaryaccountupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.tax.opt.salaryaccount.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaeinvoicetaxoptsalaryaccountupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaeinvoicetaxoptsalaryaccountupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamTaxOptimizationEmployeeAssetUpdateDTO is ParamTaxOptimizationEmployeeAssetUpdateDTO Setter
// 入参
func (r *AlibabaeinvoicetaxoptsalaryaccountupdateAPIRequest) SetParamTaxOptimizationEmployeeAssetUpdateDTO(_paramTaxOptimizationEmployeeAssetUpdateDTO *TaxOptimizationEmployeeAssetUpdateDto) error {
	r._paramTaxOptimizationEmployeeAssetUpdateDTO = _paramTaxOptimizationEmployeeAssetUpdateDTO
	r.Set("param_tax_optimization_employee_asset_update_d_t_o", _paramTaxOptimizationEmployeeAssetUpdateDTO)
	return nil
}

// GetParamTaxOptimizationEmployeeAssetUpdateDTO ParamTaxOptimizationEmployeeAssetUpdateDTO Getter
func (r AlibabaeinvoicetaxoptsalaryaccountupdateAPIRequest) GetParamTaxOptimizationEmployeeAssetUpdateDTO() *TaxOptimizationEmployeeAssetUpdateDto {
	return r._paramTaxOptimizationEmployeeAssetUpdateDTO
}
