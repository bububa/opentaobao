package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
储值账户充值前校验 API请求
alibaba.alsc.crm.recharge.chargeprecheck.get

储值账户充值前校验接口
*/
type AlibabaAlscCrmRechargeChargeprecheckGetRequest struct {
    model.Params
    // 入参
    paramChargePreCheckOpenReq   *ChargePreCheckOpenReq
}

// 初始化AlibabaAlscCrmRechargeChargeprecheckGetRequest对象
func NewAlibabaAlscCrmRechargeChargeprecheckGetRequest() *AlibabaAlscCrmRechargeChargeprecheckGetRequest{
    return &AlibabaAlscCrmRechargeChargeprecheckGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmRechargeChargeprecheckGetRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.recharge.chargeprecheck.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmRechargeChargeprecheckGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamChargePreCheckOpenReq Setter
// 入参
func (r *AlibabaAlscCrmRechargeChargeprecheckGetRequest) SetParamChargePreCheckOpenReq(paramChargePreCheckOpenReq *ChargePreCheckOpenReq) error {
    r.paramChargePreCheckOpenReq = paramChargePreCheckOpenReq
    r.Set("param_charge_pre_check_open_req", paramChargePreCheckOpenReq)
    return nil
}

// ParamChargePreCheckOpenReq Getter
func (r AlibabaAlscCrmRechargeChargeprecheckGetRequest) GetParamChargePreCheckOpenReq() *ChargePreCheckOpenReq {
    return r.paramChargePreCheckOpenReq
}
