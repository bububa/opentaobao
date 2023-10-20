package pur

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabapurbasketmergeAPIRequest 合并购物车 API请求
// alibaba.pur.basket.merge
//
// 采购商城接入第三方商家合并购物车接口服务
type AlibabapurbasketmergeAPIRequest struct {
	model.Params
	// 合并购物车入参
	_paramMallMergeCartRequestDTO *MallMergeCartRequestDto
}

// NewAlibabapurbasketmergeRequest 初始化AlibabapurbasketmergeAPIRequest对象
func NewAlibabapurbasketmergeRequest() *AlibabapurbasketmergeAPIRequest {
	return &AlibabapurbasketmergeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabapurbasketmergeAPIRequest) GetApiMethodName() string {
	return "alibaba.pur.basket.merge"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabapurbasketmergeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabapurbasketmergeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamMallMergeCartRequestDTO is ParamMallMergeCartRequestDTO Setter
// 合并购物车入参
func (r *AlibabapurbasketmergeAPIRequest) SetParamMallMergeCartRequestDTO(_paramMallMergeCartRequestDTO *MallMergeCartRequestDto) error {
	r._paramMallMergeCartRequestDTO = _paramMallMergeCartRequestDTO
	r.Set("param_mall_merge_cart_request_d_t_o", _paramMallMergeCartRequestDTO)
	return nil
}

// GetParamMallMergeCartRequestDTO ParamMallMergeCartRequestDTO Getter
func (r AlibabapurbasketmergeAPIRequest) GetParamMallMergeCartRequestDTO() *MallMergeCartRequestDto {
	return r._paramMallMergeCartRequestDTO
}
