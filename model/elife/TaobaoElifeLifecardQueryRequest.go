package elife

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询交易结果 APIRequest
taobao.elife.lifecard.query

卖家在交易状态不明的情况下, 查询交易结果.
*/
type TaobaoElifeLifecardQueryRequest struct {
    model.Params

    // 入参
    queryRequest   *ConsumeRequest 

}

func NewTaobaoElifeLifecardQueryRequest() *TaobaoElifeLifecardQueryRequest{
    return &TaobaoElifeLifecardQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoElifeLifecardQueryRequest) GetApiMethodName() string {
    return "taobao.elife.lifecard.query"
}

func (r TaobaoElifeLifecardQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoElifeLifecardQueryRequest) SetQueryRequest(queryRequest *ConsumeRequest) error {
    r.queryRequest = queryRequest
    r.Set("query_request", queryRequest)
    return nil
}

func (r TaobaoElifeLifecardQueryRequest) GetQueryRequest() *ConsumeRequest {
    return r.queryRequest
}

