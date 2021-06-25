package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
轻pos品牌营销查询接口 APIRequest
alibaba.wdk.pos.trade.query

轻pos品牌营销场景，外部商家查询营销信息
*/
type AlibabaWdkPosTradeQueryRequest struct {
    model.Params

    // 查询请求
    queryRequest   *FastBuyPosQueryRequest 

}

func NewAlibabaWdkPosTradeQueryRequest() *AlibabaWdkPosTradeQueryRequest{
    return &AlibabaWdkPosTradeQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkPosTradeQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.pos.trade.query"
}

func (r AlibabaWdkPosTradeQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkPosTradeQueryRequest) SetQueryRequest(queryRequest *FastBuyPosQueryRequest) error {
    r.queryRequest = queryRequest
    r.Set("query_request", queryRequest)
    return nil
}

func (r AlibabaWdkPosTradeQueryRequest) GetQueryRequest() *FastBuyPosQueryRequest {
    return r.queryRequest
}

