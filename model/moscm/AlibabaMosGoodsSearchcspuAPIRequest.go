package moscm

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMosGoodsSearchcspuAPIRequest) Reset() {
	r._paramCspuCriteria = nil
	r._paramPaginator = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMosGoodsSearchcspuAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.goods.searchcspu"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMosGoodsSearchcspuAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMosGoodsSearchcspuAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaMosGoodsSearchcspuAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMosGoodsSearchcspuRequest()
	},
}

// GetAlibabaMosGoodsSearchcspuRequest 从 sync.Pool 获取 AlibabaMosGoodsSearchcspuAPIRequest
func GetAlibabaMosGoodsSearchcspuAPIRequest() *AlibabaMosGoodsSearchcspuAPIRequest {
	return poolAlibabaMosGoodsSearchcspuAPIRequest.Get().(*AlibabaMosGoodsSearchcspuAPIRequest)
}

// ReleaseAlibabaMosGoodsSearchcspuAPIRequest 将 AlibabaMosGoodsSearchcspuAPIRequest 放入 sync.Pool
func ReleaseAlibabaMosGoodsSearchcspuAPIRequest(v *AlibabaMosGoodsSearchcspuAPIRequest) {
	v.Reset()
	poolAlibabaMosGoodsSearchcspuAPIRequest.Put(v)
}
