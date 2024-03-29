package tmallcar

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallCarLeaseExceptionflowsynchronizeAPIRequest 天猫开新车租后异常流线下处理状态通知接口 API请求
// tmall.car.lease.exceptionflowsynchronize
//
// 天猫开新车租后异常流线下处理状态通知接口
type TmallCarLeaseExceptionflowsynchronizeAPIRequest struct {
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

// NewTmallCarLeaseExceptionflowsynchronizeRequest 初始化TmallCarLeaseExceptionflowsynchronizeAPIRequest对象
func NewTmallCarLeaseExceptionflowsynchronizeRequest() *TmallCarLeaseExceptionflowsynchronizeAPIRequest {
	return &TmallCarLeaseExceptionflowsynchronizeAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallCarLeaseExceptionflowsynchronizeAPIRequest) Reset() {
	r._desc = ""
	r._orderId = 0
	r._status = 0
	r._flowType = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallCarLeaseExceptionflowsynchronizeAPIRequest) GetApiMethodName() string {
	return "tmall.car.lease.exceptionflowsynchronize"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallCarLeaseExceptionflowsynchronizeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallCarLeaseExceptionflowsynchronizeAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTmallCarLeaseExceptionflowsynchronizeAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallCarLeaseExceptionflowsynchronizeRequest()
	},
}

// GetTmallCarLeaseExceptionflowsynchronizeRequest 从 sync.Pool 获取 TmallCarLeaseExceptionflowsynchronizeAPIRequest
func GetTmallCarLeaseExceptionflowsynchronizeAPIRequest() *TmallCarLeaseExceptionflowsynchronizeAPIRequest {
	return poolTmallCarLeaseExceptionflowsynchronizeAPIRequest.Get().(*TmallCarLeaseExceptionflowsynchronizeAPIRequest)
}

// ReleaseTmallCarLeaseExceptionflowsynchronizeAPIRequest 将 TmallCarLeaseExceptionflowsynchronizeAPIRequest 放入 sync.Pool
func ReleaseTmallCarLeaseExceptionflowsynchronizeAPIRequest(v *TmallCarLeaseExceptionflowsynchronizeAPIRequest) {
	v.Reset()
	poolTmallCarLeaseExceptionflowsynchronizeAPIRequest.Put(v)
}
