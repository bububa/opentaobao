package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
储值账户退充值校验 APIRequest
alibaba.alsc.crm.recharge.unchargecheck.get

储值账户退充值校验接口
*/
type AlibabaAlscCrmRechargeUnchargecheckGetRequest struct {
    model.Params

    // 入参
    paramUnchargeCheckOpenReq   *UnchargeCheckOpenReq 

}

func NewAlibabaAlscCrmRechargeUnchargecheckGetRequest() *AlibabaAlscCrmRechargeUnchargecheckGetRequest{
    return &AlibabaAlscCrmRechargeUnchargecheckGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlscCrmRechargeUnchargecheckGetRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.recharge.unchargecheck.get"
}

func (r AlibabaAlscCrmRechargeUnchargecheckGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlscCrmRechargeUnchargecheckGetRequest) SetParamUnchargeCheckOpenReq(paramUnchargeCheckOpenReq *UnchargeCheckOpenReq) error {
    r.paramUnchargeCheckOpenReq = paramUnchargeCheckOpenReq
    r.Set("param_uncharge_check_open_req", paramUnchargeCheckOpenReq)
    return nil
}

func (r AlibabaAlscCrmRechargeUnchargecheckGetRequest) GetParamUnchargeCheckOpenReq() *UnchargeCheckOpenReq {
    return r.paramUnchargeCheckOpenReq
}

