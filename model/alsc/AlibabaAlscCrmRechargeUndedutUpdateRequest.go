package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
储值消费退款(逆向) APIRequest
alibaba.alsc.crm.recharge.undedut.update

新增储值消费退款接口
*/
type AlibabaAlscCrmRechargeUndedutUpdateRequest struct {
    model.Params

    // 入参
    paramUndedutOpenReq   *UndedutOpenReq 

}

func NewAlibabaAlscCrmRechargeUndedutUpdateRequest() *AlibabaAlscCrmRechargeUndedutUpdateRequest{
    return &AlibabaAlscCrmRechargeUndedutUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlscCrmRechargeUndedutUpdateRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.recharge.undedut.update"
}

func (r AlibabaAlscCrmRechargeUndedutUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlscCrmRechargeUndedutUpdateRequest) SetParamUndedutOpenReq(paramUndedutOpenReq *UndedutOpenReq) error {
    r.paramUndedutOpenReq = paramUndedutOpenReq
    r.Set("param_undedut_open_req", paramUndedutOpenReq)
    return nil
}

func (r AlibabaAlscCrmRechargeUndedutUpdateRequest) GetParamUndedutOpenReq() *UndedutOpenReq {
    return r.paramUndedutOpenReq
}

