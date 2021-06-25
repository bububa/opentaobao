package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
订单回流接口 APIRequest
alibaba.alsc.crm.open.order.backflow

回流isv订单接口
*/
type AlibabaAlscCrmOpenOrderBackflowRequest struct {
    model.Params

    // 入参
    paramOrderBackflowOpenReq   *OrderBackflowOpenReq 

}

func NewAlibabaAlscCrmOpenOrderBackflowRequest() *AlibabaAlscCrmOpenOrderBackflowRequest{
    return &AlibabaAlscCrmOpenOrderBackflowRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlscCrmOpenOrderBackflowRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.open.order.backflow"
}

func (r AlibabaAlscCrmOpenOrderBackflowRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlscCrmOpenOrderBackflowRequest) SetParamOrderBackflowOpenReq(paramOrderBackflowOpenReq *OrderBackflowOpenReq) error {
    r.paramOrderBackflowOpenReq = paramOrderBackflowOpenReq
    r.Set("param_order_backflow_open_req", paramOrderBackflowOpenReq)
    return nil
}

func (r AlibabaAlscCrmOpenOrderBackflowRequest) GetParamOrderBackflowOpenReq() *OrderBackflowOpenReq {
    return r.paramOrderBackflowOpenReq
}

