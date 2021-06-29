package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
资产核销接口 API请求
alibaba.alsc.crm.open.assert.verify

核销储值，积分，券资产
*/
type AlibabaAlscCrmOpenAssertVerifyRequest struct {
    model.Params
    // 入参
    paramPropertyVerifyOpenReq   *PropertyVerifyOpenReq
}

// 初始化AlibabaAlscCrmOpenAssertVerifyRequest对象
func NewAlibabaAlscCrmOpenAssertVerifyRequest() *AlibabaAlscCrmOpenAssertVerifyRequest{
    return &AlibabaAlscCrmOpenAssertVerifyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmOpenAssertVerifyRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.open.assert.verify"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmOpenAssertVerifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamPropertyVerifyOpenReq Setter
// 入参
func (r *AlibabaAlscCrmOpenAssertVerifyRequest) SetParamPropertyVerifyOpenReq(paramPropertyVerifyOpenReq *PropertyVerifyOpenReq) error {
    r.paramPropertyVerifyOpenReq = paramPropertyVerifyOpenReq
    r.Set("param_property_verify_open_req", paramPropertyVerifyOpenReq)
    return nil
}

// ParamPropertyVerifyOpenReq Getter
func (r AlibabaAlscCrmOpenAssertVerifyRequest) GetParamPropertyVerifyOpenReq() *PropertyVerifyOpenReq {
    return r.paramPropertyVerifyOpenReq
}
