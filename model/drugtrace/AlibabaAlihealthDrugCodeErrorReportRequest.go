package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
码信息错误上报 APIRequest
alibaba.alihealth.drug.code.error.report

提供码信息错误上报功能，用于数据校对
*/
type AlibabaAlihealthDrugCodeErrorReportRequest struct {
    model.Params

    // 追溯码
    code   string 

    // 有问题的字段名称
    fieldName   string 

    // 通过码获得的问题字段值
    codeValue   string 

    // 平台获得/期望的问题字段值
    sourceValue   string 

    // 错误信息描述
    errMsg   string 

    // 上报人员
    reporter   string 

    // 上报人员邮箱
    reporterEmail   string 

    // 上报人员手机号
    reporterMobile   string 

}

func NewAlibabaAlihealthDrugCodeErrorReportRequest() *AlibabaAlihealthDrugCodeErrorReportRequest{
    return &AlibabaAlihealthDrugCodeErrorReportRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugCodeErrorReportRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.code.error.report"
}

func (r AlibabaAlihealthDrugCodeErrorReportRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugCodeErrorReportRequest) SetCode(code string) error {
    r.code = code
    r.Set("code", code)
    return nil
}

func (r AlibabaAlihealthDrugCodeErrorReportRequest) GetCode() string {
    return r.code
}

func (r *AlibabaAlihealthDrugCodeErrorReportRequest) SetFieldName(fieldName string) error {
    r.fieldName = fieldName
    r.Set("field_name", fieldName)
    return nil
}

func (r AlibabaAlihealthDrugCodeErrorReportRequest) GetFieldName() string {
    return r.fieldName
}

func (r *AlibabaAlihealthDrugCodeErrorReportRequest) SetCodeValue(codeValue string) error {
    r.codeValue = codeValue
    r.Set("code_value", codeValue)
    return nil
}

func (r AlibabaAlihealthDrugCodeErrorReportRequest) GetCodeValue() string {
    return r.codeValue
}

func (r *AlibabaAlihealthDrugCodeErrorReportRequest) SetSourceValue(sourceValue string) error {
    r.sourceValue = sourceValue
    r.Set("source_value", sourceValue)
    return nil
}

func (r AlibabaAlihealthDrugCodeErrorReportRequest) GetSourceValue() string {
    return r.sourceValue
}

func (r *AlibabaAlihealthDrugCodeErrorReportRequest) SetErrMsg(errMsg string) error {
    r.errMsg = errMsg
    r.Set("err_msg", errMsg)
    return nil
}

func (r AlibabaAlihealthDrugCodeErrorReportRequest) GetErrMsg() string {
    return r.errMsg
}

func (r *AlibabaAlihealthDrugCodeErrorReportRequest) SetReporter(reporter string) error {
    r.reporter = reporter
    r.Set("reporter", reporter)
    return nil
}

func (r AlibabaAlihealthDrugCodeErrorReportRequest) GetReporter() string {
    return r.reporter
}

func (r *AlibabaAlihealthDrugCodeErrorReportRequest) SetReporterEmail(reporterEmail string) error {
    r.reporterEmail = reporterEmail
    r.Set("reporter_email", reporterEmail)
    return nil
}

func (r AlibabaAlihealthDrugCodeErrorReportRequest) GetReporterEmail() string {
    return r.reporterEmail
}

func (r *AlibabaAlihealthDrugCodeErrorReportRequest) SetReporterMobile(reporterMobile string) error {
    r.reporterMobile = reporterMobile
    r.Set("reporter_mobile", reporterMobile)
    return nil
}

func (r AlibabaAlihealthDrugCodeErrorReportRequest) GetReporterMobile() string {
    return r.reporterMobile
}

