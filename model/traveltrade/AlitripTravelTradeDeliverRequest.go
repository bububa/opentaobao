package traveltrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
飞猪度假-订单发货接口 APIRequest
alitrip.travel.trade.deliver

航旅度假无需物流普通商品发货接口（不支持二次预约商品），只支持子订单级别发货
*/
type AlitripTravelTradeDeliverRequest struct {
    model.Params

    // 子订单id
    subOrderId   int64 

}

func NewAlitripTravelTradeDeliverRequest() *AlitripTravelTradeDeliverRequest{
    return &AlitripTravelTradeDeliverRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripTravelTradeDeliverRequest) GetApiMethodName() string {
    return "alitrip.travel.trade.deliver"
}

func (r AlitripTravelTradeDeliverRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripTravelTradeDeliverRequest) SetSubOrderId(subOrderId int64) error {
    r.subOrderId = subOrderId
    r.Set("sub_order_id", subOrderId)
    return nil
}

func (r AlitripTravelTradeDeliverRequest) GetSubOrderId() int64 {
    return r.subOrderId
}

