package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询储值账户信息 APIRequest
alibaba.alsc.crm.recharge.account.get

查询储值账户信息接口
*/
type AlibabaAlscCrmRechargeAccountGetRequest struct {
    model.Params

    // 入参
    paramQueryRechargeAccountOpenReq   *QueryRechargeAccountOpenReq 

}

func NewAlibabaAlscCrmRechargeAccountGetRequest() *AlibabaAlscCrmRechargeAccountGetRequest{
    return &AlibabaAlscCrmRechargeAccountGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlscCrmRechargeAccountGetRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.recharge.account.get"
}

func (r AlibabaAlscCrmRechargeAccountGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlscCrmRechargeAccountGetRequest) SetParamQueryRechargeAccountOpenReq(paramQueryRechargeAccountOpenReq *QueryRechargeAccountOpenReq) error {
    r.paramQueryRechargeAccountOpenReq = paramQueryRechargeAccountOpenReq
    r.Set("param_query_recharge_account_open_req", paramQueryRechargeAccountOpenReq)
    return nil
}

func (r AlibabaAlscCrmRechargeAccountGetRequest) GetParamQueryRechargeAccountOpenReq() *QueryRechargeAccountOpenReq {
    return r.paramQueryRechargeAccountOpenReq
}

