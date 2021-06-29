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
    _code   string
    // 有问题的字段名称
    _fieldName   string
    // 通过码获得的问题字段值
    _codeValue   string
    // 平台获得/期望的问题字段值
    _sourceValue   string
    // 错误信息描述
    _errMsg   string
    // 上报人员
    _reporter   string
    // 上报人员邮箱
    _reporterEmail   string
    // 上报人员手机号
    _reporterMobile   string
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
func (r *AlibabaAlihealthDrugCodeErrorReportRequest) SetCode(_code string) error {
    r._code = _code
    r.Set("code", _code)
    return nil
}

// Code Getter
func (r AlibabaAlihealthDrugCodeErrorReportRequest) GetCode() string {
    return r._code
}
// FieldName Setter
// 有问题的字段名称
func (r *AlibabaAlihealthDrugCodeErrorReportRequest) SetFieldName(_fieldName string) error {
    r._fieldName = _fieldName
    r.Set("field_name", _fieldName)
    return nil
}

// FieldName Getter
func (r AlibabaAlihealthDrugCodeErrorReportRequest) GetFieldName() string {
    return r._fieldName
}
// CodeValue Setter
// 通过码获得的问题字段值
func (r *AlibabaAlihealthDrugCodeErrorReportRequest) SetCodeValue(_codeValue string) error {
    r._codeValue = _codeValue
    r.Set("code_value", _codeValue)
    return nil
}

// CodeValue Getter
func (r AlibabaAlihealthDrugCodeErrorReportRequest) GetCodeValue() string {
    return r._codeValue
}
// SourceValue Setter
// 平台获得/期望的问题字段值
func (r *AlibabaAlihealthDrugCodeErrorReportRequest) SetSourceValue(_sourceValue string) error {
    r._sourceValue = _sourceValue
    r.Set("source_value", _sourceValue)
    return nil
}

// SourceValue Getter
func (r AlibabaAlihealthDrugCodeErrorReportRequest) GetSourceValue() string {
    return r._sourceValue
}
// ErrMsg Setter
// 错误信息描述
func (r *AlibabaAlihealthDrugCodeErrorReportRequest) SetErrMsg(_errMsg string) error {
    r._errMsg = _errMsg
    r.Set("err_msg", _errMsg)
    return nil
}

// ErrMsg Getter
func (r AlibabaAlihealthDrugCodeErrorReportRequest) GetErrMsg() string {
    return r._errMsg
}
// Reporter Setter
// 上报人员
func (r *AlibabaAlihealthDrugCodeErrorReportRequest) SetReporter(_reporter string) error {
    r._reporter = _reporter
    r.Set("reporter", _reporter)
    return nil
}

// Reporter Getter
func (r AlibabaAlihealthDrugCodeErrorReportRequest) GetReporter() string {
    return r._reporter
}
// ReporterEmail Setter
// 上报人员邮箱
func (r *AlibabaAlihealthDrugCodeErrorReportRequest) SetReporterEmail(_reporterEmail string) error {
    r._reporterEmail = _reporterEmail
    r.Set("reporter_email", _reporterEmail)
    return nil
}

// ReporterEmail Getter
func (r AlibabaAlihealthDrugCodeErrorReportRequest) GetReporterEmail() string {
    return r._reporterEmail
}
// ReporterMobile Setter
// 上报人员手机号
func (r *AlibabaAlihealthDrugCodeErrorReportRequest) SetReporterMobile(_reporterMobile string) error {
    r._reporterMobile = _reporterMobile
    r.Set("reporter_mobile", _reporterMobile)
    return nil
}

// ReporterMobile Getter
func (r AlibabaAlihealthDrugCodeErrorReportRequest) GetReporterMobile() string {
    return r._reporterMobile
}
