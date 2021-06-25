package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
资产核销回退接口 APIRequest
alibaba.alsc.crm.open.assert.refund

回退已经核销的储值积分券资产
*/
type AlibabaAlscCrmOpenAssertRefundRequest struct {
    model.Params

    // 入参
    paramPropertyRefundOpenReq   *PropertyRefundOpenReq 

}

func NewAlibabaAlscCrmOpenAssertRefundRequest() *AlibabaAlscCrmOpenAssertRefundRequest{
    return &AlibabaAlscCrmOpenAssertRefundRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlscCrmOpenAssertRefundRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.open.assert.refund"
}

func (r AlibabaAlscCrmOpenAssertRefundRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlscCrmOpenAssertRefundRequest) SetParamPropertyRefundOpenReq(paramPropertyRefundOpenReq *PropertyRefundOpenReq) error {
    r.paramPropertyRefundOpenReq = paramPropertyRefundOpenReq
    r.Set("param_property_refund_open_req", paramPropertyRefundOpenReq)
    return nil
}

func (r AlibabaAlscCrmOpenAssertRefundRequest) GetParamPropertyRefundOpenReq() *PropertyRefundOpenReq {
    return r.paramPropertyRefundOpenReq
}

