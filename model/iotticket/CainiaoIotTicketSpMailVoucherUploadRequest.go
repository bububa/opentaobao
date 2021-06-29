package iotticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商寄出维修件上传凭证信息 APIRequest
cainiao.iot.ticket.sp.mail.voucher.upload

IoT售后服务商寄出维修件上传凭证信息
*/
type CainiaoIotTicketSpMailVoucherUploadRequest struct {
    model.Params

    // 请求参数
    param   *CommentTicketTopRequest 

}

func NewCainiaoIotTicketSpMailVoucherUploadRequest() *CainiaoIotTicketSpMailVoucherUploadRequest{
    return &CainiaoIotTicketSpMailVoucherUploadRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoIotTicketSpMailVoucherUploadRequest) GetApiMethodName() string {
    return "cainiao.iot.ticket.sp.mail.voucher.upload"
}

func (r CainiaoIotTicketSpMailVoucherUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *CainiaoIotTicketSpMailVoucherUploadRequest) SetParam(param *CommentTicketTopRequest) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r CainiaoIotTicketSpMailVoucherUploadRequest) GetParam() *CommentTicketTopRequest {
    return r.param
}

