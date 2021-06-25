package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
品牌营销的订单活动信息查询 APIRequest
alibaba.wdk.bm.trade.activity.query

品牌营销的订单活动信息查询
*/
type AlibabaWdkBmTradeActivityQueryRequest struct {
    model.Params

    // 入参
    queryParam   *IsvOrderQueryParam 

}

func NewAlibabaWdkBmTradeActivityQueryRequest() *AlibabaWdkBmTradeActivityQueryRequest{
    return &AlibabaWdkBmTradeActivityQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkBmTradeActivityQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.bm.trade.activity.query"
}

func (r AlibabaWdkBmTradeActivityQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkBmTradeActivityQueryRequest) SetQueryParam(queryParam *IsvOrderQueryParam) error {
    r.queryParam = queryParam
    r.Set("query_param", queryParam)
    return nil
}

func (r AlibabaWdkBmTradeActivityQueryRequest) GetQueryParam() *IsvOrderQueryParam {
    return r.queryParam
}

