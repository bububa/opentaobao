package alihealthcrm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
alibaba.alihealth.baby.baseinfo.order.sync APIRequest
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

func NewAlibabaAlihealthBabyBaseinfoOrderSyncRequest() *AlibabaAlihealthBabyBaseinfoOrderSyncRequest{
    return &AlibabaAlihealthBabyBaseinfoOrderSyncRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthBabyBaseinfoOrderSyncRequest) GetApiMethodName() string {
    return "alibaba.alihealth.baby.baseinfo.order.sync"
}

func (r AlibabaAlihealthBabyBaseinfoOrderSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthBabyBaseinfoOrderSyncRequest) SetTpUserId(tpUserId int64) error {
    r.tpUserId = tpUserId
    r.Set("tp_user_id", tpUserId)
    return nil
}

func (r AlibabaAlihealthBabyBaseinfoOrderSyncRequest) GetTpUserId() int64 {
    return r.tpUserId
}

func (r *AlibabaAlihealthBabyBaseinfoOrderSyncRequest) SetCommodityId(commodityId string) error {
    r.commodityId = commodityId
    r.Set("commodity_id", commodityId)
    return nil
}

func (r AlibabaAlihealthBabyBaseinfoOrderSyncRequest) GetCommodityId() string {
    return r.commodityId
}

func (r *AlibabaAlihealthBabyBaseinfoOrderSyncRequest) SetCommodityName(commodityName string) error {
    r.commodityName = commodityName
    r.Set("commodity_name", commodityName)
    return nil
}

func (r AlibabaAlihealthBabyBaseinfoOrderSyncRequest) GetCommodityName() string {
    return r.commodityName
}

func (r *AlibabaAlihealthBabyBaseinfoOrderSyncRequest) SetAmount(amount *BigDecimal) error {
    r.amount = amount
    r.Set("amount", amount)
    return nil
}

func (r AlibabaAlihealthBabyBaseinfoOrderSyncRequest) GetAmount() *BigDecimal {
    return r.amount
}

func (r *AlibabaAlihealthBabyBaseinfoOrderSyncRequest) SetStatus(status int64) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r AlibabaAlihealthBabyBaseinfoOrderSyncRequest) GetStatus() int64 {
    return r.status
}

func (r *AlibabaAlihealthBabyBaseinfoOrderSyncRequest) SetOrderTime(orderTime string) error {
    r.orderTime = orderTime
    r.Set("order_time", orderTime)
    return nil
}

func (r AlibabaAlihealthBabyBaseinfoOrderSyncRequest) GetOrderTime() string {
    return r.orderTime
}

func (r *AlibabaAlihealthBabyBaseinfoOrderSyncRequest) SetOrderId(orderId string) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r AlibabaAlihealthBabyBaseinfoOrderSyncRequest) GetOrderId() string {
    return r.orderId
}

