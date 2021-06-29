package traveltrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
飞猪度假-订单详情查询接口 APIRequest
alitrip.travel.trade.query

飞猪度假订单详情查询接口
*/
type AlitripTravelTradeQueryRequest struct {
    model.Params

    // 主订单id
    orderId   int64 

}

func NewAlitripTravelTradeQueryRequest() *AlitripTravelTradeQueryRequest{
    return &AlitripTravelTradeQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripTravelTradeQueryRequest) GetApiMethodName() string {
    return "alitrip.travel.trade.query"
}

func (r AlitripTravelTradeQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripTravelTradeQueryRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r AlitripTravelTradeQueryRequest) GetOrderId() int64 {
    return r.orderId
}

