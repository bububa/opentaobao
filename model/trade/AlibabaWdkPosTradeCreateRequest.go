package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
轻pos品牌营销下单接口 API请求
alibaba.wdk.pos.trade.create

提供给石基进行轻pos品牌营销下单
*/
type AlibabaWdkPosTradeCreateRequest struct {
    model.Params
    // 下单请求
    _createRequest   *FastBuyPosCreateRequest
}

// 初始化AlibabaWdkPosTradeCreateRequest对象
func NewAlibabaWdkPosTradeCreateRequest() *AlibabaWdkPosTradeCreateRequest{
    return &AlibabaWdkPosTradeCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkPosTradeCreateRequest) GetApiMethodName() string {
    return "alibaba.wdk.pos.trade.create"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkPosTradeCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CreateRequest Setter
// 下单请求
func (r *AlibabaWdkPosTradeCreateRequest) SetCreateRequest(_createRequest *FastBuyPosCreateRequest) error {
    r._createRequest = _createRequest
    r.Set("create_request", _createRequest)
    return nil
}

// CreateRequest Getter
func (r AlibabaWdkPosTradeCreateRequest) GetCreateRequest() *FastBuyPosCreateRequest {
    return r._createRequest
}
