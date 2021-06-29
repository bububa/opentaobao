package iotticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
IoT售后服务商工单备注 APIRequest
cainiao.iot.ticket.sp.comment

IoT售后服务商工单备注
*/
type CainiaoIotTicketSpCommentRequest struct {
    model.Params

    // 请求参数
    param   *CommentTicketTopRequest 

}

func NewCainiaoIotTicketSpCommentRequest() *CainiaoIotTicketSpCommentRequest{
    return &CainiaoIotTicketSpCommentRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoIotTicketSpCommentRequest) GetApiMethodName() string {
    return "cainiao.iot.ticket.sp.comment"
}

func (r CainiaoIotTicketSpCommentRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *CainiaoIotTicketSpCommentRequest) SetParam(param *CommentTicketTopRequest) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r CainiaoIotTicketSpCommentRequest) GetParam() *CommentTicketTopRequest {
    return r.param
}

