package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新用户发薪资产 APIRequest
alibaba.einvoice.tax.opt.salaryaccount.update

更新用户的发薪账号
*/
type AlibabaEinvoiceTaxOptSalaryaccountUpdateRequest struct {
    model.Params

    // 入参
    paramTaxOptimizationEmployeeAssetUpdateDTO   *TaxOptimizationEmployeeAssetUpdateDto 

}

func NewAlibabaEinvoiceTaxOptSalaryaccountUpdateRequest() *AlibabaEinvoiceTaxOptSalaryaccountUpdateRequest{
    return &AlibabaEinvoiceTaxOptSalaryaccountUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaEinvoiceTaxOptSalaryaccountUpdateRequest) GetApiMethodName() string {
    return "alibaba.einvoice.tax.opt.salaryaccount.update"
}

func (r AlibabaEinvoiceTaxOptSalaryaccountUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaEinvoiceTaxOptSalaryaccountUpdateRequest) SetParamTaxOptimizationEmployeeAssetUpdateDTO(paramTaxOptimizationEmployeeAssetUpdateDTO *TaxOptimizationEmployeeAssetUpdateDto) error {
    r.paramTaxOptimizationEmployeeAssetUpdateDTO = paramTaxOptimizationEmployeeAssetUpdateDTO
    r.Set("param_tax_optimization_employee_asset_update_d_t_o", paramTaxOptimizationEmployeeAssetUpdateDTO)
    return nil
}

func (r AlibabaEinvoiceTaxOptSalaryaccountUpdateRequest) GetParamTaxOptimizationEmployeeAssetUpdateDTO() *TaxOptimizationEmployeeAssetUpdateDto {
    return r.paramTaxOptimizationEmployeeAssetUpdateDTO
}

