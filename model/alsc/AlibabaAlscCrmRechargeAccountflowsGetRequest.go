package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分页查询储值流水 APIRequest
alibaba.alsc.crm.recharge.accountflows.get

增加分页查询储值流水接口
*/
type AlibabaAlscCrmRechargeAccountflowsGetRequest struct {
    model.Params

    // 入参
    paramPageQueryAccountFlowsOpenReq   *PageQueryAccountFlowsOpenReq 

}

func NewAlibabaAlscCrmRechargeAccountflowsGetRequest() *AlibabaAlscCrmRechargeAccountflowsGetRequest{
    return &AlibabaAlscCrmRechargeAccountflowsGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlscCrmRechargeAccountflowsGetRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.recharge.accountflows.get"
}

func (r AlibabaAlscCrmRechargeAccountflowsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlscCrmRechargeAccountflowsGetRequest) SetParamPageQueryAccountFlowsOpenReq(paramPageQueryAccountFlowsOpenReq *PageQueryAccountFlowsOpenReq) error {
    r.paramPageQueryAccountFlowsOpenReq = paramPageQueryAccountFlowsOpenReq
    r.Set("param_page_query_account_flows_open_req", paramPageQueryAccountFlowsOpenReq)
    return nil
}

func (r AlibabaAlscCrmRechargeAccountflowsGetRequest) GetParamPageQueryAccountFlowsOpenReq() *PageQueryAccountFlowsOpenReq {
    return r.paramPageQueryAccountFlowsOpenReq
}

