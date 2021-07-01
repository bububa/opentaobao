package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
储值充值 API请求
alibaba.alsc.crm.recharge.charge.update

顾客储值账户充值
*/
type AlibabaAlscCrmRechargeChargeUpdateAPIRequest struct {
    model.Params
    // 入参
    _paramRechargeOpenReq   *RechargeOpenReq
}

// 初始化AlibabaAlscCrmRechargeChargeUpdateAPIRequest对象
func NewAlibabaAlscCrmRechargeChargeUpdateRequest() *AlibabaAlscCrmRechargeChargeUpdateAPIRequest{
    return &AlibabaAlscCrmRechargeChargeUpdateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmRechargeChargeUpdateAPIRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.recharge.charge.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmRechargeChargeUpdateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamRechargeOpenReq Setter
// 入参
func (r *AlibabaAlscCrmRechargeChargeUpdateAPIRequest) SetParamRechargeOpenReq(_paramRechargeOpenReq *RechargeOpenReq) error {
    r._paramRechargeOpenReq = _paramRechargeOpenReq
    r.Set("param_recharge_open_req", _paramRechargeOpenReq)
    return nil
}

// ParamRechargeOpenReq Getter
func (r AlibabaAlscCrmRechargeChargeUpdateAPIRequest) GetParamRechargeOpenReq() *RechargeOpenReq {
    return r._paramRechargeOpenReq
}
