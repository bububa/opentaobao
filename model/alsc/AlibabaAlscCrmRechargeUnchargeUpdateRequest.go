package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
充值退款 API请求
alibaba.alsc.crm.recharge.uncharge.update

充值退款
*/
type AlibabaAlscCrmRechargeUnchargeUpdateRequest struct {
    model.Params
    // 入参
    paramUnchargeOpenReq   *UnchargeOpenReq
}

// 初始化AlibabaAlscCrmRechargeUnchargeUpdateRequest对象
func NewAlibabaAlscCrmRechargeUnchargeUpdateRequest() *AlibabaAlscCrmRechargeUnchargeUpdateRequest{
    return &AlibabaAlscCrmRechargeUnchargeUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmRechargeUnchargeUpdateRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.recharge.uncharge.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmRechargeUnchargeUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamUnchargeOpenReq Setter
// 入参
func (r *AlibabaAlscCrmRechargeUnchargeUpdateRequest) SetParamUnchargeOpenReq(paramUnchargeOpenReq *UnchargeOpenReq) error {
    r.paramUnchargeOpenReq = paramUnchargeOpenReq
    r.Set("param_uncharge_open_req", paramUnchargeOpenReq)
    return nil
}

// ParamUnchargeOpenReq Getter
func (r AlibabaAlscCrmRechargeUnchargeUpdateRequest) GetParamUnchargeOpenReq() *UnchargeOpenReq {
    return r.paramUnchargeOpenReq
}
