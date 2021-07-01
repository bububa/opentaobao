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
type AlibabaAlscCrmOpenOrderBackflowAPIRequest struct {
    model.Params
    // 入参
    _paramOrderBackflowOpenReq   *OrderBackflowOpenReq
}

// 初始化AlibabaAlscCrmOpenOrderBackflowAPIRequest对象
func NewAlibabaAlscCrmOpenOrderBackflowRequest() *AlibabaAlscCrmOpenOrderBackflowAPIRequest{
    return &AlibabaAlscCrmOpenOrderBackflowAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmOpenOrderBackflowAPIRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.open.order.backflow"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmOpenOrderBackflowAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamOrderBackflowOpenReq Setter
// 入参
func (r *AlibabaAlscCrmOpenOrderBackflowAPIRequest) SetParamOrderBackflowOpenReq(_paramOrderBackflowOpenReq *OrderBackflowOpenReq) error {
    r._paramOrderBackflowOpenReq = _paramOrderBackflowOpenReq
    r.Set("param_order_backflow_open_req", _paramOrderBackflowOpenReq)
    return nil
}

// ParamOrderBackflowOpenReq Getter
func (r AlibabaAlscCrmOpenOrderBackflowAPIRequest) GetParamOrderBackflowOpenReq() *OrderBackflowOpenReq {
    return r._paramOrderBackflowOpenReq
}
