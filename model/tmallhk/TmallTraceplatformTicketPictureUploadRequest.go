package tmallhk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
上传小票图片 API请求
tmall.traceplatform.ticket.picture.upload

uploadPicture
*/
type TmallTraceplatformTicketPictureUploadRequest struct {
    model.Params
    // 子订单号
    _bizOrderId   int64
    // 图片二进制流，只支持jpg/jpeg/png格式
    _file   []*model.File
}

// 初始化TmallTraceplatformTicketPictureUploadRequest对象
func NewTmallTraceplatformTicketPictureUploadRequest() *TmallTraceplatformTicketPictureUploadRequest{
    return &TmallTraceplatformTicketPictureUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallTraceplatformTicketPictureUploadRequest) GetApiMethodName() string {
    return "tmall.traceplatform.ticket.picture.upload"
}

// IRequest interface 方法, 获取API参数
func (r TmallTraceplatformTicketPictureUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BizOrderId Setter
// 子订单号
func (r *TmallTraceplatformTicketPictureUploadRequest) SetBizOrderId(_bizOrderId int64) error {
    r._bizOrderId = _bizOrderId
    r.Set("biz_order_id", _bizOrderId)
    return nil
}

// BizOrderId Getter
func (r TmallTraceplatformTicketPictureUploadRequest) GetBizOrderId() int64 {
    return r._bizOrderId
}
// File Setter
// 图片二进制流，只支持jpg/jpeg/png格式
func (r *TmallTraceplatformTicketPictureUploadRequest) SetFile(_file []*model.File) error {
    r._file = _file
    r.Set("file", _file)
    return nil
}

// File Getter
func (r TmallTraceplatformTicketPictureUploadRequest) GetFile() []*model.File {
    return r._file
}
