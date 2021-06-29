package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
轻pos品牌营销退款接口 API请求
alibaba.wdk.pos.trade.reverse

轻pos品牌营销场景，商家调用退款接口
*/
type AlibabaWdkPosTradeReverseRequest struct {
    model.Params
    // 退款请求
    reverseRequest   *FastBuyPosReverseRequest
}

// 初始化AlibabaWdkPosTradeReverseRequest对象
func NewAlibabaWdkPosTradeReverseRequest() *AlibabaWdkPosTradeReverseRequest{
    return &AlibabaWdkPosTradeReverseRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkPosTradeReverseRequest) GetApiMethodName() string {
    return "alibaba.wdk.pos.trade.reverse"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkPosTradeReverseRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ReverseRequest Setter
// 退款请求
func (r *AlibabaWdkPosTradeReverseRequest) SetReverseRequest(reverseRequest *FastBuyPosReverseRequest) error {
    r.reverseRequest = reverseRequest
    r.Set("reverse_request", reverseRequest)
    return nil
}

// ReverseRequest Getter
func (r AlibabaWdkPosTradeReverseRequest) GetReverseRequest() *FastBuyPosReverseRequest {
    return r.reverseRequest
}
