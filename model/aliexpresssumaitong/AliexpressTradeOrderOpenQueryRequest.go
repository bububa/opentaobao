package aliexpresssumaitong

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
Aliexpress开放平台订单查询 APIRequest
aliexpress.trade.order.open.query

Aliexpress开放平台订单信息查询
*/
type AliexpressTradeOrderOpenQueryRequest struct {
    model.Params

    // 买家用户id
    buyerId   int64 

    // 订单号
    orderIds   []int64 

    // 外部订单号
    outIds   []string 

    // 小程序appId
    openAppKey   string 

    // 业务编码
    bizCode   string 

}

func NewAliexpressTradeOrderOpenQueryRequest() *AliexpressTradeOrderOpenQueryRequest{
    return &AliexpressTradeOrderOpenQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AliexpressTradeOrderOpenQueryRequest) GetApiMethodName() string {
    return "aliexpress.trade.order.open.query"
}

func (r AliexpressTradeOrderOpenQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliexpressTradeOrderOpenQueryRequest) SetBuyerId(buyerId int64) error {
    r.buyerId = buyerId
    r.Set("buyer_id", buyerId)
    return nil
}

func (r AliexpressTradeOrderOpenQueryRequest) GetBuyerId() int64 {
    return r.buyerId
}

func (r *AliexpressTradeOrderOpenQueryRequest) SetOrderIds(orderIds []int64) error {
    r.orderIds = orderIds
    r.Set("order_ids", orderIds)
    return nil
}

func (r AliexpressTradeOrderOpenQueryRequest) GetOrderIds() []int64 {
    return r.orderIds
}

func (r *AliexpressTradeOrderOpenQueryRequest) SetOutIds(outIds []string) error {
    r.outIds = outIds
    r.Set("out_ids", outIds)
    return nil
}

func (r AliexpressTradeOrderOpenQueryRequest) GetOutIds() []string {
    return r.outIds
}

func (r *AliexpressTradeOrderOpenQueryRequest) SetOpenAppKey(openAppKey string) error {
    r.openAppKey = openAppKey
    r.Set("open_app_key", openAppKey)
    return nil
}

func (r AliexpressTradeOrderOpenQueryRequest) GetOpenAppKey() string {
    return r.openAppKey
}

func (r *AliexpressTradeOrderOpenQueryRequest) SetBizCode(bizCode string) error {
    r.bizCode = bizCode
    r.Set("biz_code", bizCode)
    return nil
}

func (r AliexpressTradeOrderOpenQueryRequest) GetBizCode() string {
    return r.bizCode
}

