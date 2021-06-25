package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
校验支付链路中的积分抵扣是否合法 APIRequest
alibaba.alsc.crm.point.chkpntbypay

校验支付链路中的积分抵扣是否合法
*/
type AlibabaAlscCrmPointChkpntbypayRequest struct {
    model.Params

    // 入参
    paramConsumePointByPayOpenReq   *ConsumePointByPayOpenReq 

}

func NewAlibabaAlscCrmPointChkpntbypayRequest() *AlibabaAlscCrmPointChkpntbypayRequest{
    return &AlibabaAlscCrmPointChkpntbypayRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlscCrmPointChkpntbypayRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.point.chkpntbypay"
}

func (r AlibabaAlscCrmPointChkpntbypayRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlscCrmPointChkpntbypayRequest) SetParamConsumePointByPayOpenReq(paramConsumePointByPayOpenReq *ConsumePointByPayOpenReq) error {
    r.paramConsumePointByPayOpenReq = paramConsumePointByPayOpenReq
    r.Set("param_consume_point_by_pay_open_req", paramConsumePointByPayOpenReq)
    return nil
}

func (r AlibabaAlscCrmPointChkpntbypayRequest) GetParamConsumePointByPayOpenReq() *ConsumePointByPayOpenReq {
    return r.paramConsumePointByPayOpenReq
}

