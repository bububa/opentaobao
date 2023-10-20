package tmallhk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallTraceplatformCcicTracecodeCheckAPIRequest ccic校验溯源码 API请求
// tmall.traceplatform.ccic.tracecode.check
//
// 天猫国际溯源业务，需要将溯源码校验的功能输出到ccic官方主页中以增强溯源码的可信度，故需要提供api给ccic使用以校验溯源码的正确性。
type TmallTraceplatformCcicTracecodeCheckAPIRequest struct {
	model.Params
	// 15为溯源短码，必选
	_shortTracecode string
	// 6位暗码，必选
	_hideCode string
}

// NewTmallTraceplatformCcicTracecodeCheckRequest 初始化TmallTraceplatformCcicTracecodeCheckAPIRequest对象
func NewTmallTraceplatformCcicTracecodeCheckRequest() *TmallTraceplatformCcicTracecodeCheckAPIRequest {
	return &TmallTraceplatformCcicTracecodeCheckAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallTraceplatformCcicTracecodeCheckAPIRequest) GetApiMethodName() string {
	return "tmall.traceplatform.ccic.tracecode.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallTraceplatformCcicTracecodeCheckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallTraceplatformCcicTracecodeCheckAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetShortTracecode is ShortTracecode Setter
// 15为溯源短码，必选
func (r *TmallTraceplatformCcicTracecodeCheckAPIRequest) SetShortTracecode(_shortTracecode string) error {
	r._shortTracecode = _shortTracecode
	r.Set("short_tracecode", _shortTracecode)
	return nil
}

// GetShortTracecode ShortTracecode Getter
func (r TmallTraceplatformCcicTracecodeCheckAPIRequest) GetShortTracecode() string {
	return r._shortTracecode
}

// SetHideCode is HideCode Setter
// 6位暗码，必选
func (r *TmallTraceplatformCcicTracecodeCheckAPIRequest) SetHideCode(_hideCode string) error {
	r._hideCode = _hideCode
	r.Set("hide_code", _hideCode)
	return nil
}

// GetHideCode HideCode Getter
func (r TmallTraceplatformCcicTracecodeCheckAPIRequest) GetHideCode() string {
	return r._hideCode
}
