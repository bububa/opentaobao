package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
资产核销接口 APIRequest
alibaba.alsc.crm.open.assert.verify

核销储值，积分，券资产
*/
type AlibabaAlscCrmOpenAssertVerifyRequest struct {
    model.Params

    // 入参
    paramPropertyVerifyOpenReq   *PropertyVerifyOpenReq 

}

func NewAlibabaAlscCrmOpenAssertVerifyRequest() *AlibabaAlscCrmOpenAssertVerifyRequest{
    return &AlibabaAlscCrmOpenAssertVerifyRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlscCrmOpenAssertVerifyRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.open.assert.verify"
}

func (r AlibabaAlscCrmOpenAssertVerifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlscCrmOpenAssertVerifyRequest) SetParamPropertyVerifyOpenReq(paramPropertyVerifyOpenReq *PropertyVerifyOpenReq) error {
    r.paramPropertyVerifyOpenReq = paramPropertyVerifyOpenReq
    r.Set("param_property_verify_open_req", paramPropertyVerifyOpenReq)
    return nil
}

func (r AlibabaAlscCrmOpenAssertVerifyRequest) GetParamPropertyVerifyOpenReq() *PropertyVerifyOpenReq {
    return r.paramPropertyVerifyOpenReq
}

