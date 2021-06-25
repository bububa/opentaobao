package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
积分补扣 APIRequest
alibaba.alsc.crm.point.extra.consume

积分补扣
*/
type AlibabaAlscCrmPointExtraConsumeRequest struct {
    model.Params

    // 入参
    paramExtraConsumePointOpenReq   *ExtraConsumePointOpenReq 

}

func NewAlibabaAlscCrmPointExtraConsumeRequest() *AlibabaAlscCrmPointExtraConsumeRequest{
    return &AlibabaAlscCrmPointExtraConsumeRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlscCrmPointExtraConsumeRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.point.extra.consume"
}

func (r AlibabaAlscCrmPointExtraConsumeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlscCrmPointExtraConsumeRequest) SetParamExtraConsumePointOpenReq(paramExtraConsumePointOpenReq *ExtraConsumePointOpenReq) error {
    r.paramExtraConsumePointOpenReq = paramExtraConsumePointOpenReq
    r.Set("param_extra_consume_point_open_req", paramExtraConsumePointOpenReq)
    return nil
}

func (r AlibabaAlscCrmPointExtraConsumeRequest) GetParamExtraConsumePointOpenReq() *ExtraConsumePointOpenReq {
    return r.paramExtraConsumePointOpenReq
}

