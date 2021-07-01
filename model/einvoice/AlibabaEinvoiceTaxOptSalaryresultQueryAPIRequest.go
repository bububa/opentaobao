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
type AlibabaEinvoiceTaxOptSalaryresultQueryAPIRequest struct {
    model.Params
    // 发薪流水号
    _detailIdList   []string
    // 业务方编码
    _employerCode   string
}

// 初始化AlibabaEinvoiceTaxOptSalaryresultQueryAPIRequest对象
func NewAlibabaEinvoiceTaxOptSalaryresultQueryRequest() *AlibabaEinvoiceTaxOptSalaryresultQueryAPIRequest{
    return &AlibabaEinvoiceTaxOptSalaryresultQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceTaxOptSalaryresultQueryAPIRequest) GetApiMethodName() string {
    return "alibaba.einvoice.tax.opt.salaryresult.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceTaxOptSalaryresultQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DetailIdList Setter
// 发薪流水号
func (r *AlibabaEinvoiceTaxOptSalaryresultQueryAPIRequest) SetDetailIdList(_detailIdList []string) error {
    r._detailIdList = _detailIdList
    r.Set("detail_id_list", _detailIdList)
    return nil
}

// DetailIdList Getter
func (r AlibabaEinvoiceTaxOptSalaryresultQueryAPIRequest) GetDetailIdList() []string {
    return r._detailIdList
}
// EmployerCode Setter
// 业务方编码
func (r *AlibabaEinvoiceTaxOptSalaryresultQueryAPIRequest) SetEmployerCode(_employerCode string) error {
    r._employerCode = _employerCode
    r.Set("employer_code", _employerCode)
    return nil
}

// EmployerCode Getter
func (r AlibabaEinvoiceTaxOptSalaryresultQueryAPIRequest) GetEmployerCode() string {
    return r._employerCode
}
