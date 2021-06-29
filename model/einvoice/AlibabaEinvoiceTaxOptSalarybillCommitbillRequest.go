package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
提交发薪账单 APIRequest
alibaba.einvoice.tax.opt.salarybill.commitbill

提交发薪账单
*/
type AlibabaEinvoiceTaxOptSalarybillCommitbillRequest struct {
    model.Params

    // 入参
    paramTaxOptimizationSalaryBillCommitReqDTO   *TaxOptimizationSalaryBillCommitReqDto 

}

func NewAlibabaEinvoiceTaxOptSalarybillCommitbillRequest() *AlibabaEinvoiceTaxOptSalarybillCommitbillRequest{
    return &AlibabaEinvoiceTaxOptSalarybillCommitbillRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaEinvoiceTaxOptSalarybillCommitbillRequest) GetApiMethodName() string {
    return "alibaba.einvoice.tax.opt.salarybill.commitbill"
}

func (r AlibabaEinvoiceTaxOptSalarybillCommitbillRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaEinvoiceTaxOptSalarybillCommitbillRequest) SetParamTaxOptimizationSalaryBillCommitReqDTO(paramTaxOptimizationSalaryBillCommitReqDTO *TaxOptimizationSalaryBillCommitReqDto) error {
    r.paramTaxOptimizationSalaryBillCommitReqDTO = paramTaxOptimizationSalaryBillCommitReqDTO
    r.Set("param_tax_optimization_salary_bill_commit_req_d_t_o", paramTaxOptimizationSalaryBillCommitReqDTO)
    return nil
}

func (r AlibabaEinvoiceTaxOptSalarybillCommitbillRequest) GetParamTaxOptimizationSalaryBillCommitReqDTO() *TaxOptimizationSalaryBillCommitReqDto {
    return r.paramTaxOptimizationSalaryBillCommitReqDTO
}

