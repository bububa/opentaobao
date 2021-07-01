package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
轻pos品牌营销查询接口 API请求
alibaba.wdk.pos.trade.query

轻pos品牌营销场景，外部商家查询营销信息
*/
type AlibabaWdkPosTradeQueryAPIRequest struct {
    model.Params
    // 查询请求
    _queryRequest   *FastBuyPosQueryRequest
}

// 初始化AlibabaWdkPosTradeQueryAPIRequest对象
func NewAlibabaWdkPosTradeQueryRequest() *AlibabaWdkPosTradeQueryAPIRequest{
    return &AlibabaWdkPosTradeQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkPosTradeQueryAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.pos.trade.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkPosTradeQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// QueryRequest Setter
// 查询请求
func (r *AlibabaWdkPosTradeQueryAPIRequest) SetQueryRequest(_queryRequest *FastBuyPosQueryRequest) error {
    r._queryRequest = _queryRequest
    r.Set("query_request", _queryRequest)
    return nil
}

// QueryRequest Getter
func (r AlibabaWdkPosTradeQueryAPIRequest) GetQueryRequest() *FastBuyPosQueryRequest {
    return r._queryRequest
}
