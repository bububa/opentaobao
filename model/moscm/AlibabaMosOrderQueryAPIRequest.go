package moscm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamosorderqueryAPIRequest 批量查询订单信息 API请求
// alibaba.mos.order.query
//
// 查询多笔交易信息
type AlibabamosorderqueryAPIRequest struct {
	model.Params
	// 订单查询
	_orderCriteria *OrderCriteria
	// 分页信息
	_paginator *Paginator
}

// NewAlibabamosorderqueryRequest 初始化AlibabamosorderqueryAPIRequest对象
func NewAlibabamosorderqueryRequest() *AlibabamosorderqueryAPIRequest {
	return &AlibabamosorderqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamosorderqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.order.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamosorderqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamosorderqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderCriteria is OrderCriteria Setter
// 订单查询
func (r *AlibabamosorderqueryAPIRequest) SetOrderCriteria(_orderCriteria *OrderCriteria) error {
	r._orderCriteria = _orderCriteria
	r.Set("order_criteria", _orderCriteria)
	return nil
}

// GetOrderCriteria OrderCriteria Getter
func (r AlibabamosorderqueryAPIRequest) GetOrderCriteria() *OrderCriteria {
	return r._orderCriteria
}

// SetPaginator is Paginator Setter
// 分页信息
func (r *AlibabamosorderqueryAPIRequest) SetPaginator(_paginator *Paginator) error {
	r._paginator = _paginator
	r.Set("paginator", _paginator)
	return nil
}

// GetPaginator Paginator Getter
func (r AlibabamosorderqueryAPIRequest) GetPaginator() *Paginator {
	return r._paginator
}
