package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
资产核销回退接口 API请求
alibaba.alsc.crm.open.assert.refund

回退已经核销的储值积分券资产
*/
type AlibabaAlscCrmOpenAssertRefundRequest struct {
    model.Params
    // 入参
    paramPropertyRefundOpenReq   *PropertyRefundOpenReq
}

// 初始化AlibabaAlscCrmOpenAssertRefundRequest对象
func NewAlibabaAlscCrmOpenAssertRefundRequest() *AlibabaAlscCrmOpenAssertRefundRequest{
    return &AlibabaAlscCrmOpenAssertRefundRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmOpenAssertRefundRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.open.assert.refund"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmOpenAssertRefundRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamPropertyRefundOpenReq Setter
// 入参
func (r *AlibabaAlscCrmOpenAssertRefundRequest) SetParamPropertyRefundOpenReq(paramPropertyRefundOpenReq *PropertyRefundOpenReq) error {
    r.paramPropertyRefundOpenReq = paramPropertyRefundOpenReq
    r.Set("param_property_refund_open_req", paramPropertyRefundOpenReq)
    return nil
}

// ParamPropertyRefundOpenReq Getter
func (r AlibabaAlscCrmOpenAssertRefundRequest) GetParamPropertyRefundOpenReq() *PropertyRefundOpenReq {
    return r.paramPropertyRefundOpenReq
}
