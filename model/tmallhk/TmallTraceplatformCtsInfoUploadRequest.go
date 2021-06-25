package tmallhk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
CTS提交溯源信息 APIRequest
tmall.traceplatform.cts.info.upload

cts上传溯源信息
*/
type TmallTraceplatformCtsInfoUploadRequest struct {
    model.Params

    // 入参traceInfo
    traceInfo   *CtsInfo 

}

func NewTmallTraceplatformCtsInfoUploadRequest() *TmallTraceplatformCtsInfoUploadRequest{
    return &TmallTraceplatformCtsInfoUploadRequest{
        Params: model.NewParams(),
    }
}

func (r TmallTraceplatformCtsInfoUploadRequest) GetApiMethodName() string {
    return "tmall.traceplatform.cts.info.upload"
}

func (r TmallTraceplatformCtsInfoUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallTraceplatformCtsInfoUploadRequest) SetTraceInfo(traceInfo *CtsInfo) error {
    r.traceInfo = traceInfo
    r.Set("trace_info", traceInfo)
    return nil
}

func (r TmallTraceplatformCtsInfoUploadRequest) GetTraceInfo() *CtsInfo {
    return r.traceInfo
}

