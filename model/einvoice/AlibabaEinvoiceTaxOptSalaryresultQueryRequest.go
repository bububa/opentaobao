package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询发薪结果 API请求
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

// 初始化AlibabaEinvoiceTaxOptSalaryresultQueryRequest对象
func NewAlibabaEinvoiceTaxOptSalaryresultQueryRequest() *AlibabaEinvoiceTaxOptSalaryresultQueryRequest{
    return &AlibabaEinvoiceTaxOptSalaryresultQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceTaxOptSalaryresultQueryRequest) GetApiMethodName() string {
    return "alibaba.einvoice.tax.opt.salaryresult.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceTaxOptSalaryresultQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DetailIdList Setter
// 发薪流水号
func (r *AlibabaEinvoiceTaxOptSalaryresultQueryRequest) SetDetailIdList(detailIdList []string) error {
    r.detailIdList = detailIdList
    r.Set("detail_id_list", detailIdList)
    return nil
}

// DetailIdList Getter
func (r AlibabaEinvoiceTaxOptSalaryresultQueryRequest) GetDetailIdList() []string {
    return r.detailIdList
}
// EmployerCode Setter
// 业务方编码
func (r *AlibabaEinvoiceTaxOptSalaryresultQueryRequest) SetEmployerCode(employerCode string) error {
    r.employerCode = employerCode
    r.Set("employer_code", employerCode)
    return nil
}

// EmployerCode Getter
func (r AlibabaEinvoiceTaxOptSalaryresultQueryRequest) GetEmployerCode() string {
    return r.employerCode
}
