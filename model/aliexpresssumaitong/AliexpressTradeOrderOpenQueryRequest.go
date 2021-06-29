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
    _buyerId   int64
    // 订单号
    _orderIds   []int64
    // 外部订单号
    _outIds   []string
    // 小程序appId
    _openAppKey   string
    // 业务编码
    _bizCode   string
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
func (r *AliexpressTradeOrderOpenQueryRequest) SetBuyerId(_buyerId int64) error {
    r._buyerId = _buyerId
    r.Set("buyer_id", _buyerId)
    return nil
}

// BuyerId Getter
func (r AliexpressTradeOrderOpenQueryRequest) GetBuyerId() int64 {
    return r._buyerId
}
// OrderIds Setter
// 订单号
func (r *AliexpressTradeOrderOpenQueryRequest) SetOrderIds(_orderIds []int64) error {
    r._orderIds = _orderIds
    r.Set("order_ids", _orderIds)
    return nil
}

// OrderIds Getter
func (r AliexpressTradeOrderOpenQueryRequest) GetOrderIds() []int64 {
    return r._orderIds
}
// OutIds Setter
// 外部订单号
func (r *AliexpressTradeOrderOpenQueryRequest) SetOutIds(_outIds []string) error {
    r._outIds = _outIds
    r.Set("out_ids", _outIds)
    return nil
}

// OutIds Getter
func (r AliexpressTradeOrderOpenQueryRequest) GetOutIds() []string {
    return r._outIds
}
// OpenAppKey Setter
// 小程序appId
func (r *AliexpressTradeOrderOpenQueryRequest) SetOpenAppKey(_openAppKey string) error {
    r._openAppKey = _openAppKey
    r.Set("open_app_key", _openAppKey)
    return nil
}

// OpenAppKey Getter
func (r AliexpressTradeOrderOpenQueryRequest) GetOpenAppKey() string {
    return r._openAppKey
}
// BizCode Setter
// 业务编码
func (r *AliexpressTradeOrderOpenQueryRequest) SetBizCode(_bizCode string) error {
    r._bizCode = _bizCode
    r.Set("biz_code", _bizCode)
    return nil
}

// BizCode Getter
func (r AliexpressTradeOrderOpenQueryRequest) GetBizCode() string {
    return r._bizCode
}
