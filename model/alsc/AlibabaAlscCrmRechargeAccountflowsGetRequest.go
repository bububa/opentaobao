package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分页查询储值流水 API请求
alibaba.alsc.crm.recharge.accountflows.get

增加分页查询储值流水接口
*/
type AlibabaAlscCrmRechargeAccountflowsGetRequest struct {
    model.Params
    // 入参
    paramPageQueryAccountFlowsOpenReq   *PageQueryAccountFlowsOpenReq
}

// 初始化AlibabaAlscCrmRechargeAccountflowsGetRequest对象
func NewAlibabaAlscCrmRechargeAccountflowsGetRequest() *AlibabaAlscCrmRechargeAccountflowsGetRequest{
    return &AlibabaAlscCrmRechargeAccountflowsGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmRechargeAccountflowsGetRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.recharge.accountflows.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmRechargeAccountflowsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamPageQueryAccountFlowsOpenReq Setter
// 入参
func (r *AlibabaAlscCrmRechargeAccountflowsGetRequest) SetParamPageQueryAccountFlowsOpenReq(paramPageQueryAccountFlowsOpenReq *PageQueryAccountFlowsOpenReq) error {
    r.paramPageQueryAccountFlowsOpenReq = paramPageQueryAccountFlowsOpenReq
    r.Set("param_page_query_account_flows_open_req", paramPageQueryAccountFlowsOpenReq)
    return nil
}

// ParamPageQueryAccountFlowsOpenReq Getter
func (r AlibabaAlscCrmRechargeAccountflowsGetRequest) GetParamPageQueryAccountFlowsOpenReq() *PageQueryAccountFlowsOpenReq {
    return r.paramPageQueryAccountFlowsOpenReq
}
