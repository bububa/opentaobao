package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
轻pos品牌营销支付接口 APIRequest
alibaba.wdk.pos.trade.pay

轻pos场景，外部商家支付后调用开放平台把支付信息回传给五道口交易
*/
type AlibabaWdkPosTradePayRequest struct {
    model.Params

    // 支付请求
    payRequest   *FastBuyPosPayRequest 

}

func NewAlibabaWdkPosTradePayRequest() *AlibabaWdkPosTradePayRequest{
    return &AlibabaWdkPosTradePayRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkPosTradePayRequest) GetApiMethodName() string {
    return "alibaba.wdk.pos.trade.pay"
}

func (r AlibabaWdkPosTradePayRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkPosTradePayRequest) SetPayRequest(payRequest *FastBuyPosPayRequest) error {
    r.payRequest = payRequest
    r.Set("pay_request", payRequest)
    return nil
}

func (r AlibabaWdkPosTradePayRequest) GetPayRequest() *FastBuyPosPayRequest {
    return r.payRequest
}

