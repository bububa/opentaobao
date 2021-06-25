package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
轻pos品牌营销关单接口 APIRequest
alibaba.wdk.pos.trade.close

轻pos品牌营销场景，提供关单接口给外部商家
*/
type AlibabaWdkPosTradeCloseRequest struct {
    model.Params

    // 关单请求
    closeRequest   *FastBuyPosCloseRequest 

}

func NewAlibabaWdkPosTradeCloseRequest() *AlibabaWdkPosTradeCloseRequest{
    return &AlibabaWdkPosTradeCloseRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkPosTradeCloseRequest) GetApiMethodName() string {
    return "alibaba.wdk.pos.trade.close"
}

func (r AlibabaWdkPosTradeCloseRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkPosTradeCloseRequest) SetCloseRequest(closeRequest *FastBuyPosCloseRequest) error {
    r.closeRequest = closeRequest
    r.Set("close_request", closeRequest)
    return nil
}

func (r AlibabaWdkPosTradeCloseRequest) GetCloseRequest() *FastBuyPosCloseRequest {
    return r.closeRequest
}

