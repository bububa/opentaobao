package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
轻pos品牌营销支付接口 API请求
alibaba.wdk.pos.trade.pay

轻pos场景，外部商家支付后调用开放平台把支付信息回传给五道口交易
*/
type AlibabaWdkPosTradePayRequest struct {
    model.Params
    // 支付请求
    _payRequest   *FastBuyPosPayRequest
}

// 初始化AlibabaWdkPosTradePayRequest对象
func NewAlibabaWdkPosTradePayRequest() *AlibabaWdkPosTradePayRequest{
    return &AlibabaWdkPosTradePayRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkPosTradePayRequest) GetApiMethodName() string {
    return "alibaba.wdk.pos.trade.pay"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkPosTradePayRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PayRequest Setter
// 支付请求
func (r *AlibabaWdkPosTradePayRequest) SetPayRequest(_payRequest *FastBuyPosPayRequest) error {
    r._payRequest = _payRequest
    r.Set("pay_request", _payRequest)
    return nil
}

// PayRequest Getter
func (r AlibabaWdkPosTradePayRequest) GetPayRequest() *FastBuyPosPayRequest {
    return r._payRequest
}
