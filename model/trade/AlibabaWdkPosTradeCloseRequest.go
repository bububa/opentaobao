package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
轻pos品牌营销关单接口 API请求
alibaba.wdk.pos.trade.close

轻pos品牌营销场景，提供关单接口给外部商家
*/
type AlibabaWdkPosTradeCloseRequest struct {
    model.Params
    // 关单请求
    closeRequest   *FastBuyPosCloseRequest
}

// 初始化AlibabaWdkPosTradeCloseRequest对象
func NewAlibabaWdkPosTradeCloseRequest() *AlibabaWdkPosTradeCloseRequest{
    return &AlibabaWdkPosTradeCloseRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkPosTradeCloseRequest) GetApiMethodName() string {
    return "alibaba.wdk.pos.trade.close"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkPosTradeCloseRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CloseRequest Setter
// 关单请求
func (r *AlibabaWdkPosTradeCloseRequest) SetCloseRequest(closeRequest *FastBuyPosCloseRequest) error {
    r.closeRequest = closeRequest
    r.Set("close_request", closeRequest)
    return nil
}

// CloseRequest Getter
func (r AlibabaWdkPosTradeCloseRequest) GetCloseRequest() *FastBuyPosCloseRequest {
    return r.closeRequest
}
