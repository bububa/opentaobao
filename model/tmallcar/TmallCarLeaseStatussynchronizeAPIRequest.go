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
type TmallCarLeaseStatussynchronizeAPIRequest struct {
    model.Params
    // 天猫开新车订单号
    _orderId   int64
    // 业务类型:0.退车,1.买断,2.分期，3.续租
    _bizType   int64
    // 1:过户,2:续租,3.额外费用审核，4.退车完成
    _actionType   int64
    // 1:通过，-1:拒绝
    _actionValue   int64
    // 拒绝原因
    _refuseReason   string
}

// 初始化TmallCarLeaseStatussynchronizeAPIRequest对象
func NewTmallCarLeaseStatussynchronizeRequest() *TmallCarLeaseStatussynchronizeAPIRequest{
    return &TmallCarLeaseStatussynchronizeAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallCarLeaseStatussynchronizeAPIRequest) GetApiMethodName() string {
    return "tmall.car.lease.statussynchronize"
}

// IRequest interface 方法, 获取API参数
func (r TmallCarLeaseStatussynchronizeAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 天猫开新车订单号
func (r *TmallCarLeaseStatussynchronizeAPIRequest) SetOrderId(_orderId int64) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r TmallCarLeaseStatussynchronizeAPIRequest) GetOrderId() int64 {
    return r._orderId
}
// BizType Setter
// 业务类型:0.退车,1.买断,2.分期，3.续租
func (r *TmallCarLeaseStatussynchronizeAPIRequest) SetBizType(_bizType int64) error {
    r._bizType = _bizType
    r.Set("biz_type", _bizType)
    return nil
}

// BizType Getter
func (r TmallCarLeaseStatussynchronizeAPIRequest) GetBizType() int64 {
    return r._bizType
}
// ActionType Setter
// 1:过户,2:续租,3.额外费用审核，4.退车完成
func (r *TmallCarLeaseStatussynchronizeAPIRequest) SetActionType(_actionType int64) error {
    r._actionType = _actionType
    r.Set("action_type", _actionType)
    return nil
}

// ActionType Getter
func (r TmallCarLeaseStatussynchronizeAPIRequest) GetActionType() int64 {
    return r._actionType
}
// ActionValue Setter
// 1:通过，-1:拒绝
func (r *TmallCarLeaseStatussynchronizeAPIRequest) SetActionValue(_actionValue int64) error {
    r._actionValue = _actionValue
    r.Set("action_value", _actionValue)
    return nil
}

// ActionValue Getter
func (r TmallCarLeaseStatussynchronizeAPIRequest) GetActionValue() int64 {
    return r._actionValue
}
// RefuseReason Setter
// 拒绝原因
func (r *TmallCarLeaseStatussynchronizeAPIRequest) SetRefuseReason(_refuseReason string) error {
    r._refuseReason = _refuseReason
    r.Set("refuse_reason", _refuseReason)
    return nil
}

// RefuseReason Getter
func (r TmallCarLeaseStatussynchronizeAPIRequest) GetRefuseReason() string {
    return r._refuseReason
}
