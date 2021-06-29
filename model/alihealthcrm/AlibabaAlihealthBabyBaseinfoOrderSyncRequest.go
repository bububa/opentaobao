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
    _tpUserId   int64
    // 商品id
    _commodityId   string
    // 商品名称
    _commodityName   string
    // 价钱
    _amount   *BigDecimal
    // 状态，1是已支付，2是已退款
    _status   int64
    // 订单时间
    _orderTime   string
    // 订单id
    _orderId   string
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
func (r *AlibabaAlihealthBabyBaseinfoOrderSyncRequest) SetTpUserId(_tpUserId int64) error {
    r._tpUserId = _tpUserId
    r.Set("tp_user_id", _tpUserId)
    return nil
}

// TpUserId Getter
func (r AlibabaAlihealthBabyBaseinfoOrderSyncRequest) GetTpUserId() int64 {
    return r._tpUserId
}
// CommodityId Setter
// 商品id
func (r *AlibabaAlihealthBabyBaseinfoOrderSyncRequest) SetCommodityId(_commodityId string) error {
    r._commodityId = _commodityId
    r.Set("commodity_id", _commodityId)
    return nil
}

// CommodityId Getter
func (r AlibabaAlihealthBabyBaseinfoOrderSyncRequest) GetCommodityId() string {
    return r._commodityId
}
// CommodityName Setter
// 商品名称
func (r *AlibabaAlihealthBabyBaseinfoOrderSyncRequest) SetCommodityName(_commodityName string) error {
    r._commodityName = _commodityName
    r.Set("commodity_name", _commodityName)
    return nil
}

// CommodityName Getter
func (r AlibabaAlihealthBabyBaseinfoOrderSyncRequest) GetCommodityName() string {
    return r._commodityName
}
// Amount Setter
// 价钱
func (r *AlibabaAlihealthBabyBaseinfoOrderSyncRequest) SetAmount(_amount *BigDecimal) error {
    r._amount = _amount
    r.Set("amount", _amount)
    return nil
}

// Amount Getter
func (r AlibabaAlihealthBabyBaseinfoOrderSyncRequest) GetAmount() *BigDecimal {
    return r._amount
}
// Status Setter
// 状态，1是已支付，2是已退款
func (r *AlibabaAlihealthBabyBaseinfoOrderSyncRequest) SetStatus(_status int64) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r AlibabaAlihealthBabyBaseinfoOrderSyncRequest) GetStatus() int64 {
    return r._status
}
// OrderTime Setter
// 订单时间
func (r *AlibabaAlihealthBabyBaseinfoOrderSyncRequest) SetOrderTime(_orderTime string) error {
    r._orderTime = _orderTime
    r.Set("order_time", _orderTime)
    return nil
}

// OrderTime Getter
func (r AlibabaAlihealthBabyBaseinfoOrderSyncRequest) GetOrderTime() string {
    return r._orderTime
}
// OrderId Setter
// 订单id
func (r *AlibabaAlihealthBabyBaseinfoOrderSyncRequest) SetOrderId(_orderId string) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r AlibabaAlihealthBabyBaseinfoOrderSyncRequest) GetOrderId() string {
    return r._orderId
}
