package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
校验支付链路中的积分抵扣是否合法 API请求
alibaba.alsc.crm.point.chkpntbypay

校验支付链路中的积分抵扣是否合法
*/
type AlibabaAlscCrmPointChkpntbypayRequest struct {
    model.Params
    // 入参
    _paramConsumePointByPayOpenReq   *ConsumePointByPayOpenReq
}

// 初始化AlibabaAlscCrmPointChkpntbypayRequest对象
func NewAlibabaAlscCrmPointChkpntbypayRequest() *AlibabaAlscCrmPointChkpntbypayRequest{
    return &AlibabaAlscCrmPointChkpntbypayRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmPointChkpntbypayRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.point.chkpntbypay"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmPointChkpntbypayRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamConsumePointByPayOpenReq Setter
// 入参
func (r *AlibabaAlscCrmPointChkpntbypayRequest) SetParamConsumePointByPayOpenReq(_paramConsumePointByPayOpenReq *ConsumePointByPayOpenReq) error {
    r._paramConsumePointByPayOpenReq = _paramConsumePointByPayOpenReq
    r.Set("param_consume_point_by_pay_open_req", _paramConsumePointByPayOpenReq)
    return nil
}

// ParamConsumePointByPayOpenReq Getter
func (r AlibabaAlscCrmPointChkpntbypayRequest) GetParamConsumePointByPayOpenReq() *ConsumePointByPayOpenReq {
    return r._paramConsumePointByPayOpenReq
}
