package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugCodeErrorReportAPIRequest 码信息错误上报 API请求
// alibaba.alihealth.drug.code.error.report
//
// 提供码信息错误上报功能，用于数据校对
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
		Params: model.NewParams(8),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugCodeErrorReportAPIRequest) Reset() {
	r._code = ""
	r._fieldName = ""
	r._codeValue = ""
	r._sourceValue = ""
	r._errMsg = ""
	r._reporter = ""
	r._reporterEmail = ""
	r._reporterMobile = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugCodeErrorReportAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.code.error.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugCodeErrorReportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugCodeErrorReportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCode is Code Setter
// 追溯码
func (r *AlibabaAlihealthDrugCodeErrorReportAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabaAlihealthDrugCodeErrorReportAPIRequest) GetCode() string {
	return r._code
}

// SetFieldName is FieldName Setter
// 有问题的字段名称
func (r *AlibabaAlihealthDrugCodeErrorReportAPIRequest) SetFieldName(_fieldName string) error {
	r._fieldName = _fieldName
	r.Set("field_name", _fieldName)
	return nil
}

// GetFieldName FieldName Getter
func (r AlibabaAlihealthDrugCodeErrorReportAPIRequest) GetFieldName() string {
	return r._fieldName
}

// SetCodeValue is CodeValue Setter
// 通过码获得的问题字段值
func (r *AlibabaAlihealthDrugCodeErrorReportAPIRequest) SetCodeValue(_codeValue string) error {
	r._codeValue = _codeValue
	r.Set("code_value", _codeValue)
	return nil
}

// GetCodeValue CodeValue Getter
func (r AlibabaAlihealthDrugCodeErrorReportAPIRequest) GetCodeValue() string {
	return r._codeValue
}

// SetSourceValue is SourceValue Setter
// 平台获得/期望的问题字段值
func (r *AlibabaAlihealthDrugCodeErrorReportAPIRequest) SetSourceValue(_sourceValue string) error {
	r._sourceValue = _sourceValue
	r.Set("source_value", _sourceValue)
	return nil
}

// GetSourceValue SourceValue Getter
func (r AlibabaAlihealthDrugCodeErrorReportAPIRequest) GetSourceValue() string {
	return r._sourceValue
}

// SetErrMsg is ErrMsg Setter
// 错误信息描述
func (r *AlibabaAlihealthDrugCodeErrorReportAPIRequest) SetErrMsg(_errMsg string) error {
	r._errMsg = _errMsg
	r.Set("err_msg", _errMsg)
	return nil
}

// GetErrMsg ErrMsg Getter
func (r AlibabaAlihealthDrugCodeErrorReportAPIRequest) GetErrMsg() string {
	return r._errMsg
}

// SetReporter is Reporter Setter
// 上报人员
func (r *AlibabaAlihealthDrugCodeErrorReportAPIRequest) SetReporter(_reporter string) error {
	r._reporter = _reporter
	r.Set("reporter", _reporter)
	return nil
}

// GetReporter Reporter Getter
func (r AlibabaAlihealthDrugCodeErrorReportAPIRequest) GetReporter() string {
	return r._reporter
}

// SetReporterEmail is ReporterEmail Setter
// 上报人员邮箱
func (r *AlibabaAlihealthDrugCodeErrorReportAPIRequest) SetReporterEmail(_reporterEmail string) error {
	r._reporterEmail = _reporterEmail
	r.Set("reporter_email", _reporterEmail)
	return nil
}

// GetReporterEmail ReporterEmail Getter
func (r AlibabaAlihealthDrugCodeErrorReportAPIRequest) GetReporterEmail() string {
	return r._reporterEmail
}

// SetReporterMobile is ReporterMobile Setter
// 上报人员手机号
func (r *AlibabaAlihealthDrugCodeErrorReportAPIRequest) SetReporterMobile(_reporterMobile string) error {
	r._reporterMobile = _reporterMobile
	r.Set("reporter_mobile", _reporterMobile)
	return nil
}

// GetReporterMobile ReporterMobile Getter
func (r AlibabaAlihealthDrugCodeErrorReportAPIRequest) GetReporterMobile() string {
	return r._reporterMobile
}

var poolAlibabaAlihealthDrugCodeErrorReportAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugCodeErrorReportRequest()
	},
}

// GetAlibabaAlihealthDrugCodeErrorReportRequest 从 sync.Pool 获取 AlibabaAlihealthDrugCodeErrorReportAPIRequest
func GetAlibabaAlihealthDrugCodeErrorReportAPIRequest() *AlibabaAlihealthDrugCodeErrorReportAPIRequest {
	return poolAlibabaAlihealthDrugCodeErrorReportAPIRequest.Get().(*AlibabaAlihealthDrugCodeErrorReportAPIRequest)
}

// ReleaseAlibabaAlihealthDrugCodeErrorReportAPIRequest 将 AlibabaAlihealthDrugCodeErrorReportAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugCodeErrorReportAPIRequest(v *AlibabaAlihealthDrugCodeErrorReportAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugCodeErrorReportAPIRequest.Put(v)
}
