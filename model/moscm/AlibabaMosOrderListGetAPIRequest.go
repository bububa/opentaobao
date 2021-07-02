package moscm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosOrderListGetAPIRequest 批量查询订单交易 API请求
// alibaba.mos.order.list.get
//
// 批量查询交易信息
type AlibabaMosOrderListGetAPIRequest struct {
	model.Params
	// 订单查询条件
	_orderCriteria *OrderCriteria
	// 分页信息
	_paginator *Paginator
}

// NewAlibabaMosOrderListGetRequest 初始化AlibabaMosOrderListGetAPIRequest对象
func NewAlibabaMosOrderListGetRequest() *AlibabaMosOrderListGetAPIRequest {
	return &AlibabaMosOrderListGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMosOrderListGetAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.order.list.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMosOrderListGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OrderCriteria Setter
// 订单查询条件
func (r *AlibabaMosOrderListGetAPIRequest) SetOrderCriteria(_orderCriteria *OrderCriteria) error {
	r._orderCriteria = _orderCriteria
	r.Set("order_criteria", _orderCriteria)
	return nil
}

// Get OrderCriteria Getter
func (r AlibabaMosOrderListGetAPIRequest) GetOrderCriteria() *OrderCriteria {
	return r._orderCriteria
}

// Set is Paginator Setter
// 分页信息
func (r *AlibabaMosOrderListGetAPIRequest) SetPaginator(_paginator *Paginator) error {
	r._paginator = _paginator
	r.Set("paginator", _paginator)
	return nil
}

// Get Paginator Getter
func (r AlibabaMosOrderListGetAPIRequest) GetPaginator() *Paginator {
	return r._paginator
}
