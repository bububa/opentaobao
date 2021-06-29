package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
储值流水详细 API请求
alibaba.alsc.crm.recharge.account.flowdetail.get

查询储值流水详细接口
*/
type AlibabaAlscCrmRechargeAccountFlowdetailGetRequest struct {
    model.Params
    // 入参
    _paramQueryRechargeAccountFlowOpenReq   *QueryRechargeAccountFlowOpenReq
}

// 初始化AlibabaAlscCrmRechargeAccountFlowdetailGetRequest对象
func NewAlibabaAlscCrmRechargeAccountFlowdetailGetRequest() *AlibabaAlscCrmRechargeAccountFlowdetailGetRequest{
    return &AlibabaAlscCrmRechargeAccountFlowdetailGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmRechargeAccountFlowdetailGetRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.recharge.account.flowdetail.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmRechargeAccountFlowdetailGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamQueryRechargeAccountFlowOpenReq Setter
// 入参
func (r *AlibabaAlscCrmRechargeAccountFlowdetailGetRequest) SetParamQueryRechargeAccountFlowOpenReq(_paramQueryRechargeAccountFlowOpenReq *QueryRechargeAccountFlowOpenReq) error {
    r._paramQueryRechargeAccountFlowOpenReq = _paramQueryRechargeAccountFlowOpenReq
    r.Set("param_query_recharge_account_flow_open_req", _paramQueryRechargeAccountFlowOpenReq)
    return nil
}

// ParamQueryRechargeAccountFlowOpenReq Getter
func (r AlibabaAlscCrmRechargeAccountFlowdetailGetRequest) GetParamQueryRechargeAccountFlowOpenReq() *QueryRechargeAccountFlowOpenReq {
    return r._paramQueryRechargeAccountFlowOpenReq
}
