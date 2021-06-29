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
type AlibabaWdkPosTradeQueryRequest struct {
    model.Params
    // 查询请求
    _queryRequest   *FastBuyPosQueryRequest
}

// 初始化AlibabaWdkPosTradeQueryRequest对象
func NewAlibabaWdkPosTradeQueryRequest() *AlibabaWdkPosTradeQueryRequest{
    return &AlibabaWdkPosTradeQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkPosTradeQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.pos.trade.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkPosTradeQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// QueryRequest Setter
// 查询请求
func (r *AlibabaWdkPosTradeQueryRequest) SetQueryRequest(_queryRequest *FastBuyPosQueryRequest) error {
    r._queryRequest = _queryRequest
    r.Set("query_request", _queryRequest)
    return nil
}

// QueryRequest Getter
func (r AlibabaWdkPosTradeQueryRequest) GetQueryRequest() *FastBuyPosQueryRequest {
    return r._queryRequest
}
