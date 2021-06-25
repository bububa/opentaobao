package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
积分抵现 APIRequest
alibaba.alsc.crm.point.consumepoint

积分抵现
*/
type AlibabaAlscCrmPointConsumepointRequest struct {
    model.Params

    // 入参
    paramConsumePointOpenReq   *ConsumePointOpenReq 

}

func NewAlibabaAlscCrmPointConsumepointRequest() *AlibabaAlscCrmPointConsumepointRequest{
    return &AlibabaAlscCrmPointConsumepointRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlscCrmPointConsumepointRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.point.consumepoint"
}

func (r AlibabaAlscCrmPointConsumepointRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlscCrmPointConsumepointRequest) SetParamConsumePointOpenReq(paramConsumePointOpenReq *ConsumePointOpenReq) error {
    r.paramConsumePointOpenReq = paramConsumePointOpenReq
    r.Set("param_consume_point_open_req", paramConsumePointOpenReq)
    return nil
}

func (r AlibabaAlscCrmPointConsumepointRequest) GetParamConsumePointOpenReq() *ConsumePointOpenReq {
    return r.paramConsumePointOpenReq
}

