package aliexpresssumaitong

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
Aliexpress开放平台订单查询 API请求
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

// 初始化AliexpressTradeOrderOpenQueryRequest对象
func NewAliexpressTradeOrderOpenQueryRequest() *AliexpressTradeOrderOpenQueryRequest{
    return &AliexpressTradeOrderOpenQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressTradeOrderOpenQueryRequest) GetApiMethodName() string {
    return "aliexpress.trade.order.open.query"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressTradeOrderOpenQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BuyerId Setter
// 买家用户id
func (r *AliexpressTradeOrderOpenQueryRequest) SetBuyerId(buyerId int64) error {
    r.buyerId = buyerId
    r.Set("buyer_id", buyerId)
    return nil
}

// BuyerId Getter
func (r AliexpressTradeOrderOpenQueryRequest) GetBuyerId() int64 {
    return r.buyerId
}
// OrderIds Setter
// 订单号
func (r *AliexpressTradeOrderOpenQueryRequest) SetOrderIds(orderIds []int64) error {
    r.orderIds = orderIds
    r.Set("order_ids", orderIds)
    return nil
}

// OrderIds Getter
func (r AliexpressTradeOrderOpenQueryRequest) GetOrderIds() []int64 {
    return r.orderIds
}
// OutIds Setter
// 外部订单号
func (r *AliexpressTradeOrderOpenQueryRequest) SetOutIds(outIds []string) error {
    r.outIds = outIds
    r.Set("out_ids", outIds)
    return nil
}

// OutIds Getter
func (r AliexpressTradeOrderOpenQueryRequest) GetOutIds() []string {
    return r.outIds
}
// OpenAppKey Setter
// 小程序appId
func (r *AliexpressTradeOrderOpenQueryRequest) SetOpenAppKey(openAppKey string) error {
    r.openAppKey = openAppKey
    r.Set("open_app_key", openAppKey)
    return nil
}

// OpenAppKey Getter
func (r AliexpressTradeOrderOpenQueryRequest) GetOpenAppKey() string {
    return r.openAppKey
}
// BizCode Setter
// 业务编码
func (r *AliexpressTradeOrderOpenQueryRequest) SetBizCode(bizCode string) error {
    r.bizCode = bizCode
    r.Set("biz_code", bizCode)
    return nil
}

// BizCode Getter
func (r AliexpressTradeOrderOpenQueryRequest) GetBizCode() string {
    return r.bizCode
}
