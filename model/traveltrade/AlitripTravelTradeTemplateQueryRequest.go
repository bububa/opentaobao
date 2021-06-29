package traveltrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
订单服务详情模版查询 API请求
alitrip.travel.trade.template.query

通过订单ID获取标注模版信息，商家可以根据模版来填充行业信息
*/
type AlitripTravelTradeTemplateQueryRequest struct {
    model.Params
    // 是否取最新的模版
    _isNew   bool
    // 淘宝平台订单ID
    _orderId   int64
}

// 初始化AlitripTravelTradeTemplateQueryRequest对象
func NewAlitripTravelTradeTemplateQueryRequest() *AlitripTravelTradeTemplateQueryRequest{
    return &AlitripTravelTradeTemplateQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripTravelTradeTemplateQueryRequest) GetApiMethodName() string {
    return "alitrip.travel.trade.template.query"
}

// IRequest interface 方法, 获取API参数
func (r AlitripTravelTradeTemplateQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// IsNew Setter
// 是否取最新的模版
func (r *AlitripTravelTradeTemplateQueryRequest) SetIsNew(_isNew bool) error {
    r._isNew = _isNew
    r.Set("is_new", _isNew)
    return nil
}

// IsNew Getter
func (r AlitripTravelTradeTemplateQueryRequest) GetIsNew() bool {
    return r._isNew
}
// OrderId Setter
// 淘宝平台订单ID
func (r *AlitripTravelTradeTemplateQueryRequest) SetOrderId(_orderId int64) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r AlitripTravelTradeTemplateQueryRequest) GetOrderId() int64 {
    return r._orderId
}
