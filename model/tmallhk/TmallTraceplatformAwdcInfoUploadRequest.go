package tmallhk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
AWDC提交溯源信息 APIRequest
tmall.traceplatform.awdc.info.upload

天猫溯源-AWDC-上传溯源信息
*/
type TmallTraceplatformAwdcInfoUploadRequest struct {
    model.Params

    // 入参traceInfo
    traceInfo   *AwdcInfo 

}

func NewTmallTraceplatformAwdcInfoUploadRequest() *TmallTraceplatformAwdcInfoUploadRequest{
    return &TmallTraceplatformAwdcInfoUploadRequest{
        Params: model.NewParams(),
    }
}

func (r TmallTraceplatformAwdcInfoUploadRequest) GetApiMethodName() string {
    return "tmall.traceplatform.awdc.info.upload"
}

func (r TmallTraceplatformAwdcInfoUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallTraceplatformAwdcInfoUploadRequest) SetTraceInfo(traceInfo *AwdcInfo) error {
    r.traceInfo = traceInfo
    r.Set("trace_info", traceInfo)
    return nil
}

func (r TmallTraceplatformAwdcInfoUploadRequest) GetTraceInfo() *AwdcInfo {
    return r.traceInfo
}

