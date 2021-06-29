package alihealthcrm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
订单数据回传接口 API请求
alibaba.alihealth.alipaypfm.order.sync

订单数据回传接口，各个isv通过我们渠道产生订单需要回传进行统计
*/
type AlibabaAlihealthAlipaypfmOrderSyncRequest struct {
    model.Params
    // user_id
    _userId   int64
    // 订单id
    _orderId   string
    // 订单价格
    _orderPrice   string
    // 订单状态
    _orderStatus   string
    // 扩展参数
    _extParam   string
}

// 初始化AlibabaAlihealthAlipaypfmOrderSyncRequest对象
func NewAlibabaAlihealthAlipaypfmOrderSyncRequest() *AlibabaAlihealthAlipaypfmOrderSyncRequest{
    return &AlibabaAlihealthAlipaypfmOrderSyncRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthAlipaypfmOrderSyncRequest) GetApiMethodName() string {
    return "alibaba.alihealth.alipaypfm.order.sync"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthAlipaypfmOrderSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UserId Setter
// user_id
func (r *AlibabaAlihealthAlipaypfmOrderSyncRequest) SetUserId(_userId int64) error {
    r._userId = _userId
    r.Set("user_id", _userId)
    return nil
}

// UserId Getter
func (r AlibabaAlihealthAlipaypfmOrderSyncRequest) GetUserId() int64 {
    return r._userId
}
// OrderId Setter
// 订单id
func (r *AlibabaAlihealthAlipaypfmOrderSyncRequest) SetOrderId(_orderId string) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r AlibabaAlihealthAlipaypfmOrderSyncRequest) GetOrderId() string {
    return r._orderId
}
// OrderPrice Setter
// 订单价格
func (r *AlibabaAlihealthAlipaypfmOrderSyncRequest) SetOrderPrice(_orderPrice string) error {
    r._orderPrice = _orderPrice
    r.Set("order_price", _orderPrice)
    return nil
}

// OrderPrice Getter
func (r AlibabaAlihealthAlipaypfmOrderSyncRequest) GetOrderPrice() string {
    return r._orderPrice
}
// OrderStatus Setter
// 订单状态
func (r *AlibabaAlihealthAlipaypfmOrderSyncRequest) SetOrderStatus(_orderStatus string) error {
    r._orderStatus = _orderStatus
    r.Set("order_status", _orderStatus)
    return nil
}

// OrderStatus Getter
func (r AlibabaAlihealthAlipaypfmOrderSyncRequest) GetOrderStatus() string {
    return r._orderStatus
}
// ExtParam Setter
// 扩展参数
func (r *AlibabaAlihealthAlipaypfmOrderSyncRequest) SetExtParam(_extParam string) error {
    r._extParam = _extParam
    r.Set("ext_param", _extParam)
    return nil
}

// ExtParam Getter
func (r AlibabaAlihealthAlipaypfmOrderSyncRequest) GetExtParam() string {
    return r._extParam
}
