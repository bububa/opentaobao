package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
储值消费 APIRequest
alibaba.alsc.crm.recharge.dedut.update

增加储值消费接口
*/
type AlibabaAlscCrmRechargeDedutUpdateRequest struct {
    model.Params

    // 入参
    paramDedutOpenReq   *DedutOpenReq 

}

func NewAlibabaAlscCrmRechargeDedutUpdateRequest() *AlibabaAlscCrmRechargeDedutUpdateRequest{
    return &AlibabaAlscCrmRechargeDedutUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlscCrmRechargeDedutUpdateRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.recharge.dedut.update"
}

func (r AlibabaAlscCrmRechargeDedutUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlscCrmRechargeDedutUpdateRequest) SetParamDedutOpenReq(paramDedutOpenReq *DedutOpenReq) error {
    r.paramDedutOpenReq = paramDedutOpenReq
    r.Set("param_dedut_open_req", paramDedutOpenReq)
    return nil
}

func (r AlibabaAlscCrmRechargeDedutUpdateRequest) GetParamDedutOpenReq() *DedutOpenReq {
    return r.paramDedutOpenReq
}

