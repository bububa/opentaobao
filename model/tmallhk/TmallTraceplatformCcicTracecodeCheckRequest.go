package tmallhk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
ccic校验溯源码 API请求
tmall.traceplatform.ccic.tracecode.check

天猫国际溯源业务，需要将溯源码校验的功能输出到ccic官方主页中以增强溯源码的可信度，故需要提供api给ccic使用以校验溯源码的正确性。
*/
type TmallTraceplatformCcicTracecodeCheckRequest struct {
    model.Params
    // 15为溯源短码，必选
    _shortTracecode   string
    // 6位暗码，必选
    _hideCode   string
}

// 初始化TmallTraceplatformCcicTracecodeCheckRequest对象
func NewTmallTraceplatformCcicTracecodeCheckRequest() *TmallTraceplatformCcicTracecodeCheckRequest{
    return &TmallTraceplatformCcicTracecodeCheckRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallTraceplatformCcicTracecodeCheckRequest) GetApiMethodName() string {
    return "tmall.traceplatform.ccic.tracecode.check"
}

// IRequest interface 方法, 获取API参数
func (r TmallTraceplatformCcicTracecodeCheckRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ShortTracecode Setter
// 15为溯源短码，必选
func (r *TmallTraceplatformCcicTracecodeCheckRequest) SetShortTracecode(_shortTracecode string) error {
    r._shortTracecode = _shortTracecode
    r.Set("short_tracecode", _shortTracecode)
    return nil
}

// ShortTracecode Getter
func (r TmallTraceplatformCcicTracecodeCheckRequest) GetShortTracecode() string {
    return r._shortTracecode
}
// HideCode Setter
// 6位暗码，必选
func (r *TmallTraceplatformCcicTracecodeCheckRequest) SetHideCode(_hideCode string) error {
    r._hideCode = _hideCode
    r.Set("hide_code", _hideCode)
    return nil
}

// HideCode Getter
func (r TmallTraceplatformCcicTracecodeCheckRequest) GetHideCode() string {
    return r._hideCode
}
