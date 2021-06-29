package traveltrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
飞猪度假-订单发货接口 API请求
alitrip.travel.trade.deliver

航旅度假无需物流普通商品发货接口（不支持二次预约商品），只支持子订单级别发货
*/
type AlitripTravelTradeDeliverRequest struct {
    model.Params
    // 子订单id
    _subOrderId   int64
}

// 初始化AlitripTravelTradeDeliverRequest对象
func NewAlitripTravelTradeDeliverRequest() *AlitripTravelTradeDeliverRequest{
    return &AlitripTravelTradeDeliverRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripTravelTradeDeliverRequest) GetApiMethodName() string {
    return "alitrip.travel.trade.deliver"
}

// IRequest interface 方法, 获取API参数
func (r AlitripTravelTradeDeliverRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SubOrderId Setter
// 子订单id
func (r *AlitripTravelTradeDeliverRequest) SetSubOrderId(_subOrderId int64) error {
    r._subOrderId = _subOrderId
    r.Set("sub_order_id", _subOrderId)
    return nil
}

// SubOrderId Getter
func (r AlitripTravelTradeDeliverRequest) GetSubOrderId() int64 {
    return r._subOrderId
}
