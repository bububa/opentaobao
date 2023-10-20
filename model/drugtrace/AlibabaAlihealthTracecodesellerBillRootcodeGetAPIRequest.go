package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthtracecodesellerbillrootcodegetAPIRequest 获取最外层包装码 API请求
// alibaba.alihealth.tracecodeseller.bill.rootcode.get
//
// 获取最外层包装码
type AlibabaalihealthtracecodesellerbillrootcodegetAPIRequest struct {
	model.Params
	// 用户身份认证
	_appCode string
	// 码
	_code string
}

// NewAlibabaalihealthtracecodesellerbillrootcodegetRequest 初始化AlibabaalihealthtracecodesellerbillrootcodegetAPIRequest对象
func NewAlibabaalihealthtracecodesellerbillrootcodegetRequest() *AlibabaalihealthtracecodesellerbillrootcodegetAPIRequest {
	return &AlibabaalihealthtracecodesellerbillrootcodegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthtracecodesellerbillrootcodegetAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.tracecodeseller.bill.rootcode.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthtracecodesellerbillrootcodegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthtracecodesellerbillrootcodegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAppCode is AppCode Setter
// 用户身份认证
func (r *AlibabaalihealthtracecodesellerbillrootcodegetAPIRequest) SetAppCode(_appCode string) error {
	r._appCode = _appCode
	r.Set("app_code", _appCode)
	return nil
}

// GetAppCode AppCode Getter
func (r AlibabaalihealthtracecodesellerbillrootcodegetAPIRequest) GetAppCode() string {
	return r._appCode
}

// SetCode is Code Setter
// 码
func (r *AlibabaalihealthtracecodesellerbillrootcodegetAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabaalihealthtracecodesellerbillrootcodegetAPIRequest) GetCode() string {
	return r._code
}
