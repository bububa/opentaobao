package tmallhk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
CTS提交溯源信息 API请求
tmall.traceplatform.cts.info.upload

cts上传溯源信息
*/
type TmallTraceplatformCtsInfoUploadRequest struct {
    model.Params
    // 入参traceInfo
    _traceInfo   *CtsInfo
}

// 初始化TmallTraceplatformCtsInfoUploadRequest对象
func NewTmallTraceplatformCtsInfoUploadRequest() *TmallTraceplatformCtsInfoUploadRequest{
    return &TmallTraceplatformCtsInfoUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallTraceplatformCtsInfoUploadRequest) GetApiMethodName() string {
    return "tmall.traceplatform.cts.info.upload"
}

// IRequest interface 方法, 获取API参数
func (r TmallTraceplatformCtsInfoUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TraceInfo Setter
// 入参traceInfo
func (r *TmallTraceplatformCtsInfoUploadRequest) SetTraceInfo(_traceInfo *CtsInfo) error {
    r._traceInfo = _traceInfo
    r.Set("trace_info", _traceInfo)
    return nil
}

// TraceInfo Getter
func (r TmallTraceplatformCtsInfoUploadRequest) GetTraceInfo() *CtsInfo {
    return r._traceInfo
}
