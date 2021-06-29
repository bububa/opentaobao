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
    isNew   bool
    // 淘宝平台订单ID
    orderId   int64
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
func (r *AlitripTravelTradeTemplateQueryRequest) SetIsNew(isNew bool) error {
    r.isNew = isNew
    r.Set("is_new", isNew)
    return nil
}

// IsNew Getter
func (r AlitripTravelTradeTemplateQueryRequest) GetIsNew() bool {
    return r.isNew
}
// OrderId Setter
// 淘宝平台订单ID
func (r *AlitripTravelTradeTemplateQueryRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

// OrderId Getter
func (r AlitripTravelTradeTemplateQueryRequest) GetOrderId() int64 {
    return r.orderId
}
