package tmallhk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
AWDC提交溯源信息 API请求
tmall.traceplatform.awdc.info.upload

天猫溯源-AWDC-上传溯源信息
*/
type TmallTraceplatformAwdcInfoUploadRequest struct {
    model.Params
    // 入参traceInfo
    _traceInfo   *AwdcInfo
}

// 初始化TmallTraceplatformAwdcInfoUploadRequest对象
func NewTmallTraceplatformAwdcInfoUploadRequest() *TmallTraceplatformAwdcInfoUploadRequest{
    return &TmallTraceplatformAwdcInfoUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallTraceplatformAwdcInfoUploadRequest) GetApiMethodName() string {
    return "tmall.traceplatform.awdc.info.upload"
}

// IRequest interface 方法, 获取API参数
func (r TmallTraceplatformAwdcInfoUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TraceInfo Setter
// 入参traceInfo
func (r *TmallTraceplatformAwdcInfoUploadRequest) SetTraceInfo(_traceInfo *AwdcInfo) error {
    r._traceInfo = _traceInfo
    r.Set("trace_info", _traceInfo)
    return nil
}

// TraceInfo Getter
func (r TmallTraceplatformAwdcInfoUploadRequest) GetTraceInfo() *AwdcInfo {
    return r._traceInfo
}
