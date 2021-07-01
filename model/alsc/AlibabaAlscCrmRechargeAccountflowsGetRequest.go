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
type AlibabaAlscCrmRechargeAccountflowsGetAPIRequest struct {
    model.Params
    // 入参
    _paramPageQueryAccountFlowsOpenReq   *PageQueryAccountFlowsOpenReq
}

// 初始化AlibabaAlscCrmRechargeAccountflowsGetAPIRequest对象
func NewAlibabaAlscCrmRechargeAccountflowsGetRequest() *AlibabaAlscCrmRechargeAccountflowsGetAPIRequest{
    return &AlibabaAlscCrmRechargeAccountflowsGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmRechargeAccountflowsGetAPIRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.recharge.accountflows.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmRechargeAccountflowsGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamPageQueryAccountFlowsOpenReq Setter
// 入参
func (r *AlibabaAlscCrmRechargeAccountflowsGetAPIRequest) SetParamPageQueryAccountFlowsOpenReq(_paramPageQueryAccountFlowsOpenReq *PageQueryAccountFlowsOpenReq) error {
    r._paramPageQueryAccountFlowsOpenReq = _paramPageQueryAccountFlowsOpenReq
    r.Set("param_page_query_account_flows_open_req", _paramPageQueryAccountFlowsOpenReq)
    return nil
}

// ParamPageQueryAccountFlowsOpenReq Getter
func (r AlibabaAlscCrmRechargeAccountflowsGetAPIRequest) GetParamPageQueryAccountFlowsOpenReq() *PageQueryAccountFlowsOpenReq {
    return r._paramPageQueryAccountFlowsOpenReq
}
