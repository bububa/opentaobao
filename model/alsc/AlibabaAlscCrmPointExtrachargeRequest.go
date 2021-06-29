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
type AlibabaAlscCrmPointExtrachargeRequest struct {
    model.Params
    // 入参
    _paramExtraChargePointOpenReq   *ExtraChargePointOpenReq
}

// 初始化AlibabaAlscCrmPointExtrachargeRequest对象
func NewAlibabaAlscCrmPointExtrachargeRequest() *AlibabaAlscCrmPointExtrachargeRequest{
    return &AlibabaAlscCrmPointExtrachargeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmPointExtrachargeRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.point.extracharge"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmPointExtrachargeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamExtraChargePointOpenReq Setter
// 入参
func (r *AlibabaAlscCrmPointExtrachargeRequest) SetParamExtraChargePointOpenReq(_paramExtraChargePointOpenReq *ExtraChargePointOpenReq) error {
    r._paramExtraChargePointOpenReq = _paramExtraChargePointOpenReq
    r.Set("param_extra_charge_point_open_req", _paramExtraChargePointOpenReq)
    return nil
}

// ParamExtraChargePointOpenReq Getter
func (r AlibabaAlscCrmPointExtrachargeRequest) GetParamExtraChargePointOpenReq() *ExtraChargePointOpenReq {
    return r._paramExtraChargePointOpenReq
}
