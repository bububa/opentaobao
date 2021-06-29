package fundplatform

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
异步批量生成储值卡 API请求
alibaba.fundplatform.cardorder.fetch.card.async

外部业务方异步批量生成储值卡的接口。同步只返回接受成功，异步会通知制卡成功
*/
type AlibabaFundplatformCardorderFetchCardAsyncRequest struct {
    model.Params
    // 入参复杂对象
    _paramCardFetchAsyncRequest   *CardFetchAsyncRequest
}

// 初始化AlibabaFundplatformCardorderFetchCardAsyncRequest对象
func NewAlibabaFundplatformCardorderFetchCardAsyncRequest() *AlibabaFundplatformCardorderFetchCardAsyncRequest{
    return &AlibabaFundplatformCardorderFetchCardAsyncRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaFundplatformCardorderFetchCardAsyncRequest) GetApiMethodName() string {
    return "alibaba.fundplatform.cardorder.fetch.card.async"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaFundplatformCardorderFetchCardAsyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamCardFetchAsyncRequest Setter
// 入参复杂对象
func (r *AlibabaFundplatformCardorderFetchCardAsyncRequest) SetParamCardFetchAsyncRequest(_paramCardFetchAsyncRequest *CardFetchAsyncRequest) error {
    r._paramCardFetchAsyncRequest = _paramCardFetchAsyncRequest
    r.Set("param_card_fetch_async_request", _paramCardFetchAsyncRequest)
    return nil
}

// ParamCardFetchAsyncRequest Getter
func (r AlibabaFundplatformCardorderFetchCardAsyncRequest) GetParamCardFetchAsyncRequest() *CardFetchAsyncRequest {
    return r._paramCardFetchAsyncRequest
}
