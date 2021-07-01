package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
品牌营销的订单活动信息查询 API请求
alibaba.wdk.bm.trade.activity.query

品牌营销的订单活动信息查询
*/
type AlibabaWdkBmTradeActivityQueryAPIRequest struct {
    model.Params
    // 入参
    _queryParam   *IsvOrderQueryParam
}

// 初始化AlibabaWdkBmTradeActivityQueryAPIRequest对象
func NewAlibabaWdkBmTradeActivityQueryRequest() *AlibabaWdkBmTradeActivityQueryAPIRequest{
    return &AlibabaWdkBmTradeActivityQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkBmTradeActivityQueryAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.bm.trade.activity.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkBmTradeActivityQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// QueryParam Setter
// 入参
func (r *AlibabaWdkBmTradeActivityQueryAPIRequest) SetQueryParam(_queryParam *IsvOrderQueryParam) error {
    r._queryParam = _queryParam
    r.Set("query_param", _queryParam)
    return nil
}

// QueryParam Getter
func (r AlibabaWdkBmTradeActivityQueryAPIRequest) GetQueryParam() *IsvOrderQueryParam {
    return r._queryParam
}
