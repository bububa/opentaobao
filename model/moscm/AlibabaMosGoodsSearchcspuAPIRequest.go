package moscm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosGoodsSearchcspuAPIRequest cspu查询 API请求
// alibaba.mos.goods.searchcspu
//
// 商品信息查询（仅用于商品上传数据验证，不能用于商品下载，有限流）
type AlibabaMosGoodsSearchcspuAPIRequest struct {
	model.Params
	// 组合查询对象
	_paramCspuCriteria *CspuCriteria
	// 分页参数
	_paramPaginator *Paginator
}

// NewAlibabaMosGoodsSearchcspuRequest 初始化AlibabaMosGoodsSearchcspuAPIRequest对象
func NewAlibabaMosGoodsSearchcspuRequest() *AlibabaMosGoodsSearchcspuAPIRequest {
	return &AlibabaMosGoodsSearchcspuAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMosGoodsSearchcspuAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.goods.searchcspu"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMosGoodsSearchcspuAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParamCspuCriteria is ParamCspuCriteria Setter
// 组合查询对象
func (r *AlibabaMosGoodsSearchcspuAPIRequest) SetParamCspuCriteria(_paramCspuCriteria *CspuCriteria) error {
	r._paramCspuCriteria = _paramCspuCriteria
	r.Set("param_cspu_criteria", _paramCspuCriteria)
	return nil
}

// GetParamCspuCriteria ParamCspuCriteria Getter
func (r AlibabaMosGoodsSearchcspuAPIRequest) GetParamCspuCriteria() *CspuCriteria {
	return r._paramCspuCriteria
}

// SetParamPaginator is ParamPaginator Setter
// 分页参数
func (r *AlibabaMosGoodsSearchcspuAPIRequest) SetParamPaginator(_paramPaginator *Paginator) error {
	r._paramPaginator = _paramPaginator
	r.Set("param_paginator", _paramPaginator)
	return nil
}

// GetParamPaginator ParamPaginator Getter
func (r AlibabaMosGoodsSearchcspuAPIRequest) GetParamPaginator() *Paginator {
	return r._paramPaginator
}
