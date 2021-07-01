package iotticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
IoT售后服务商工单备注 API请求
cainiao.iot.ticket.sp.comment

IoT售后服务商工单备注
*/
type CainiaoIotTicketSpCommentAPIRequest struct {
    model.Params
    // 请求参数
    _param   *CommentTicketTopRequest
}

// 初始化CainiaoIotTicketSpCommentAPIRequest对象
func NewCainiaoIotTicketSpCommentRequest() *CainiaoIotTicketSpCommentAPIRequest{
    return &CainiaoIotTicketSpCommentAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoIotTicketSpCommentAPIRequest) GetApiMethodName() string {
    return "cainiao.iot.ticket.sp.comment"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoIotTicketSpCommentAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 请求参数
func (r *CainiaoIotTicketSpCommentAPIRequest) SetParam(_param *CommentTicketTopRequest) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r CainiaoIotTicketSpCommentAPIRequest) GetParam() *CommentTicketTopRequest {
    return r._param
}
