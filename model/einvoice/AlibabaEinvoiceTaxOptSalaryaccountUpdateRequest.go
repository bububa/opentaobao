package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新用户发薪资产 API请求
alibaba.einvoice.tax.opt.salaryaccount.update

更新用户的发薪账号
*/
type AlibabaEinvoiceTaxOptSalaryaccountUpdateRequest struct {
    model.Params
    // 入参
    _paramTaxOptimizationEmployeeAssetUpdateDTO   *TaxOptimizationEmployeeAssetUpdateDTO
}

// 初始化AlibabaEinvoiceTaxOptSalaryaccountUpdateRequest对象
func NewAlibabaEinvoiceTaxOptSalaryaccountUpdateRequest() *AlibabaEinvoiceTaxOptSalaryaccountUpdateRequest{
    return &AlibabaEinvoiceTaxOptSalaryaccountUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceTaxOptSalaryaccountUpdateRequest) GetApiMethodName() string {
    return "alibaba.einvoice.tax.opt.salaryaccount.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceTaxOptSalaryaccountUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamTaxOptimizationEmployeeAssetUpdateDTO Setter
// 入参
func (r *AlibabaEinvoiceTaxOptSalaryaccountUpdateRequest) SetParamTaxOptimizationEmployeeAssetUpdateDTO(_paramTaxOptimizationEmployeeAssetUpdateDTO *TaxOptimizationEmployeeAssetUpdateDTO) error {
    r._paramTaxOptimizationEmployeeAssetUpdateDTO = _paramTaxOptimizationEmployeeAssetUpdateDTO
    r.Set("param_tax_optimization_employee_asset_update_d_t_o", _paramTaxOptimizationEmployeeAssetUpdateDTO)
    return nil
}

// ParamTaxOptimizationEmployeeAssetUpdateDTO Getter
func (r AlibabaEinvoiceTaxOptSalaryaccountUpdateRequest) GetParamTaxOptimizationEmployeeAssetUpdateDTO() *TaxOptimizationEmployeeAssetUpdateDTO {
    return r._paramTaxOptimizationEmployeeAssetUpdateDTO
}
