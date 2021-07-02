package moscm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosOrderQueryAPIRequest 批量查询订单信息 API请求
// alibaba.mos.order.query
//
// 查询多笔交易信息
type AlibabaMosOrderQueryAPIRequest struct {
	model.Params
	// 订单查询
	_orderCriteria *OrderCriteria
	// 分页信息
	_paginator *Paginator
}

// NewAlibabaMosOrderQueryRequest 初始化AlibabaMosOrderQueryAPIRequest对象
func NewAlibabaMosOrderQueryRequest() *AlibabaMosOrderQueryAPIRequest {
	return &AlibabaMosOrderQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMosOrderQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.order.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMosOrderQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetOrderCriteria is OrderCriteria Setter
// 订单查询
func (r *AlibabaMosOrderQueryAPIRequest) SetOrderCriteria(_orderCriteria *OrderCriteria) error {
	r._orderCriteria = _orderCriteria
	r.Set("order_criteria", _orderCriteria)
	return nil
}

// GetOrderCriteria OrderCriteria Getter
func (r AlibabaMosOrderQueryAPIRequest) GetOrderCriteria() *OrderCriteria {
	return r._orderCriteria
}

// SetPaginator is Paginator Setter
// 分页信息
func (r *AlibabaMosOrderQueryAPIRequest) SetPaginator(_paginator *Paginator) error {
	r._paginator = _paginator
	r.Set("paginator", _paginator)
	return nil
}

// GetPaginator Paginator Getter
func (r AlibabaMosOrderQueryAPIRequest) GetPaginator() *Paginator {
	return r._paginator
}
