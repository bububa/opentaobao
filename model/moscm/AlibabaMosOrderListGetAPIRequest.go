package moscm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamosorderlistgetAPIRequest 批量查询订单交易 API请求
// alibaba.mos.order.list.get
//
// 批量查询交易信息
type AlibabamosorderlistgetAPIRequest struct {
	model.Params
	// 订单查询条件
	_orderCriteria *OrderCriteria
	// 分页信息
	_paginator *Paginator
}

// NewAlibabamosorderlistgetRequest 初始化AlibabamosorderlistgetAPIRequest对象
func NewAlibabamosorderlistgetRequest() *AlibabamosorderlistgetAPIRequest {
	return &AlibabamosorderlistgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamosorderlistgetAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.order.list.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamosorderlistgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamosorderlistgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderCriteria is OrderCriteria Setter
// 订单查询条件
func (r *AlibabamosorderlistgetAPIRequest) SetOrderCriteria(_orderCriteria *OrderCriteria) error {
	r._orderCriteria = _orderCriteria
	r.Set("order_criteria", _orderCriteria)
	return nil
}

// GetOrderCriteria OrderCriteria Getter
func (r AlibabamosorderlistgetAPIRequest) GetOrderCriteria() *OrderCriteria {
	return r._orderCriteria
}

// SetPaginator is Paginator Setter
// 分页信息
func (r *AlibabamosorderlistgetAPIRequest) SetPaginator(_paginator *Paginator) error {
	r._paginator = _paginator
	r.Set("paginator", _paginator)
	return nil
}

// GetPaginator Paginator Getter
func (r AlibabamosorderlistgetAPIRequest) GetPaginator() *Paginator {
	return r._paginator
}
