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
    _subOrderIds   []int64
    // 事件发生时间
    _actionTime   string
    // 操作人
    _operator   string
    // 业务类型
    _type   string
    // 订单状态
    _status   string
    // 目标门店的商户中心门店编码
    _storeId   int64
    // 交易订单
    _parentOrderCode   int64
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
func (r *TaobaoJstAstrolabeOrderstatusSyncRequest) SetSubOrderIds(_subOrderIds []int64) error {
    r._subOrderIds = _subOrderIds
    r.Set("sub_order_ids", _subOrderIds)
    return nil
}

// SubOrderIds Getter
func (r TaobaoJstAstrolabeOrderstatusSyncRequest) GetSubOrderIds() []int64 {
    return r._subOrderIds
}
// ActionTime Setter
// 事件发生时间
func (r *TaobaoJstAstrolabeOrderstatusSyncRequest) SetActionTime(_actionTime string) error {
    r._actionTime = _actionTime
    r.Set("action_time", _actionTime)
    return nil
}

// ActionTime Getter
func (r TaobaoJstAstrolabeOrderstatusSyncRequest) GetActionTime() string {
    return r._actionTime
}
// Operator Setter
// 操作人
func (r *TaobaoJstAstrolabeOrderstatusSyncRequest) SetOperator(_operator string) error {
    r._operator = _operator
    r.Set("operator", _operator)
    return nil
}

// Operator Getter
func (r TaobaoJstAstrolabeOrderstatusSyncRequest) GetOperator() string {
    return r._operator
}
// Type Setter
// 业务类型
func (r *TaobaoJstAstrolabeOrderstatusSyncRequest) SetType(_type string) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r TaobaoJstAstrolabeOrderstatusSyncRequest) GetType() string {
    return r._type
}
// Status Setter
// 订单状态
func (r *TaobaoJstAstrolabeOrderstatusSyncRequest) SetStatus(_status string) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r TaobaoJstAstrolabeOrderstatusSyncRequest) GetStatus() string {
    return r._status
}
// StoreId Setter
// 目标门店的商户中心门店编码
func (r *TaobaoJstAstrolabeOrderstatusSyncRequest) SetStoreId(_storeId int64) error {
    r._storeId = _storeId
    r.Set("store_id", _storeId)
    return nil
}

// StoreId Getter
func (r TaobaoJstAstrolabeOrderstatusSyncRequest) GetStoreId() int64 {
    return r._storeId
}
// ParentOrderCode Setter
// 交易订单
func (r *TaobaoJstAstrolabeOrderstatusSyncRequest) SetParentOrderCode(_parentOrderCode int64) error {
    r._parentOrderCode = _parentOrderCode
    r.Set("parent_order_code", _parentOrderCode)
    return nil
}

// ParentOrderCode Getter
func (r TaobaoJstAstrolabeOrderstatusSyncRequest) GetParentOrderCode() int64 {
    return r._parentOrderCode
}
