package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
积分消费回退 API请求
alibaba.alsc.crm.point.reversepoint

积分消费回退
*/
type AlibabaAlscCrmPointReversepointRequest struct {
    model.Params
    // 入参
    paramReverseConsumePointOpenReq   *ReverseConsumePointOpenReq
}

// 初始化AlibabaAlscCrmPointReversepointRequest对象
func NewAlibabaAlscCrmPointReversepointRequest() *AlibabaAlscCrmPointReversepointRequest{
    return &AlibabaAlscCrmPointReversepointRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmPointReversepointRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.point.reversepoint"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmPointReversepointRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamReverseConsumePointOpenReq Setter
// 入参
func (r *AlibabaAlscCrmPointReversepointRequest) SetParamReverseConsumePointOpenReq(paramReverseConsumePointOpenReq *ReverseConsumePointOpenReq) error {
    r.paramReverseConsumePointOpenReq = paramReverseConsumePointOpenReq
    r.Set("param_reverse_consume_point_open_req", paramReverseConsumePointOpenReq)
    return nil
}

// ParamReverseConsumePointOpenReq Getter
func (r AlibabaAlscCrmPointReversepointRequest) GetParamReverseConsumePointOpenReq() *ReverseConsumePointOpenReq {
    return r.paramReverseConsumePointOpenReq
}
