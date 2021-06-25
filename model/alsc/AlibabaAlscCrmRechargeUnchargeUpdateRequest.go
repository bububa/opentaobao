package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
充值退款 APIRequest
alibaba.alsc.crm.recharge.uncharge.update

充值退款
*/
type AlibabaAlscCrmRechargeUnchargeUpdateRequest struct {
    model.Params

    // 入参
    paramUnchargeOpenReq   *UnchargeOpenReq 

}

func NewAlibabaAlscCrmRechargeUnchargeUpdateRequest() *AlibabaAlscCrmRechargeUnchargeUpdateRequest{
    return &AlibabaAlscCrmRechargeUnchargeUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlscCrmRechargeUnchargeUpdateRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.recharge.uncharge.update"
}

func (r AlibabaAlscCrmRechargeUnchargeUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlscCrmRechargeUnchargeUpdateRequest) SetParamUnchargeOpenReq(paramUnchargeOpenReq *UnchargeOpenReq) error {
    r.paramUnchargeOpenReq = paramUnchargeOpenReq
    r.Set("param_uncharge_open_req", paramUnchargeOpenReq)
    return nil
}

func (r AlibabaAlscCrmRechargeUnchargeUpdateRequest) GetParamUnchargeOpenReq() *UnchargeOpenReq {
    return r.paramUnchargeOpenReq
}

