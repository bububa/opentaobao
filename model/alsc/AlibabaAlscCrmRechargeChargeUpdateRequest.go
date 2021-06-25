package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
储值充值 APIRequest
alibaba.alsc.crm.recharge.charge.update

顾客储值账户充值
*/
type AlibabaAlscCrmRechargeChargeUpdateRequest struct {
    model.Params

    // 入参
    paramRechargeOpenReq   *RechargeOpenReq 

}

func NewAlibabaAlscCrmRechargeChargeUpdateRequest() *AlibabaAlscCrmRechargeChargeUpdateRequest{
    return &AlibabaAlscCrmRechargeChargeUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlscCrmRechargeChargeUpdateRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.recharge.charge.update"
}

func (r AlibabaAlscCrmRechargeChargeUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlscCrmRechargeChargeUpdateRequest) SetParamRechargeOpenReq(paramRechargeOpenReq *RechargeOpenReq) error {
    r.paramRechargeOpenReq = paramRechargeOpenReq
    r.Set("param_recharge_open_req", paramRechargeOpenReq)
    return nil
}

func (r AlibabaAlscCrmRechargeChargeUpdateRequest) GetParamRechargeOpenReq() *RechargeOpenReq {
    return r.paramRechargeOpenReq
}

