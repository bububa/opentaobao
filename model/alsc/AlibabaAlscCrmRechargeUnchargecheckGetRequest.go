package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
储值账户退充值校验 API请求
alibaba.alsc.crm.recharge.unchargecheck.get

储值账户退充值校验接口
*/
type AlibabaAlscCrmRechargeUnchargecheckGetRequest struct {
    model.Params
    // 入参
    _paramUnchargeCheckOpenReq   *UnchargeCheckOpenReq
}

// 初始化AlibabaAlscCrmRechargeUnchargecheckGetRequest对象
func NewAlibabaAlscCrmRechargeUnchargecheckGetRequest() *AlibabaAlscCrmRechargeUnchargecheckGetRequest{
    return &AlibabaAlscCrmRechargeUnchargecheckGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmRechargeUnchargecheckGetRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.recharge.unchargecheck.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmRechargeUnchargecheckGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamUnchargeCheckOpenReq Setter
// 入参
func (r *AlibabaAlscCrmRechargeUnchargecheckGetRequest) SetParamUnchargeCheckOpenReq(_paramUnchargeCheckOpenReq *UnchargeCheckOpenReq) error {
    r._paramUnchargeCheckOpenReq = _paramUnchargeCheckOpenReq
    r.Set("param_uncharge_check_open_req", _paramUnchargeCheckOpenReq)
    return nil
}

// ParamUnchargeCheckOpenReq Getter
func (r AlibabaAlscCrmRechargeUnchargecheckGetRequest) GetParamUnchargeCheckOpenReq() *UnchargeCheckOpenReq {
    return r._paramUnchargeCheckOpenReq
}
