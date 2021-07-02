package tmallcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallCarLeaseExceptionflowsynchronizeAPIRequest 天猫开新车租后异常流线下处理状态通知接口 API请求
// tmall.car.lease.exceptionflowsynchronize
//
// 天猫开新车租后异常流线下处理状态通知接口
type TmallCarLeaseExceptionflowsynchronizeAPIRequest struct {
	model.Params
	// 天猫开新车订单id
	_orderId int64
	// 1:开始切换为异常流，2:线下处理完成
	_status int64
	// 异常流类型,0.退车,1.买断,2.分期，3.续租
	_flowType int64
	// 切换原因描述
	_desc string
}

// NewTmallCarLeaseExceptionflowsynchronizeRequest 初始化TmallCarLeaseExceptionflowsynchronizeAPIRequest对象
func NewTmallCarLeaseExceptionflowsynchronizeRequest() *TmallCarLeaseExceptionflowsynchronizeAPIRequest {
	return &TmallCarLeaseExceptionflowsynchronizeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallCarLeaseExceptionflowsynchronizeAPIRequest) GetApiMethodName() string {
	return "tmall.car.lease.exceptionflowsynchronize"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallCarLeaseExceptionflowsynchronizeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetOrderId is OrderId Setter
// 天猫开新车订单id
func (r *TmallCarLeaseExceptionflowsynchronizeAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TmallCarLeaseExceptionflowsynchronizeAPIRequest) GetOrderId() int64 {
	return r._orderId
}

// SetStatus is Status Setter
// 1:开始切换为异常流，2:线下处理完成
func (r *TmallCarLeaseExceptionflowsynchronizeAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TmallCarLeaseExceptionflowsynchronizeAPIRequest) GetStatus() int64 {
	return r._status
}

// SetFlowType is FlowType Setter
// 异常流类型,0.退车,1.买断,2.分期，3.续租
func (r *TmallCarLeaseExceptionflowsynchronizeAPIRequest) SetFlowType(_flowType int64) error {
	r._flowType = _flowType
	r.Set("flow_type", _flowType)
	return nil
}

// GetFlowType FlowType Getter
func (r TmallCarLeaseExceptionflowsynchronizeAPIRequest) GetFlowType() int64 {
	return r._flowType
}

// SetDesc is Desc Setter
// 切换原因描述
func (r *TmallCarLeaseExceptionflowsynchronizeAPIRequest) SetDesc(_desc string) error {
	r._desc = _desc
	r.Set("desc", _desc)
	return nil
}

// GetDesc Desc Getter
func (r TmallCarLeaseExceptionflowsynchronizeAPIRequest) GetDesc() string {
	return r._desc
}
