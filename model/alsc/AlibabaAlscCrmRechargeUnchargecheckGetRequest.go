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
type AlibabaAlscCrmRechargeUnchargecheckGetAPIRequest struct {
    model.Params
    // 入参
    _paramUnchargeCheckOpenReq   *UnchargeCheckOpenReq
}

// 初始化AlibabaAlscCrmRechargeUnchargecheckGetAPIRequest对象
func NewAlibabaAlscCrmRechargeUnchargecheckGetRequest() *AlibabaAlscCrmRechargeUnchargecheckGetAPIRequest{
    return &AlibabaAlscCrmRechargeUnchargecheckGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmRechargeUnchargecheckGetAPIRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.recharge.unchargecheck.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmRechargeUnchargecheckGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamUnchargeCheckOpenReq Setter
// 入参
func (r *AlibabaAlscCrmRechargeUnchargecheckGetAPIRequest) SetParamUnchargeCheckOpenReq(_paramUnchargeCheckOpenReq *UnchargeCheckOpenReq) error {
    r._paramUnchargeCheckOpenReq = _paramUnchargeCheckOpenReq
    r.Set("param_uncharge_check_open_req", _paramUnchargeCheckOpenReq)
    return nil
}

// ParamUnchargeCheckOpenReq Getter
func (r AlibabaAlscCrmRechargeUnchargecheckGetAPIRequest) GetParamUnchargeCheckOpenReq() *UnchargeCheckOpenReq {
    return r._paramUnchargeCheckOpenReq
}
