package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdruggetbarcodebytraccodeAPIRequest 根据追溯码获取69码 API请求
// alibaba.alihealth.drug.getbarcode.bytraccode
//
// 根据追溯码获取69码
type AlibabaalihealthdruggetbarcodebytraccodeAPIRequest struct {
	model.Params
	// 追溯码
	_traceCode string
}

// NewAlibabaalihealthdruggetbarcodebytraccodeRequest 初始化AlibabaalihealthdruggetbarcodebytraccodeAPIRequest对象
func NewAlibabaalihealthdruggetbarcodebytraccodeRequest() *AlibabaalihealthdruggetbarcodebytraccodeAPIRequest {
	return &AlibabaalihealthdruggetbarcodebytraccodeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdruggetbarcodebytraccodeAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.getbarcode.bytraccode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdruggetbarcodebytraccodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdruggetbarcodebytraccodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTraceCode is TraceCode Setter
// 追溯码
func (r *AlibabaalihealthdruggetbarcodebytraccodeAPIRequest) SetTraceCode(_traceCode string) error {
	r._traceCode = _traceCode
	r.Set("trace_code", _traceCode)
	return nil
}

// GetTraceCode TraceCode Getter
func (r AlibabaalihealthdruggetbarcodebytraccodeAPIRequest) GetTraceCode() string {
	return r._traceCode
}
