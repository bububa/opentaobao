package tmallcar

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫开新车租后状态同步 API请求
tmall.car.lease.statussynchronize

天猫开新车租后状态同步
*/
type TmallCarLeaseStatussynchronizeRequest struct {
    model.Params
    // 天猫开新车订单号
    orderId   int64
    // 业务类型:0.退车,1.买断,2.分期，3.续租
    bizType   int64
    // 1:过户,2:续租,3.额外费用审核，4.退车完成
    actionType   int64
    // 1:通过，-1:拒绝
    actionValue   int64
    // 拒绝原因
    refuseReason   string
}

// 初始化TmallCarLeaseStatussynchronizeRequest对象
func NewTmallCarLeaseStatussynchronizeRequest() *TmallCarLeaseStatussynchronizeRequest{
    return &TmallCarLeaseStatussynchronizeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallCarLeaseStatussynchronizeRequest) GetApiMethodName() string {
    return "tmall.car.lease.statussynchronize"
}

// IRequest interface 方法, 获取API参数
func (r TmallCarLeaseStatussynchronizeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 天猫开新车订单号
func (r *TmallCarLeaseStatussynchronizeRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

// OrderId Getter
func (r TmallCarLeaseStatussynchronizeRequest) GetOrderId() int64 {
    return r.orderId
}
// BizType Setter
// 业务类型:0.退车,1.买断,2.分期，3.续租
func (r *TmallCarLeaseStatussynchronizeRequest) SetBizType(bizType int64) error {
    r.bizType = bizType
    r.Set("biz_type", bizType)
    return nil
}

// BizType Getter
func (r TmallCarLeaseStatussynchronizeRequest) GetBizType() int64 {
    return r.bizType
}
// ActionType Setter
// 1:过户,2:续租,3.额外费用审核，4.退车完成
func (r *TmallCarLeaseStatussynchronizeRequest) SetActionType(actionType int64) error {
    r.actionType = actionType
    r.Set("action_type", actionType)
    return nil
}

// ActionType Getter
func (r TmallCarLeaseStatussynchronizeRequest) GetActionType() int64 {
    return r.actionType
}
// ActionValue Setter
// 1:通过，-1:拒绝
func (r *TmallCarLeaseStatussynchronizeRequest) SetActionValue(actionValue int64) error {
    r.actionValue = actionValue
    r.Set("action_value", actionValue)
    return nil
}

// ActionValue Getter
func (r TmallCarLeaseStatussynchronizeRequest) GetActionValue() int64 {
    return r.actionValue
}
// RefuseReason Setter
// 拒绝原因
func (r *TmallCarLeaseStatussynchronizeRequest) SetRefuseReason(refuseReason string) error {
    r.refuseReason = refuseReason
    r.Set("refuse_reason", refuseReason)
    return nil
}

// RefuseReason Getter
func (r TmallCarLeaseStatussynchronizeRequest) GetRefuseReason() string {
    return r.refuseReason
}
