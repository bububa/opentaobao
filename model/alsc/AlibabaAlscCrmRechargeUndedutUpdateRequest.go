package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
储值消费退款(逆向) API请求
alibaba.alsc.crm.recharge.undedut.update

新增储值消费退款接口
*/
type AlibabaAlscCrmRechargeUndedutUpdateRequest struct {
    model.Params
    // 入参
    _paramUndedutOpenReq   *UndedutOpenReq
}

// 初始化AlibabaAlscCrmRechargeUndedutUpdateRequest对象
func NewAlibabaAlscCrmRechargeUndedutUpdateRequest() *AlibabaAlscCrmRechargeUndedutUpdateRequest{
    return &AlibabaAlscCrmRechargeUndedutUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmRechargeUndedutUpdateRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.recharge.undedut.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmRechargeUndedutUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamUndedutOpenReq Setter
// 入参
func (r *AlibabaAlscCrmRechargeUndedutUpdateRequest) SetParamUndedutOpenReq(_paramUndedutOpenReq *UndedutOpenReq) error {
    r._paramUndedutOpenReq = _paramUndedutOpenReq
    r.Set("param_undedut_open_req", _paramUndedutOpenReq)
    return nil
}

// ParamUndedutOpenReq Getter
func (r AlibabaAlscCrmRechargeUndedutUpdateRequest) GetParamUndedutOpenReq() *UndedutOpenReq {
    return r._paramUndedutOpenReq
}
