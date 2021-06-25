package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
积分消费回退 APIRequest
alibaba.alsc.crm.point.reversepoint

积分消费回退
*/
type AlibabaAlscCrmPointReversepointRequest struct {
    model.Params

    // 入参
    paramReverseConsumePointOpenReq   *ReverseConsumePointOpenReq 

}

func NewAlibabaAlscCrmPointReversepointRequest() *AlibabaAlscCrmPointReversepointRequest{
    return &AlibabaAlscCrmPointReversepointRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlscCrmPointReversepointRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.point.reversepoint"
}

func (r AlibabaAlscCrmPointReversepointRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlscCrmPointReversepointRequest) SetParamReverseConsumePointOpenReq(paramReverseConsumePointOpenReq *ReverseConsumePointOpenReq) error {
    r.paramReverseConsumePointOpenReq = paramReverseConsumePointOpenReq
    r.Set("param_reverse_consume_point_open_req", paramReverseConsumePointOpenReq)
    return nil
}

func (r AlibabaAlscCrmPointReversepointRequest) GetParamReverseConsumePointOpenReq() *ReverseConsumePointOpenReq {
    return r.paramReverseConsumePointOpenReq
}

