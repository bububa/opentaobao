package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询发薪结果 APIRequest
alibaba.einvoice.tax.opt.salaryresult.query

查询发薪结果
*/
type AlibabaEinvoiceTaxOptSalaryresultQueryRequest struct {
    model.Params

    // 发薪流水号
    detailIdList   []string 

    // 业务方编码
    employerCode   string 

}

func NewAlibabaEinvoiceTaxOptSalaryresultQueryRequest() *AlibabaEinvoiceTaxOptSalaryresultQueryRequest{
    return &AlibabaEinvoiceTaxOptSalaryresultQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaEinvoiceTaxOptSalaryresultQueryRequest) GetApiMethodName() string {
    return "alibaba.einvoice.tax.opt.salaryresult.query"
}

func (r AlibabaEinvoiceTaxOptSalaryresultQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaEinvoiceTaxOptSalaryresultQueryRequest) SetDetailIdList(detailIdList []string) error {
    r.detailIdList = detailIdList
    r.Set("detail_id_list", detailIdList)
    return nil
}

func (r AlibabaEinvoiceTaxOptSalaryresultQueryRequest) GetDetailIdList() []string {
    return r.detailIdList
}

func (r *AlibabaEinvoiceTaxOptSalaryresultQueryRequest) SetEmployerCode(employerCode string) error {
    r.employerCode = employerCode
    r.Set("employer_code", employerCode)
    return nil
}

func (r AlibabaEinvoiceTaxOptSalaryresultQueryRequest) GetEmployerCode() string {
    return r.employerCode
}

