package moscm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamosorderrefundlistgetAPIRequest 批量查询交易退货信息 API请求
// alibaba.mos.order.refund.list.get
//
// 批量查询多个退货单的退货明细
type AlibabamosorderrefundlistgetAPIRequest struct {
	model.Params
	// 退换货查询条件
	_rmaCriteria *RmaCriteria
	// 分页信息
	_paginator *Paginator
}

// NewAlibabamosorderrefundlistgetRequest 初始化AlibabamosorderrefundlistgetAPIRequest对象
func NewAlibabamosorderrefundlistgetRequest() *AlibabamosorderrefundlistgetAPIRequest {
	return &AlibabamosorderrefundlistgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamosorderrefundlistgetAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.order.refund.list.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamosorderrefundlistgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamosorderrefundlistgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRmaCriteria is RmaCriteria Setter
// 退换货查询条件
func (r *AlibabamosorderrefundlistgetAPIRequest) SetRmaCriteria(_rmaCriteria *RmaCriteria) error {
	r._rmaCriteria = _rmaCriteria
	r.Set("rma_criteria", _rmaCriteria)
	return nil
}

// GetRmaCriteria RmaCriteria Getter
func (r AlibabamosorderrefundlistgetAPIRequest) GetRmaCriteria() *RmaCriteria {
	return r._rmaCriteria
}

// SetPaginator is Paginator Setter
// 分页信息
func (r *AlibabamosorderrefundlistgetAPIRequest) SetPaginator(_paginator *Paginator) error {
	r._paginator = _paginator
	r.Set("paginator", _paginator)
	return nil
}

// GetPaginator Paginator Getter
func (r AlibabamosorderrefundlistgetAPIRequest) GetPaginator() *Paginator {
	return r._paginator
}
