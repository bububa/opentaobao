package traveltrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
订单服务详情模版查询 APIRequest
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

func NewAlitripTravelTradeTemplateQueryRequest() *AlitripTravelTradeTemplateQueryRequest{
    return &AlitripTravelTradeTemplateQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripTravelTradeTemplateQueryRequest) GetApiMethodName() string {
    return "alitrip.travel.trade.template.query"
}

func (r AlitripTravelTradeTemplateQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripTravelTradeTemplateQueryRequest) SetIsNew(isNew bool) error {
    r.isNew = isNew
    r.Set("is_new", isNew)
    return nil
}

func (r AlitripTravelTradeTemplateQueryRequest) GetIsNew() bool {
    return r.isNew
}

func (r *AlitripTravelTradeTemplateQueryRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r AlitripTravelTradeTemplateQueryRequest) GetOrderId() int64 {
    return r.orderId
}

