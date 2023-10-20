package moscm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamosgoodssearchcspuAPIRequest cspu查询 API请求
// alibaba.mos.goods.searchcspu
//
// 商品信息查询（仅用于商品上传数据验证，不能用于商品下载，有限流）
type AlibabamosgoodssearchcspuAPIRequest struct {
	model.Params
	// 组合查询对象
	_paramCspuCriteria *CspuCriteria
	// 分页参数
	_paramPaginator *Paginator
}

// NewAlibabamosgoodssearchcspuRequest 初始化AlibabamosgoodssearchcspuAPIRequest对象
func NewAlibabamosgoodssearchcspuRequest() *AlibabamosgoodssearchcspuAPIRequest {
	return &AlibabamosgoodssearchcspuAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamosgoodssearchcspuAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.goods.searchcspu"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamosgoodssearchcspuAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamosgoodssearchcspuAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamCspuCriteria is ParamCspuCriteria Setter
// 组合查询对象
func (r *AlibabamosgoodssearchcspuAPIRequest) SetParamCspuCriteria(_paramCspuCriteria *CspuCriteria) error {
	r._paramCspuCriteria = _paramCspuCriteria
	r.Set("param_cspu_criteria", _paramCspuCriteria)
	return nil
}

// GetParamCspuCriteria ParamCspuCriteria Getter
func (r AlibabamosgoodssearchcspuAPIRequest) GetParamCspuCriteria() *CspuCriteria {
	return r._paramCspuCriteria
}

// SetParamPaginator is ParamPaginator Setter
// 分页参数
func (r *AlibabamosgoodssearchcspuAPIRequest) SetParamPaginator(_paramPaginator *Paginator) error {
	r._paramPaginator = _paramPaginator
	r.Set("param_paginator", _paramPaginator)
	return nil
}

// GetParamPaginator ParamPaginator Getter
func (r AlibabamosgoodssearchcspuAPIRequest) GetParamPaginator() *Paginator {
	return r._paramPaginator
}
