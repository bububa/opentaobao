package tmallhk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallhk"
)

/* 
CTS提交溯源信息 
tmall.traceplatform.cts.info.upload

cts上传溯源信息
*/
func TmallTraceplatformCtsInfoUpload(clt *core.SDKClient, req *tmallhk.TmallTraceplatformCtsInfoUploadRequest, session string) (*tmallhk.TmallTraceplatformCtsInfoUploadAPIResponse, error) {
    var resp tmallhk.TmallTraceplatformCtsInfoUploadAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
