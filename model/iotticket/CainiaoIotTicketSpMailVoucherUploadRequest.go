package iotticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商寄出维修件上传凭证信息 API请求
cainiao.iot.ticket.sp.mail.voucher.upload

IoT售后服务商寄出维修件上传凭证信息
*/
type CainiaoIotTicketSpMailVoucherUploadRequest struct {
    model.Params
    // 请求参数
    _param   *CommentTicketTopRequest
}

// 初始化CainiaoIotTicketSpMailVoucherUploadRequest对象
func NewCainiaoIotTicketSpMailVoucherUploadRequest() *CainiaoIotTicketSpMailVoucherUploadRequest{
    return &CainiaoIotTicketSpMailVoucherUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoIotTicketSpMailVoucherUploadRequest) GetApiMethodName() string {
    return "cainiao.iot.ticket.sp.mail.voucher.upload"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoIotTicketSpMailVoucherUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 请求参数
func (r *CainiaoIotTicketSpMailVoucherUploadRequest) SetParam(_param *CommentTicketTopRequest) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r CainiaoIotTicketSpMailVoucherUploadRequest) GetParam() *CommentTicketTopRequest {
    return r._param
}
