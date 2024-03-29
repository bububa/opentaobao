package tmallcar

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallCarLeaseStatussynchronizeAPIRequest 天猫开新车租后状态同步 API请求
// tmall.car.lease.statussynchronize
//
// 天猫开新车租后状态同步
type TmallCarLeaseStatussynchronizeAPIRequest struct {
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

// NewTmallCarLeaseStatussynchronizeRequest 初始化TmallCarLeaseStatussynchronizeAPIRequest对象
func NewTmallCarLeaseStatussynchronizeRequest() *TmallCarLeaseStatussynchronizeAPIRequest {
	return &TmallCarLeaseStatussynchronizeAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallCarLeaseStatussynchronizeAPIRequest) Reset() {
	r._refuseReason = ""
	r._orderId = 0
	r._bizType = 0
	r._actionType = 0
	r._actionValue = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallCarLeaseStatussynchronizeAPIRequest) GetApiMethodName() string {
	return "tmall.car.lease.statussynchronize"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallCarLeaseStatussynchronizeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallCarLeaseStatussynchronizeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefuseReason is RefuseReason Setter
// 拒绝原因
func (r *TmallCarLeaseStatussynchronizeAPIRequest) SetRefuseReason(_refuseReason string) error {
	r._refuseReason = _refuseReason
	r.Set("refuse_reason", _refuseReason)
	return nil
}

// GetRefuseReason RefuseReason Getter
func (r TmallCarLeaseStatussynchronizeAPIRequest) GetRefuseReason() string {
	return r._refuseReason
}

// SetOrderId is OrderId Setter
// 天猫开新车订单号
func (r *TmallCarLeaseStatussynchronizeAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TmallCarLeaseStatussynchronizeAPIRequest) GetOrderId() int64 {
	return r._orderId
}

// SetBizType is BizType Setter
// 业务类型:0.退车,1.买断,2.分期，3.续租
func (r *TmallCarLeaseStatussynchronizeAPIRequest) SetBizType(_bizType int64) error {
	r._bizType = _bizType
	r.Set("biz_type", _bizType)
	return nil
}

// GetBizType BizType Getter
func (r TmallCarLeaseStatussynchronizeAPIRequest) GetBizType() int64 {
	return r._bizType
}

// SetActionType is ActionType Setter
// 1:过户,2:续租,3.额外费用审核，4.退车完成
func (r *TmallCarLeaseStatussynchronizeAPIRequest) SetActionType(_actionType int64) error {
	r._actionType = _actionType
	r.Set("action_type", _actionType)
	return nil
}

// GetActionType ActionType Getter
func (r TmallCarLeaseStatussynchronizeAPIRequest) GetActionType() int64 {
	return r._actionType
}

// SetActionValue is ActionValue Setter
// 1:通过，-1:拒绝
func (r *TmallCarLeaseStatussynchronizeAPIRequest) SetActionValue(_actionValue int64) error {
	r._actionValue = _actionValue
	r.Set("action_value", _actionValue)
	return nil
}

// GetActionValue ActionValue Getter
func (r TmallCarLeaseStatussynchronizeAPIRequest) GetActionValue() int64 {
	return r._actionValue
}

var poolTmallCarLeaseStatussynchronizeAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallCarLeaseStatussynchronizeRequest()
	},
}

// GetTmallCarLeaseStatussynchronizeRequest 从 sync.Pool 获取 TmallCarLeaseStatussynchronizeAPIRequest
func GetTmallCarLeaseStatussynchronizeAPIRequest() *TmallCarLeaseStatussynchronizeAPIRequest {
	return poolTmallCarLeaseStatussynchronizeAPIRequest.Get().(*TmallCarLeaseStatussynchronizeAPIRequest)
}

// ReleaseTmallCarLeaseStatussynchronizeAPIRequest 将 TmallCarLeaseStatussynchronizeAPIRequest 放入 sync.Pool
func ReleaseTmallCarLeaseStatussynchronizeAPIRequest(v *TmallCarLeaseStatussynchronizeAPIRequest) {
	v.Reset()
	poolTmallCarLeaseStatussynchronizeAPIRequest.Put(v)
}
