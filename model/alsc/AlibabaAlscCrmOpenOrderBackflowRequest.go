package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
订单回流接口 API请求
alibaba.alsc.crm.open.order.backflow

回流isv订单接口
*/
type AlibabaAlscCrmOpenOrderBackflowRequest struct {
    model.Params
    // 入参
    paramOrderBackflowOpenReq   *OrderBackflowOpenReq
}

// 初始化AlibabaAlscCrmOpenOrderBackflowRequest对象
func NewAlibabaAlscCrmOpenOrderBackflowRequest() *AlibabaAlscCrmOpenOrderBackflowRequest{
    return &AlibabaAlscCrmOpenOrderBackflowRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmOpenOrderBackflowRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.open.order.backflow"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmOpenOrderBackflowRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamOrderBackflowOpenReq Setter
// 入参
func (r *AlibabaAlscCrmOpenOrderBackflowRequest) SetParamOrderBackflowOpenReq(paramOrderBackflowOpenReq *OrderBackflowOpenReq) error {
    r.paramOrderBackflowOpenReq = paramOrderBackflowOpenReq
    r.Set("param_order_backflow_open_req", paramOrderBackflowOpenReq)
    return nil
}

// ParamOrderBackflowOpenReq Getter
func (r AlibabaAlscCrmOpenOrderBackflowRequest) GetParamOrderBackflowOpenReq() *OrderBackflowOpenReq {
    return r.paramOrderBackflowOpenReq
}
