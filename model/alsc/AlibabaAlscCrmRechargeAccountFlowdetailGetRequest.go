package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
储值流水详细 APIRequest
alibaba.alsc.crm.recharge.account.flowdetail.get

查询储值流水详细接口
*/
type AlibabaAlscCrmRechargeAccountFlowdetailGetRequest struct {
    model.Params

    // 入参
    paramQueryRechargeAccountFlowOpenReq   *QueryRechargeAccountFlowOpenReq 

}

func NewAlibabaAlscCrmRechargeAccountFlowdetailGetRequest() *AlibabaAlscCrmRechargeAccountFlowdetailGetRequest{
    return &AlibabaAlscCrmRechargeAccountFlowdetailGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlscCrmRechargeAccountFlowdetailGetRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.recharge.account.flowdetail.get"
}

func (r AlibabaAlscCrmRechargeAccountFlowdetailGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlscCrmRechargeAccountFlowdetailGetRequest) SetParamQueryRechargeAccountFlowOpenReq(paramQueryRechargeAccountFlowOpenReq *QueryRechargeAccountFlowOpenReq) error {
    r.paramQueryRechargeAccountFlowOpenReq = paramQueryRechargeAccountFlowOpenReq
    r.Set("param_query_recharge_account_flow_open_req", paramQueryRechargeAccountFlowOpenReq)
    return nil
}

func (r AlibabaAlscCrmRechargeAccountFlowdetailGetRequest) GetParamQueryRechargeAccountFlowOpenReq() *QueryRechargeAccountFlowOpenReq {
    return r.paramQueryRechargeAccountFlowOpenReq
}

