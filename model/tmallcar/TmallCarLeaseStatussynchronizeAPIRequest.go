package tmallcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallcarleasestatussynchronizeAPIRequest 天猫开新车租后状态同步 API请求
// tmall.car.lease.statussynchronize
//
// 天猫开新车租后状态同步
type TmallcarleasestatussynchronizeAPIRequest struct {
	model.Params
	// 拒绝原因
	_refuseReason string
	// 天猫开新车订单号
	_orderId int64
	// 业务类型:0.退车,1.买断,2.分期，3.续租
	_bizType int64
	// 1:过户,2:续租,3.额外费用审核，4.退车完成
	_actionType int64
	// 1:通过，-1:拒绝
	_actionValue int64
}

// NewTmallcarleasestatussynchronizeRequest 初始化TmallcarleasestatussynchronizeAPIRequest对象
func NewTmallcarleasestatussynchronizeRequest() *TmallcarleasestatussynchronizeAPIRequest {
	return &TmallcarleasestatussynchronizeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallcarleasestatussynchronizeAPIRequest) GetApiMethodName() string {
	return "tmall.car.lease.statussynchronize"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallcarleasestatussynchronizeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallcarleasestatussynchronizeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefuseReason is RefuseReason Setter
// 拒绝原因
func (r *TmallcarleasestatussynchronizeAPIRequest) SetRefuseReason(_refuseReason string) error {
	r._refuseReason = _refuseReason
	r.Set("refuse_reason", _refuseReason)
	return nil
}

// GetRefuseReason RefuseReason Getter
func (r TmallcarleasestatussynchronizeAPIRequest) GetRefuseReason() string {
	return r._refuseReason
}

// SetOrderId is OrderId Setter
// 天猫开新车订单号
func (r *TmallcarleasestatussynchronizeAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TmallcarleasestatussynchronizeAPIRequest) GetOrderId() int64 {
	return r._orderId
}

// SetBizType is BizType Setter
// 业务类型:0.退车,1.买断,2.分期，3.续租
func (r *TmallcarleasestatussynchronizeAPIRequest) SetBizType(_bizType int64) error {
	r._bizType = _bizType
	r.Set("biz_type", _bizType)
	return nil
}

// GetBizType BizType Getter
func (r TmallcarleasestatussynchronizeAPIRequest) GetBizType() int64 {
	return r._bizType
}

// SetActionType is ActionType Setter
// 1:过户,2:续租,3.额外费用审核，4.退车完成
func (r *TmallcarleasestatussynchronizeAPIRequest) SetActionType(_actionType int64) error {
	r._actionType = _actionType
	r.Set("action_type", _actionType)
	return nil
}

// GetActionType ActionType Getter
func (r TmallcarleasestatussynchronizeAPIRequest) GetActionType() int64 {
	return r._actionType
}

// SetActionValue is ActionValue Setter
// 1:通过，-1:拒绝
func (r *TmallcarleasestatussynchronizeAPIRequest) SetActionValue(_actionValue int64) error {
	r._actionValue = _actionValue
	r.Set("action_value", _actionValue)
	return nil
}

// GetActionValue ActionValue Getter
func (r TmallcarleasestatussynchronizeAPIRequest) GetActionValue() int64 {
	return r._actionValue
}
