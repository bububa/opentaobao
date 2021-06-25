package tmallhk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
ccic校验溯源码 APIRequest
tmall.traceplatform.ccic.tracecode.check

天猫国际溯源业务，需要将溯源码校验的功能输出到ccic官方主页中以增强溯源码的可信度，故需要提供api给ccic使用以校验溯源码的正确性。
*/
type TmallTraceplatformCcicTracecodeCheckRequest struct {
    model.Params

    // 15为溯源短码，必选
    shortTracecode   string 

    // 6位暗码，必选
    hideCode   string 

}

func NewTmallTraceplatformCcicTracecodeCheckRequest() *TmallTraceplatformCcicTracecodeCheckRequest{
    return &TmallTraceplatformCcicTracecodeCheckRequest{
        Params: model.NewParams(),
    }
}

func (r TmallTraceplatformCcicTracecodeCheckRequest) GetApiMethodName() string {
    return "tmall.traceplatform.ccic.tracecode.check"
}

func (r TmallTraceplatformCcicTracecodeCheckRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallTraceplatformCcicTracecodeCheckRequest) SetShortTracecode(shortTracecode string) error {
    r.shortTracecode = shortTracecode
    r.Set("short_tracecode", shortTracecode)
    return nil
}

func (r TmallTraceplatformCcicTracecodeCheckRequest) GetShortTracecode() string {
    return r.shortTracecode
}

func (r *TmallTraceplatformCcicTracecodeCheckRequest) SetHideCode(hideCode string) error {
    r.hideCode = hideCode
    r.Set("hide_code", hideCode)
    return nil
}

func (r TmallTraceplatformCcicTracecodeCheckRequest) GetHideCode() string {
    return r.hideCode
}

