package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
轻pos品牌营销退款接口 APIRequest
alibaba.wdk.pos.trade.reverse

轻pos品牌营销场景，商家调用退款接口
*/
type AlibabaWdkPosTradeReverseRequest struct {
    model.Params

    // 退款请求
    reverseRequest   *FastBuyPosReverseRequest 

}

func NewAlibabaWdkPosTradeReverseRequest() *AlibabaWdkPosTradeReverseRequest{
    return &AlibabaWdkPosTradeReverseRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkPosTradeReverseRequest) GetApiMethodName() string {
    return "alibaba.wdk.pos.trade.reverse"
}

func (r AlibabaWdkPosTradeReverseRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkPosTradeReverseRequest) SetReverseRequest(reverseRequest *FastBuyPosReverseRequest) error {
    r.reverseRequest = reverseRequest
    r.Set("reverse_request", reverseRequest)
    return nil
}

func (r AlibabaWdkPosTradeReverseRequest) GetReverseRequest() *FastBuyPosReverseRequest {
    return r.reverseRequest
}

