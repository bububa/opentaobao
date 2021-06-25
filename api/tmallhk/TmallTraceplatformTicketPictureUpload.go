package tmallhk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallhk"
)

/* 
上传小票图片 
tmall.traceplatform.ticket.picture.upload

uploadPicture
*/
func TmallTraceplatformTicketPictureUpload(clt *core.SDKClient, req *tmallhk.TmallTraceplatformTicketPictureUploadRequest, session string) (*tmallhk.TmallTraceplatformTicketPictureUploadResponse, error) {
    var resp tmallhk.TmallTraceplatformTicketPictureUploadAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
