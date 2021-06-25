package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
储值操作接口 APIRequest
alibaba.alsc.crm.open.recharge.operate

储值操作接口
*/
type AlibabaAlscCrmOpenRechargeOperateRequest struct {
    model.Params

    // 储值操作参数
    paramRechargeOperateOpenReq   *RechargeOperateOpenReq 

}

func NewAlibabaAlscCrmOpenRechargeOperateRequest() *AlibabaAlscCrmOpenRechargeOperateRequest{
    return &AlibabaAlscCrmOpenRechargeOperateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlscCrmOpenRechargeOperateRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.open.recharge.operate"
}

func (r AlibabaAlscCrmOpenRechargeOperateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlscCrmOpenRechargeOperateRequest) SetParamRechargeOperateOpenReq(paramRechargeOperateOpenReq *RechargeOperateOpenReq) error {
    r.paramRechargeOperateOpenReq = paramRechargeOperateOpenReq
    r.Set("param_recharge_operate_open_req", paramRechargeOperateOpenReq)
    return nil
}

func (r AlibabaAlscCrmOpenRechargeOperateRequest) GetParamRechargeOperateOpenReq() *RechargeOperateOpenReq {
    return r.paramRechargeOperateOpenReq
}

