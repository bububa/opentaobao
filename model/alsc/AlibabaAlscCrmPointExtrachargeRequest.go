package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
积分补录 API请求
alibaba.alsc.crm.point.extracharge

积分补录
*/
type AlibabaAlscCrmPointExtrachargeAPIRequest struct {
    model.Params
    // 入参
    _paramExtraChargePointOpenReq   *ExtraChargePointOpenReq
}

// 初始化AlibabaAlscCrmPointExtrachargeAPIRequest对象
func NewAlibabaAlscCrmPointExtrachargeRequest() *AlibabaAlscCrmPointExtrachargeAPIRequest{
    return &AlibabaAlscCrmPointExtrachargeAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmPointExtrachargeAPIRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.point.extracharge"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmPointExtrachargeAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamExtraChargePointOpenReq Setter
// 入参
func (r *AlibabaAlscCrmPointExtrachargeAPIRequest) SetParamExtraChargePointOpenReq(_paramExtraChargePointOpenReq *ExtraChargePointOpenReq) error {
    r._paramExtraChargePointOpenReq = _paramExtraChargePointOpenReq
    r.Set("param_extra_charge_point_open_req", _paramExtraChargePointOpenReq)
    return nil
}

// ParamExtraChargePointOpenReq Getter
func (r AlibabaAlscCrmPointExtrachargeAPIRequest) GetParamExtraChargePointOpenReq() *ExtraChargePointOpenReq {
    return r._paramExtraChargePointOpenReq
}
