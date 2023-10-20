package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugcodeerrorreportAPIRequest 码信息错误上报 API请求
// alibaba.alihealth.drug.code.error.report
//
// 提供码信息错误上报功能，用于数据校对
type AlibabaalihealthdrugcodeerrorreportAPIRequest struct {
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

// NewAlibabaalihealthdrugcodeerrorreportRequest 初始化AlibabaalihealthdrugcodeerrorreportAPIRequest对象
func NewAlibabaalihealthdrugcodeerrorreportRequest() *AlibabaalihealthdrugcodeerrorreportAPIRequest {
	return &AlibabaalihealthdrugcodeerrorreportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugcodeerrorreportAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.code.error.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugcodeerrorreportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugcodeerrorreportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCode is Code Setter
// 追溯码
func (r *AlibabaalihealthdrugcodeerrorreportAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabaalihealthdrugcodeerrorreportAPIRequest) GetCode() string {
	return r._code
}

// SetFieldName is FieldName Setter
// 有问题的字段名称
func (r *AlibabaalihealthdrugcodeerrorreportAPIRequest) SetFieldName(_fieldName string) error {
	r._fieldName = _fieldName
	r.Set("field_name", _fieldName)
	return nil
}

// GetFieldName FieldName Getter
func (r AlibabaalihealthdrugcodeerrorreportAPIRequest) GetFieldName() string {
	return r._fieldName
}

// SetCodeValue is CodeValue Setter
// 通过码获得的问题字段值
func (r *AlibabaalihealthdrugcodeerrorreportAPIRequest) SetCodeValue(_codeValue string) error {
	r._codeValue = _codeValue
	r.Set("code_value", _codeValue)
	return nil
}

// GetCodeValue CodeValue Getter
func (r AlibabaalihealthdrugcodeerrorreportAPIRequest) GetCodeValue() string {
	return r._codeValue
}

// SetSourceValue is SourceValue Setter
// 平台获得/期望的问题字段值
func (r *AlibabaalihealthdrugcodeerrorreportAPIRequest) SetSourceValue(_sourceValue string) error {
	r._sourceValue = _sourceValue
	r.Set("source_value", _sourceValue)
	return nil
}

// GetSourceValue SourceValue Getter
func (r AlibabaalihealthdrugcodeerrorreportAPIRequest) GetSourceValue() string {
	return r._sourceValue
}

// SetErrMsg is ErrMsg Setter
// 错误信息描述
func (r *AlibabaalihealthdrugcodeerrorreportAPIRequest) SetErrMsg(_errMsg string) error {
	r._errMsg = _errMsg
	r.Set("err_msg", _errMsg)
	return nil
}

// GetErrMsg ErrMsg Getter
func (r AlibabaalihealthdrugcodeerrorreportAPIRequest) GetErrMsg() string {
	return r._errMsg
}

// SetReporter is Reporter Setter
// 上报人员
func (r *AlibabaalihealthdrugcodeerrorreportAPIRequest) SetReporter(_reporter string) error {
	r._reporter = _reporter
	r.Set("reporter", _reporter)
	return nil
}

// GetReporter Reporter Getter
func (r AlibabaalihealthdrugcodeerrorreportAPIRequest) GetReporter() string {
	return r._reporter
}

// SetReporterEmail is ReporterEmail Setter
// 上报人员邮箱
func (r *AlibabaalihealthdrugcodeerrorreportAPIRequest) SetReporterEmail(_reporterEmail string) error {
	r._reporterEmail = _reporterEmail
	r.Set("reporter_email", _reporterEmail)
	return nil
}

// GetReporterEmail ReporterEmail Getter
func (r AlibabaalihealthdrugcodeerrorreportAPIRequest) GetReporterEmail() string {
	return r._reporterEmail
}

// SetReporterMobile is ReporterMobile Setter
// 上报人员手机号
func (r *AlibabaalihealthdrugcodeerrorreportAPIRequest) SetReporterMobile(_reporterMobile string) error {
	r._reporterMobile = _reporterMobile
	r.Set("reporter_mobile", _reporterMobile)
	return nil
}

// GetReporterMobile ReporterMobile Getter
func (r AlibabaalihealthdrugcodeerrorreportAPIRequest) GetReporterMobile() string {
	return r._reporterMobile
}
