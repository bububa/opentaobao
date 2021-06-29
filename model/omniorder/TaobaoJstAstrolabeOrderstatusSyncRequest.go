package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
线下门店派单以及单据相关操作接口 API请求
taobao.jst.astrolabe.orderstatus.sync

针对ERP系统部署在门店的商家，将派单状态回流到星盘
*/
type TaobaoJstAstrolabeOrderstatusSyncRequest struct {
    model.Params
    // 子订单Id
    subOrderIds   []int64
    // 事件发生时间
    actionTime   string
    // 操作人
    operator   string
    // 业务类型
    type   string
    // 订单状态
    status   string
    // 目标门店的商户中心门店编码
    storeId   int64
    // 交易订单
    parentOrderCode   int64
}

// 初始化TaobaoJstAstrolabeOrderstatusSyncRequest对象
func NewTaobaoJstAstrolabeOrderstatusSyncRequest() *TaobaoJstAstrolabeOrderstatusSyncRequest{
    return &TaobaoJstAstrolabeOrderstatusSyncRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoJstAstrolabeOrderstatusSyncRequest) GetApiMethodName() string {
    return "taobao.jst.astrolabe.orderstatus.sync"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoJstAstrolabeOrderstatusSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SubOrderIds Setter
// 子订单Id
func (r *TaobaoJstAstrolabeOrderstatusSyncRequest) SetSubOrderIds(subOrderIds []int64) error {
    r.subOrderIds = subOrderIds
    r.Set("sub_order_ids", subOrderIds)
    return nil
}

// SubOrderIds Getter
func (r TaobaoJstAstrolabeOrderstatusSyncRequest) GetSubOrderIds() []int64 {
    return r.subOrderIds
}
// ActionTime Setter
// 事件发生时间
func (r *TaobaoJstAstrolabeOrderstatusSyncRequest) SetActionTime(actionTime string) error {
    r.actionTime = actionTime
    r.Set("action_time", actionTime)
    return nil
}

// ActionTime Getter
func (r TaobaoJstAstrolabeOrderstatusSyncRequest) GetActionTime() string {
    return r.actionTime
}
// Operator Setter
// 操作人
func (r *TaobaoJstAstrolabeOrderstatusSyncRequest) SetOperator(operator string) error {
    r.operator = operator
    r.Set("operator", operator)
    return nil
}

// Operator Getter
func (r TaobaoJstAstrolabeOrderstatusSyncRequest) GetOperator() string {
    return r.operator
}
// Type Setter
// 业务类型
func (r *TaobaoJstAstrolabeOrderstatusSyncRequest) SetType(type string) error {
    r.type = type
    r.Set("type", type)
    return nil
}

// Type Getter
func (r TaobaoJstAstrolabeOrderstatusSyncRequest) GetType() string {
    return r.type
}
// Status Setter
// 订单状态
func (r *TaobaoJstAstrolabeOrderstatusSyncRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

// Status Getter
func (r TaobaoJstAstrolabeOrderstatusSyncRequest) GetStatus() string {
    return r.status
}
// StoreId Setter
// 目标门店的商户中心门店编码
func (r *TaobaoJstAstrolabeOrderstatusSyncRequest) SetStoreId(storeId int64) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

// StoreId Getter
func (r TaobaoJstAstrolabeOrderstatusSyncRequest) GetStoreId() int64 {
    return r.storeId
}
// ParentOrderCode Setter
// 交易订单
func (r *TaobaoJstAstrolabeOrderstatusSyncRequest) SetParentOrderCode(parentOrderCode int64) error {
    r.parentOrderCode = parentOrderCode
    r.Set("parent_order_code", parentOrderCode)
    return nil
}

// ParentOrderCode Getter
func (r TaobaoJstAstrolabeOrderstatusSyncRequest) GetParentOrderCode() int64 {
    return r.parentOrderCode
}
