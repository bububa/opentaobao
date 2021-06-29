package aliospay

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询支付结果接口 API请求
aliyun.alios.pay.trade.query

商户用来查询支付结果接口
*/
type AliyunAliosPayTradeQueryRequest struct {
    model.Params
    // 请求参数
    queryTradeRequest   *QueryTradeRequest
}

// 初始化AliyunAliosPayTradeQueryRequest对象
func NewAliyunAliosPayTradeQueryRequest() *AliyunAliosPayTradeQueryRequest{
    return &AliyunAliosPayTradeQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliyunAliosPayTradeQueryRequest) GetApiMethodName() string {
    return "aliyun.alios.pay.trade.query"
}

// IRequest interface 方法, 获取API参数
func (r AliyunAliosPayTradeQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// QueryTradeRequest Setter
// 请求参数
func (r *AliyunAliosPayTradeQueryRequest) SetQueryTradeRequest(queryTradeRequest *QueryTradeRequest) error {
    r.queryTradeRequest = queryTradeRequest
    r.Set("query_trade_request", queryTradeRequest)
    return nil
}

// QueryTradeRequest Getter
func (r AliyunAliosPayTradeQueryRequest) GetQueryTradeRequest() *QueryTradeRequest {
    return r.queryTradeRequest
}
