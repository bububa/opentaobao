package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
提交发薪账单 API请求
alibaba.einvoice.tax.opt.salarybill.commitbill

提交发薪账单
*/
type AlibabaEinvoiceTaxOptSalarybillCommitbillRequest struct {
    model.Params
    // 入参
    _paramTaxOptimizationSalaryBillCommitReqDTO   *TaxOptimizationSalaryBillCommitReqDto
}

// 初始化AlibabaEinvoiceTaxOptSalarybillCommitbillRequest对象
func NewAlibabaEinvoiceTaxOptSalarybillCommitbillRequest() *AlibabaEinvoiceTaxOptSalarybillCommitbillRequest{
    return &AlibabaEinvoiceTaxOptSalarybillCommitbillRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceTaxOptSalarybillCommitbillRequest) GetApiMethodName() string {
    return "alibaba.einvoice.tax.opt.salarybill.commitbill"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceTaxOptSalarybillCommitbillRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamTaxOptimizationSalaryBillCommitReqDTO Setter
// 入参
func (r *AlibabaEinvoiceTaxOptSalarybillCommitbillRequest) SetParamTaxOptimizationSalaryBillCommitReqDTO(_paramTaxOptimizationSalaryBillCommitReqDTO *TaxOptimizationSalaryBillCommitReqDto) error {
    r._paramTaxOptimizationSalaryBillCommitReqDTO = _paramTaxOptimizationSalaryBillCommitReqDTO
    r.Set("param_tax_optimization_salary_bill_commit_req_d_t_o", _paramTaxOptimizationSalaryBillCommitReqDTO)
    return nil
}

// ParamTaxOptimizationSalaryBillCommitReqDTO Getter
func (r AlibabaEinvoiceTaxOptSalarybillCommitbillRequest) GetParamTaxOptimizationSalaryBillCommitReqDTO() *TaxOptimizationSalaryBillCommitReqDto {
    return r._paramTaxOptimizationSalaryBillCommitReqDTO
}
