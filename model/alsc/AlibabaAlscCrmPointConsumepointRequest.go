package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
积分抵现 API请求
alibaba.alsc.crm.point.consumepoint

积分抵现
*/
type AlibabaAlscCrmPointConsumepointRequest struct {
    model.Params
    // 入参
    _paramConsumePointOpenReq   *ConsumePointOpenReq
}

// 初始化AlibabaAlscCrmPointConsumepointRequest对象
func NewAlibabaAlscCrmPointConsumepointRequest() *AlibabaAlscCrmPointConsumepointRequest{
    return &AlibabaAlscCrmPointConsumepointRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmPointConsumepointRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.point.consumepoint"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmPointConsumepointRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamConsumePointOpenReq Setter
// 入参
func (r *AlibabaAlscCrmPointConsumepointRequest) SetParamConsumePointOpenReq(_paramConsumePointOpenReq *ConsumePointOpenReq) error {
    r._paramConsumePointOpenReq = _paramConsumePointOpenReq
    r.Set("param_consume_point_open_req", _paramConsumePointOpenReq)
    return nil
}

// ParamConsumePointOpenReq Getter
func (r AlibabaAlscCrmPointConsumepointRequest) GetParamConsumePointOpenReq() *ConsumePointOpenReq {
    return r._paramConsumePointOpenReq
}
