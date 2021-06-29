package fundplatform

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
异步批量生成储值卡 APIRequest
alibaba.fundplatform.cardorder.fetch.card.async

外部业务方异步批量生成储值卡的接口。同步只返回接受成功，异步会通知制卡成功
*/
type AlibabaFundplatformCardorderFetchCardAsyncRequest struct {
    model.Params

    // 入参复杂对象
    paramCardFetchAsyncRequest   *CardFetchAsyncRequest 

}

func NewAlibabaFundplatformCardorderFetchCardAsyncRequest() *AlibabaFundplatformCardorderFetchCardAsyncRequest{
    return &AlibabaFundplatformCardorderFetchCardAsyncRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaFundplatformCardorderFetchCardAsyncRequest) GetApiMethodName() string {
    return "alibaba.fundplatform.cardorder.fetch.card.async"
}

func (r AlibabaFundplatformCardorderFetchCardAsyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaFundplatformCardorderFetchCardAsyncRequest) SetParamCardFetchAsyncRequest(paramCardFetchAsyncRequest *CardFetchAsyncRequest) error {
    r.paramCardFetchAsyncRequest = paramCardFetchAsyncRequest
    r.Set("param_card_fetch_async_request", paramCardFetchAsyncRequest)
    return nil
}

func (r AlibabaFundplatformCardorderFetchCardAsyncRequest) GetParamCardFetchAsyncRequest() *CardFetchAsyncRequest {
    return r.paramCardFetchAsyncRequest
}

