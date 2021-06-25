package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
积分补录 APIRequest
alibaba.alsc.crm.point.extracharge

积分补录
*/
type AlibabaAlscCrmPointExtrachargeRequest struct {
    model.Params

    // 入参
    paramExtraChargePointOpenReq   *ExtraChargePointOpenReq 

}

func NewAlibabaAlscCrmPointExtrachargeRequest() *AlibabaAlscCrmPointExtrachargeRequest{
    return &AlibabaAlscCrmPointExtrachargeRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlscCrmPointExtrachargeRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.point.extracharge"
}

func (r AlibabaAlscCrmPointExtrachargeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlscCrmPointExtrachargeRequest) SetParamExtraChargePointOpenReq(paramExtraChargePointOpenReq *ExtraChargePointOpenReq) error {
    r.paramExtraChargePointOpenReq = paramExtraChargePointOpenReq
    r.Set("param_extra_charge_point_open_req", paramExtraChargePointOpenReq)
    return nil
}

func (r AlibabaAlscCrmPointExtrachargeRequest) GetParamExtraChargePointOpenReq() *ExtraChargePointOpenReq {
    return r.paramExtraChargePointOpenReq
}

