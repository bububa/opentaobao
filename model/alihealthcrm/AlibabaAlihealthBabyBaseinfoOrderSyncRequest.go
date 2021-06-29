package alihealthcrm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
alibaba.alihealth.baby.baseinfo.order.sync API请求
alibaba.alihealth.baby.baseinfo.order.sync

育学园将订单信息回传给我们
*/
type AlibabaAlihealthBabyBaseinfoOrderSyncRequest struct {
    model.Params
    // 健康id
    tpUserId   int64
    // 商品id
    commodityId   string
    // 商品名称
    commodityName   string
    // 价钱
    amount   *BigDecimal
    // 状态，1是已支付，2是已退款
    status   int64
    // 订单时间
    orderTime   string
    // 订单id
    orderId   string
}

// 初始化AlibabaAlihealthBabyBaseinfoOrderSyncRequest对象
func NewAlibabaAlihealthBabyBaseinfoOrderSyncRequest() *AlibabaAlihealthBabyBaseinfoOrderSyncRequest{
    return &AlibabaAlihealthBabyBaseinfoOrderSyncRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthBabyBaseinfoOrderSyncRequest) GetApiMethodName() string {
    return "alibaba.alihealth.baby.baseinfo.order.sync"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthBabyBaseinfoOrderSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TpUserId Setter
// 健康id
func (r *AlibabaAlihealthBabyBaseinfoOrderSyncRequest) SetTpUserId(tpUserId int64) error {
    r.tpUserId = tpUserId
    r.Set("tp_user_id", tpUserId)
    return nil
}

// TpUserId Getter
func (r AlibabaAlihealthBabyBaseinfoOrderSyncRequest) GetTpUserId() int64 {
    return r.tpUserId
}
// CommodityId Setter
// 商品id
func (r *AlibabaAlihealthBabyBaseinfoOrderSyncRequest) SetCommodityId(commodityId string) error {
    r.commodityId = commodityId
    r.Set("commodity_id", commodityId)
    return nil
}

// CommodityId Getter
func (r AlibabaAlihealthBabyBaseinfoOrderSyncRequest) GetCommodityId() string {
    return r.commodityId
}
// CommodityName Setter
// 商品名称
func (r *AlibabaAlihealthBabyBaseinfoOrderSyncRequest) SetCommodityName(commodityName string) error {
    r.commodityName = commodityName
    r.Set("commodity_name", commodityName)
    return nil
}

// CommodityName Getter
func (r AlibabaAlihealthBabyBaseinfoOrderSyncRequest) GetCommodityName() string {
    return r.commodityName
}
// Amount Setter
// 价钱
func (r *AlibabaAlihealthBabyBaseinfoOrderSyncRequest) SetAmount(amount *BigDecimal) error {
    r.amount = amount
    r.Set("amount", amount)
    return nil
}

// Amount Getter
func (r AlibabaAlihealthBabyBaseinfoOrderSyncRequest) GetAmount() *BigDecimal {
    return r.amount
}
// Status Setter
// 状态，1是已支付，2是已退款
func (r *AlibabaAlihealthBabyBaseinfoOrderSyncRequest) SetStatus(status int64) error {
    r.status = status
    r.Set("status", status)
    return nil
}

// Status Getter
func (r AlibabaAlihealthBabyBaseinfoOrderSyncRequest) GetStatus() int64 {
    return r.status
}
// OrderTime Setter
// 订单时间
func (r *AlibabaAlihealthBabyBaseinfoOrderSyncRequest) SetOrderTime(orderTime string) error {
    r.orderTime = orderTime
    r.Set("order_time", orderTime)
    return nil
}

// OrderTime Getter
func (r AlibabaAlihealthBabyBaseinfoOrderSyncRequest) GetOrderTime() string {
    return r.orderTime
}
// OrderId Setter
// 订单id
func (r *AlibabaAlihealthBabyBaseinfoOrderSyncRequest) SetOrderId(orderId string) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

// OrderId Getter
func (r AlibabaAlihealthBabyBaseinfoOrderSyncRequest) GetOrderId() string {
    return r.orderId
}
