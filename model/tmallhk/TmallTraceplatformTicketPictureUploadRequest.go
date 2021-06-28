package tmallhk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
上传小票图片 APIRequest
tmall.traceplatform.ticket.picture.upload

uploadPicture
*/
type TmallTraceplatformTicketPictureUploadRequest struct {
    model.Params

    // 子订单号
    bizOrderId   int64 

    // 图片二进制流，只支持jpg/jpeg/png格式
    file   []*model.File 

}

func NewTmallTraceplatformTicketPictureUploadRequest() *TmallTraceplatformTicketPictureUploadRequest{
    return &TmallTraceplatformTicketPictureUploadRequest{
        Params: model.NewParams(),
    }
}

func (r TmallTraceplatformTicketPictureUploadRequest) GetApiMethodName() string {
    return "tmall.traceplatform.ticket.picture.upload"
}

func (r TmallTraceplatformTicketPictureUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallTraceplatformTicketPictureUploadRequest) SetBizOrderId(bizOrderId int64) error {
    r.bizOrderId = bizOrderId
    r.Set("biz_order_id", bizOrderId)
    return nil
}

func (r TmallTraceplatformTicketPictureUploadRequest) GetBizOrderId() int64 {
    return r.bizOrderId
}

func (r *TmallTraceplatformTicketPictureUploadRequest) SetFile(file []*model.File) error {
    r.file = file
    r.Set("file", file)
    return nil
}

func (r TmallTraceplatformTicketPictureUploadRequest) GetFile() []*model.File {
    return r.file
}

