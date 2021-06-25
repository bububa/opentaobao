package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
储值账户充值前校验 APIRequest
alibaba.alsc.crm.recharge.chargeprecheck.get

储值账户充值前校验接口
*/
type AlibabaAlscCrmRechargeChargeprecheckGetRequest struct {
    model.Params

    // 入参
    paramChargePreCheckOpenReq   *ChargePreCheckOpenReq 

}

func NewAlibabaAlscCrmRechargeChargeprecheckGetRequest() *AlibabaAlscCrmRechargeChargeprecheckGetRequest{
    return &AlibabaAlscCrmRechargeChargeprecheckGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlscCrmRechargeChargeprecheckGetRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.recharge.chargeprecheck.get"
}

func (r AlibabaAlscCrmRechargeChargeprecheckGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlscCrmRechargeChargeprecheckGetRequest) SetParamChargePreCheckOpenReq(paramChargePreCheckOpenReq *ChargePreCheckOpenReq) error {
    r.paramChargePreCheckOpenReq = paramChargePreCheckOpenReq
    r.Set("param_charge_pre_check_open_req", paramChargePreCheckOpenReq)
    return nil
}

func (r AlibabaAlscCrmRechargeChargeprecheckGetRequest) GetParamChargePreCheckOpenReq() *ChargePreCheckOpenReq {
    return r.paramChargePreCheckOpenReq
}

