package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
码信息错误上报 API请求
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

// 初始化AlibabaAlihealthDrugCodeErrorReportRequest对象
func NewAlibabaAlihealthDrugCodeErrorReportRequest() *AlibabaAlihealthDrugCodeErrorReportRequest{
    return &AlibabaAlihealthDrugCodeErrorReportRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugCodeErrorReportRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.code.error.report"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugCodeErrorReportRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Code Setter
// 追溯码
func (r *AlibabaAlihealthDrugCodeErrorReportRequest) SetCode(code string) error {
    r.code = code
    r.Set("code", code)
    return nil
}

// Code Getter
func (r AlibabaAlihealthDrugCodeErrorReportRequest) GetCode() string {
    return r.code
}
// FieldName Setter
// 有问题的字段名称
func (r *AlibabaAlihealthDrugCodeErrorReportRequest) SetFieldName(fieldName string) error {
    r.fieldName = fieldName
    r.Set("field_name", fieldName)
    return nil
}

// FieldName Getter
func (r AlibabaAlihealthDrugCodeErrorReportRequest) GetFieldName() string {
    return r.fieldName
}
// CodeValue Setter
// 通过码获得的问题字段值
func (r *AlibabaAlihealthDrugCodeErrorReportRequest) SetCodeValue(codeValue string) error {
    r.codeValue = codeValue
    r.Set("code_value", codeValue)
    return nil
}

// CodeValue Getter
func (r AlibabaAlihealthDrugCodeErrorReportRequest) GetCodeValue() string {
    return r.codeValue
}
// SourceValue Setter
// 平台获得/期望的问题字段值
func (r *AlibabaAlihealthDrugCodeErrorReportRequest) SetSourceValue(sourceValue string) error {
    r.sourceValue = sourceValue
    r.Set("source_value", sourceValue)
    return nil
}

// SourceValue Getter
func (r AlibabaAlihealthDrugCodeErrorReportRequest) GetSourceValue() string {
    return r.sourceValue
}
// ErrMsg Setter
// 错误信息描述
func (r *AlibabaAlihealthDrugCodeErrorReportRequest) SetErrMsg(errMsg string) error {
    r.errMsg = errMsg
    r.Set("err_msg", errMsg)
    return nil
}

// ErrMsg Getter
func (r AlibabaAlihealthDrugCodeErrorReportRequest) GetErrMsg() string {
    return r.errMsg
}
// Reporter Setter
// 上报人员
func (r *AlibabaAlihealthDrugCodeErrorReportRequest) SetReporter(reporter string) error {
    r.reporter = reporter
    r.Set("reporter", reporter)
    return nil
}

// Reporter Getter
func (r AlibabaAlihealthDrugCodeErrorReportRequest) GetReporter() string {
    return r.reporter
}
// ReporterEmail Setter
// 上报人员邮箱
func (r *AlibabaAlihealthDrugCodeErrorReportRequest) SetReporterEmail(reporterEmail string) error {
    r.reporterEmail = reporterEmail
    r.Set("reporter_email", reporterEmail)
    return nil
}

// ReporterEmail Getter
func (r AlibabaAlihealthDrugCodeErrorReportRequest) GetReporterEmail() string {
    return r.reporterEmail
}
// ReporterMobile Setter
// 上报人员手机号
func (r *AlibabaAlihealthDrugCodeErrorReportRequest) SetReporterMobile(reporterMobile string) error {
    r.reporterMobile = reporterMobile
    r.Set("reporter_mobile", reporterMobile)
    return nil
}

// ReporterMobile Getter
func (r AlibabaAlihealthDrugCodeErrorReportRequest) GetReporterMobile() string {
    return r.reporterMobile
}
