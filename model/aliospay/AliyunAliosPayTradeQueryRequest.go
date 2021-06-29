package aliospay

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询支付结果接口 APIRequest
aliyun.alios.pay.trade.query

商户用来查询支付结果接口
*/
type AliyunAliosPayTradeQueryRequest struct {
    model.Params

    // 请求参数
    queryTradeRequest   *QueryTradeRequest 

}

func NewAliyunAliosPayTradeQueryRequest() *AliyunAliosPayTradeQueryRequest{
    return &AliyunAliosPayTradeQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AliyunAliosPayTradeQueryRequest) GetApiMethodName() string {
    return "aliyun.alios.pay.trade.query"
}

func (r AliyunAliosPayTradeQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliyunAliosPayTradeQueryRequest) SetQueryTradeRequest(queryTradeRequest *QueryTradeRequest) error {
    r.queryTradeRequest = queryTradeRequest
    r.Set("query_trade_request", queryTradeRequest)
    return nil
}

func (r AliyunAliosPayTradeQueryRequest) GetQueryTradeRequest() *QueryTradeRequest {
    return r.queryTradeRequest
}

