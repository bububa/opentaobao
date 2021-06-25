package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
储值核销预先校验 APIRequest
alibaba.alsc.crm.recharge.dedutprecheck.get

储值核销预先校验接口
*/
type AlibabaAlscCrmRechargeDedutprecheckGetRequest struct {
    model.Params

    // 入参
    paramDeductPreCheckOpenReq   *DeductPreCheckOpenReq 

}

func NewAlibabaAlscCrmRechargeDedutprecheckGetRequest() *AlibabaAlscCrmRechargeDedutprecheckGetRequest{
    return &AlibabaAlscCrmRechargeDedutprecheckGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlscCrmRechargeDedutprecheckGetRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.recharge.dedutprecheck.get"
}

func (r AlibabaAlscCrmRechargeDedutprecheckGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlscCrmRechargeDedutprecheckGetRequest) SetParamDeductPreCheckOpenReq(paramDeductPreCheckOpenReq *DeductPreCheckOpenReq) error {
    r.paramDeductPreCheckOpenReq = paramDeductPreCheckOpenReq
    r.Set("param_deduct_pre_check_open_req", paramDeductPreCheckOpenReq)
    return nil
}

func (r AlibabaAlscCrmRechargeDedutprecheckGetRequest) GetParamDeductPreCheckOpenReq() *DeductPreCheckOpenReq {
    return r.paramDeductPreCheckOpenReq
}

