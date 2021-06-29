package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询用户签约税优结果 API请求
alibaba.einvoice.tax.opt.esignresult.query

查询用户是否已经签约
*/
type AlibabaEinvoiceTaxOptEsignresultQueryRequest struct {
    model.Params
    // 业务方编码
    employerCode   string
    // 用户在业务方平台的userid
    identificationInBelongingEmployer   string
}

// 初始化AlibabaEinvoiceTaxOptEsignresultQueryRequest对象
func NewAlibabaEinvoiceTaxOptEsignresultQueryRequest() *AlibabaEinvoiceTaxOptEsignresultQueryRequest{
    return &AlibabaEinvoiceTaxOptEsignresultQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceTaxOptEsignresultQueryRequest) GetApiMethodName() string {
    return "alibaba.einvoice.tax.opt.esignresult.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceTaxOptEsignresultQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// EmployerCode Setter
// 业务方编码
func (r *AlibabaEinvoiceTaxOptEsignresultQueryRequest) SetEmployerCode(employerCode string) error {
    r.employerCode = employerCode
    r.Set("employer_code", employerCode)
    return nil
}

// EmployerCode Getter
func (r AlibabaEinvoiceTaxOptEsignresultQueryRequest) GetEmployerCode() string {
    return r.employerCode
}
// IdentificationInBelongingEmployer Setter
// 用户在业务方平台的userid
func (r *AlibabaEinvoiceTaxOptEsignresultQueryRequest) SetIdentificationInBelongingEmployer(identificationInBelongingEmployer string) error {
    r.identificationInBelongingEmployer = identificationInBelongingEmployer
    r.Set("identification_in_belonging_employer", identificationInBelongingEmployer)
    return nil
}

// IdentificationInBelongingEmployer Getter
func (r AlibabaEinvoiceTaxOptEsignresultQueryRequest) GetIdentificationInBelongingEmployer() string {
    return r.identificationInBelongingEmployer
}
