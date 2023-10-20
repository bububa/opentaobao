package tmallcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallcarleaseexceptionflowsynchronizeAPIRequest 天猫开新车租后异常流线下处理状态通知接口 API请求
// tmall.car.lease.exceptionflowsynchronize
//
// 天猫开新车租后异常流线下处理状态通知接口
type TmallcarleaseexceptionflowsynchronizeAPIRequest struct {
	model.Params
	// 切换原因描述
	_desc string
	// 天猫开新车订单id
	_orderId int64
	// 1:开始切换为异常流，2:线下处理完成
	_status int64
	// 异常流类型,0.退车,1.买断,2.分期，3.续租
	_flowType int64
}

// NewTmallcarleaseexceptionflowsynchronizeRequest 初始化TmallcarleaseexceptionflowsynchronizeAPIRequest对象
func NewTmallcarleaseexceptionflowsynchronizeRequest() *TmallcarleaseexceptionflowsynchronizeAPIRequest {
	return &TmallcarleaseexceptionflowsynchronizeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallcarleaseexceptionflowsynchronizeAPIRequest) GetApiMethodName() string {
	return "tmall.car.lease.exceptionflowsynchronize"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallcarleaseexceptionflowsynchronizeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallcarleaseexceptionflowsynchronizeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDesc is Desc Setter
// 切换原因描述
func (r *TmallcarleaseexceptionflowsynchronizeAPIRequest) SetDesc(_desc string) error {
	r._desc = _desc
	r.Set("desc", _desc)
	return nil
}

// GetDesc Desc Getter
func (r TmallcarleaseexceptionflowsynchronizeAPIRequest) GetDesc() string {
	return r._desc
}

// SetOrderId is OrderId Setter
// 天猫开新车订单id
func (r *TmallcarleaseexceptionflowsynchronizeAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TmallcarleaseexceptionflowsynchronizeAPIRequest) GetOrderId() int64 {
	return r._orderId
}

// SetStatus is Status Setter
// 1:开始切换为异常流，2:线下处理完成
func (r *TmallcarleaseexceptionflowsynchronizeAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TmallcarleaseexceptionflowsynchronizeAPIRequest) GetStatus() int64 {
	return r._status
}

// SetFlowType is FlowType Setter
// 异常流类型,0.退车,1.买断,2.分期，3.续租
func (r *TmallcarleaseexceptionflowsynchronizeAPIRequest) SetFlowType(_flowType int64) error {
	r._flowType = _flowType
	r.Set("flow_type", _flowType)
	return nil
}

// GetFlowType FlowType Getter
func (r TmallcarleaseexceptionflowsynchronizeAPIRequest) GetFlowType() int64 {
	return r._flowType
}
