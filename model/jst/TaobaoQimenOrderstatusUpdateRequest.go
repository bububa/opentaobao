package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
订单状态更新接口 APIRequest
taobao.qimen.orderstatus.update

星盘和ISV，可以通过此接口，来更新订单状态。此接口应用于使用阿里星盘分单，且使用商家系统（非阿里掌柜）接单/拒单的模式下更新订单状态。
*/
type TaobaoQimenOrderstatusUpdateRequest struct {
    model.Params

    // 星盘派单号
    allocationCode   string 

    // 淘系子订单号
    orderCodes   []Number 

    // 目标门店的商户中心门店编码
    storeId   int64 

    // 业务类型，（枚举值：FAHUO、ZITI）
    type   string 

    // 订单状态，门店发货包括X_SHOP_ALLOCATION、X_SHOP_DENYX_SHOP_HANDLED、X_SHOP_CANCEL_CONFIRM、X_SHOP_CANCEL_DENIED、X_MATCHED；门店自提包括X_COMMODITY_CONFIRMX_COMMODITY_TRANSER、X_TRANSFER _SUCCESS、X_SHOP_CANCEL_CONFIRM、X_MATCHED、X_SHOP_DENY
    status   string 

    // 操作人
    operator   string 

    // 事件发生时间
    actionTime   string 

}

func NewTaobaoQimenOrderstatusUpdateRequest() *TaobaoQimenOrderstatusUpdateRequest{
    return &TaobaoQimenOrderstatusUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQimenOrderstatusUpdateRequest) GetApiMethodName() string {
    return "taobao.qimen.orderstatus.update"
}

func (r TaobaoQimenOrderstatusUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQimenOrderstatusUpdateRequest) SetAllocationCode(allocationCode string) error {
    r.allocationCode = allocationCode
    r.Set("allocation_code", allocationCode)
    return nil
}

func (r TaobaoQimenOrderstatusUpdateRequest) GetAllocationCode() string {
    return r.allocationCode
}

func (r *TaobaoQimenOrderstatusUpdateRequest) SetOrderCodes(orderCodes []Number) error {
    r.orderCodes = orderCodes
    r.Set("order_codes", orderCodes)
    return nil
}

func (r TaobaoQimenOrderstatusUpdateRequest) GetOrderCodes() []Number {
    return r.orderCodes
}

func (r *TaobaoQimenOrderstatusUpdateRequest) SetStoreId(storeId int64) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

func (r TaobaoQimenOrderstatusUpdateRequest) GetStoreId() int64 {
    return r.storeId
}

func (r *TaobaoQimenOrderstatusUpdateRequest) SetType(type string) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r TaobaoQimenOrderstatusUpdateRequest) GetType() string {
    return r.type
}

func (r *TaobaoQimenOrderstatusUpdateRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r TaobaoQimenOrderstatusUpdateRequest) GetStatus() string {
    return r.status
}

func (r *TaobaoQimenOrderstatusUpdateRequest) SetOperator(operator string) error {
    r.operator = operator
    r.Set("operator", operator)
    return nil
}

func (r TaobaoQimenOrderstatusUpdateRequest) GetOperator() string {
    return r.operator
}

func (r *TaobaoQimenOrderstatusUpdateRequest) SetActionTime(actionTime string) error {
    r.actionTime = actionTime
    r.Set("action_time", actionTime)
    return nil
}

func (r TaobaoQimenOrderstatusUpdateRequest) GetActionTime() string {
    return r.actionTime
}

