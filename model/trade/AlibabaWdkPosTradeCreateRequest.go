package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
轻pos品牌营销下单接口 APIRequest
alibaba.wdk.pos.trade.create

提供给石基进行轻pos品牌营销下单
*/
type AlibabaWdkPosTradeCreateRequest struct {
    model.Params

    // 下单请求
    createRequest   *FastBuyPosCreateRequest 

}

func NewAlibabaWdkPosTradeCreateRequest() *AlibabaWdkPosTradeCreateRequest{
    return &AlibabaWdkPosTradeCreateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkPosTradeCreateRequest) GetApiMethodName() string {
    return "alibaba.wdk.pos.trade.create"
}

func (r AlibabaWdkPosTradeCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkPosTradeCreateRequest) SetCreateRequest(createRequest *FastBuyPosCreateRequest) error {
    r.createRequest = createRequest
    r.Set("create_request", createRequest)
    return nil
}

func (r AlibabaWdkPosTradeCreateRequest) GetCreateRequest() *FastBuyPosCreateRequest {
    return r.createRequest
}

