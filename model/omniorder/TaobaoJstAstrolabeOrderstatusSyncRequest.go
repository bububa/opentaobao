package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
线下门店派单以及单据相关操作接口 APIRequest
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

func NewTaobaoJstAstrolabeOrderstatusSyncRequest() *TaobaoJstAstrolabeOrderstatusSyncRequest{
    return &TaobaoJstAstrolabeOrderstatusSyncRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoJstAstrolabeOrderstatusSyncRequest) GetApiMethodName() string {
    return "taobao.jst.astrolabe.orderstatus.sync"
}

func (r TaobaoJstAstrolabeOrderstatusSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoJstAstrolabeOrderstatusSyncRequest) SetSubOrderIds(subOrderIds []int64) error {
    r.subOrderIds = subOrderIds
    r.Set("sub_order_ids", subOrderIds)
    return nil
}

func (r TaobaoJstAstrolabeOrderstatusSyncRequest) GetSubOrderIds() []int64 {
    return r.subOrderIds
}

func (r *TaobaoJstAstrolabeOrderstatusSyncRequest) SetActionTime(actionTime string) error {
    r.actionTime = actionTime
    r.Set("action_time", actionTime)
    return nil
}

func (r TaobaoJstAstrolabeOrderstatusSyncRequest) GetActionTime() string {
    return r.actionTime
}

func (r *TaobaoJstAstrolabeOrderstatusSyncRequest) SetOperator(operator string) error {
    r.operator = operator
    r.Set("operator", operator)
    return nil
}

func (r TaobaoJstAstrolabeOrderstatusSyncRequest) GetOperator() string {
    return r.operator
}

func (r *TaobaoJstAstrolabeOrderstatusSyncRequest) SetType(type string) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r TaobaoJstAstrolabeOrderstatusSyncRequest) GetType() string {
    return r.type
}

func (r *TaobaoJstAstrolabeOrderstatusSyncRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r TaobaoJstAstrolabeOrderstatusSyncRequest) GetStatus() string {
    return r.status
}

func (r *TaobaoJstAstrolabeOrderstatusSyncRequest) SetStoreId(storeId int64) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

func (r TaobaoJstAstrolabeOrderstatusSyncRequest) GetStoreId() int64 {
    return r.storeId
}

func (r *TaobaoJstAstrolabeOrderstatusSyncRequest) SetParentOrderCode(parentOrderCode int64) error {
    r.parentOrderCode = parentOrderCode
    r.Set("parent_order_code", parentOrderCode)
    return nil
}

func (r TaobaoJstAstrolabeOrderstatusSyncRequest) GetParentOrderCode() int64 {
    return r.parentOrderCode
}

