package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugCodeErrorReportAPIRequest
码信息错误上报 API请求
alibaba.alihealth.drug.code.error.report

提供码信息错误上报功能，用于数据校对 */
type AlibabaAlihealthDrugCodeErrorReportAPIRequest struct {
	model.Params
	// 追溯码
	_code string
	// 有问题的字段名称
	_fieldName string
	// 通过码获得的问题字段值
	_codeValue string
	// 平台获得/期望的问题字段值
	_sourceValue string
	// 错误信息描述
	_errMsg string
	// 上报人员
	_reporter string
	// 上报人员邮箱
	_reporterEmail string
	// 上报人员手机号
	_reporterMobile string
}

// NewAlibabaAlihealthDrugCodeErrorReportRequest 初始化AlibabaAlihealthDrugCodeErrorReportAPIRequest对象
func NewAlibabaAlihealthDrugCodeErrorReportRequest() *AlibabaAlihealthDrugCodeErrorReportAPIRequest {
	return &AlibabaAlihealthDrugCodeErrorReportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugCodeErrorReportAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.code.error.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugCodeErrorReportAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Code Setter
// 追溯码
func (r *AlibabaAlihealthDrugCodeErrorReportAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// Get Code Getter
func (r AlibabaAlihealthDrugCodeErrorReportAPIRequest) GetCode() string {
	return r._code
}

// Set is FieldName Setter
// 有问题的字段名称
func (r *AlibabaAlihealthDrugCodeErrorReportAPIRequest) SetFieldName(_fieldName string) error {
	r._fieldName = _fieldName
	r.Set("field_name", _fieldName)
	return nil
}

// Get FieldName Getter
func (r AlibabaAlihealthDrugCodeErrorReportAPIRequest) GetFieldName() string {
	return r._fieldName
}

// Set is CodeValue Setter
// 通过码获得的问题字段值
func (r *AlibabaAlihealthDrugCodeErrorReportAPIRequest) SetCodeValue(_codeValue string) error {
	r._codeValue = _codeValue
	r.Set("code_value", _codeValue)
	return nil
}

// Get CodeValue Getter
func (r AlibabaAlihealthDrugCodeErrorReportAPIRequest) GetCodeValue() string {
	return r._codeValue
}

// Set is SourceValue Setter
// 平台获得/期望的问题字段值
func (r *AlibabaAlihealthDrugCodeErrorReportAPIRequest) SetSourceValue(_sourceValue string) error {
	r._sourceValue = _sourceValue
	r.Set("source_value", _sourceValue)
	return nil
}

// Get SourceValue Getter
func (r AlibabaAlihealthDrugCodeErrorReportAPIRequest) GetSourceValue() string {
	return r._sourceValue
}

// Set is ErrMsg Setter
// 错误信息描述
func (r *AlibabaAlihealthDrugCodeErrorReportAPIRequest) SetErrMsg(_errMsg string) error {
	r._errMsg = _errMsg
	r.Set("err_msg", _errMsg)
	return nil
}

// Get ErrMsg Getter
func (r AlibabaAlihealthDrugCodeErrorReportAPIRequest) GetErrMsg() string {
	return r._errMsg
}

// Set is Reporter Setter
// 上报人员
func (r *AlibabaAlihealthDrugCodeErrorReportAPIRequest) SetReporter(_reporter string) error {
	r._reporter = _reporter
	r.Set("reporter", _reporter)
	return nil
}

// Get Reporter Getter
func (r AlibabaAlihealthDrugCodeErrorReportAPIRequest) GetReporter() string {
	return r._reporter
}

// Set is ReporterEmail Setter
// 上报人员邮箱
func (r *AlibabaAlihealthDrugCodeErrorReportAPIRequest) SetReporterEmail(_reporterEmail string) error {
	r._reporterEmail = _reporterEmail
	r.Set("reporter_email", _reporterEmail)
	return nil
}

// Get ReporterEmail Getter
func (r AlibabaAlihealthDrugCodeErrorReportAPIRequest) GetReporterEmail() string {
	return r._reporterEmail
}

// Set is ReporterMobile Setter
// 上报人员手机号
func (r *AlibabaAlihealthDrugCodeErrorReportAPIRequest) SetReporterMobile(_reporterMobile string) error {
	r._reporterMobile = _reporterMobile
	r.Set("reporter_mobile", _reporterMobile)
	return nil
}

// Get ReporterMobile Getter
func (r AlibabaAlihealthDrugCodeErrorReportAPIRequest) GetReporterMobile() string {
	return r._reporterMobile
}
