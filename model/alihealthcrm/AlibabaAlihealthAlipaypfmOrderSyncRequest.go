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
    userId   int64
    // 订单id
    orderId   string
    // 订单价格
    orderPrice   string
    // 订单状态
    orderStatus   string
    // 扩展参数
    extParam   string
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
func (r *AlibabaAlihealthAlipaypfmOrderSyncRequest) SetUserId(userId int64) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

// UserId Getter
func (r AlibabaAlihealthAlipaypfmOrderSyncRequest) GetUserId() int64 {
    return r.userId
}
// OrderId Setter
// 订单id
func (r *AlibabaAlihealthAlipaypfmOrderSyncRequest) SetOrderId(orderId string) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

// OrderId Getter
func (r AlibabaAlihealthAlipaypfmOrderSyncRequest) GetOrderId() string {
    return r.orderId
}
// OrderPrice Setter
// 订单价格
func (r *AlibabaAlihealthAlipaypfmOrderSyncRequest) SetOrderPrice(orderPrice string) error {
    r.orderPrice = orderPrice
    r.Set("order_price", orderPrice)
    return nil
}

// OrderPrice Getter
func (r AlibabaAlihealthAlipaypfmOrderSyncRequest) GetOrderPrice() string {
    return r.orderPrice
}
// OrderStatus Setter
// 订单状态
func (r *AlibabaAlihealthAlipaypfmOrderSyncRequest) SetOrderStatus(orderStatus string) error {
    r.orderStatus = orderStatus
    r.Set("order_status", orderStatus)
    return nil
}

// OrderStatus Getter
func (r AlibabaAlihealthAlipaypfmOrderSyncRequest) GetOrderStatus() string {
    return r.orderStatus
}
// ExtParam Setter
// 扩展参数
func (r *AlibabaAlihealthAlipaypfmOrderSyncRequest) SetExtParam(extParam string) error {
    r.extParam = extParam
    r.Set("ext_param", extParam)
    return nil
}

// ExtParam Getter
func (r AlibabaAlihealthAlipaypfmOrderSyncRequest) GetExtParam() string {
    return r.extParam
}
