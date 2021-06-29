package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
积分补扣 API请求
alibaba.alsc.crm.point.extra.consume

积分补扣
*/
type AlibabaAlscCrmPointExtraConsumeRequest struct {
    model.Params
    // 入参
    _paramExtraConsumePointOpenReq   *ExtraConsumePointOpenReq
}

// 初始化AlibabaAlscCrmPointExtraConsumeRequest对象
func NewAlibabaAlscCrmPointExtraConsumeRequest() *AlibabaAlscCrmPointExtraConsumeRequest{
    return &AlibabaAlscCrmPointExtraConsumeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmPointExtraConsumeRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.point.extra.consume"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmPointExtraConsumeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamExtraConsumePointOpenReq Setter
// 入参
func (r *AlibabaAlscCrmPointExtraConsumeRequest) SetParamExtraConsumePointOpenReq(_paramExtraConsumePointOpenReq *ExtraConsumePointOpenReq) error {
    r._paramExtraConsumePointOpenReq = _paramExtraConsumePointOpenReq
    r.Set("param_extra_consume_point_open_req", _paramExtraConsumePointOpenReq)
    return nil
}

// ParamExtraConsumePointOpenReq Getter
func (r AlibabaAlscCrmPointExtraConsumeRequest) GetParamExtraConsumePointOpenReq() *ExtraConsumePointOpenReq {
    return r._paramExtraConsumePointOpenReq
}
