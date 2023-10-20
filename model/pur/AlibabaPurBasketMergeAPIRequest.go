package pur

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaPurBasketMergeAPIRequest 合并购物车 API请求
// alibaba.pur.basket.merge
//
// 采购商城接入第三方商家合并购物车接口服务
type AlibabaPurBasketMergeAPIRequest struct {
	model.Params
	// 合并购物车入参
	_paramMallMergeCartRequestDTO *MallMergeCartRequestDto
}

// NewAlibabaPurBasketMergeRequest 初始化AlibabaPurBasketMergeAPIRequest对象
func NewAlibabaPurBasketMergeRequest() *AlibabaPurBasketMergeAPIRequest {
	return &AlibabaPurBasketMergeAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaPurBasketMergeAPIRequest) Reset() {
	r._paramMallMergeCartRequestDTO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaPurBasketMergeAPIRequest) GetApiMethodName() string {
	return "alibaba.pur.basket.merge"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaPurBasketMergeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaPurBasketMergeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamMallMergeCartRequestDTO is ParamMallMergeCartRequestDTO Setter
// 合并购物车入参
func (r *AlibabaPurBasketMergeAPIRequest) SetParamMallMergeCartRequestDTO(_paramMallMergeCartRequestDTO *MallMergeCartRequestDto) error {
	r._paramMallMergeCartRequestDTO = _paramMallMergeCartRequestDTO
	r.Set("param_mall_merge_cart_request_d_t_o", _paramMallMergeCartRequestDTO)
	return nil
}

// GetParamMallMergeCartRequestDTO ParamMallMergeCartRequestDTO Getter
func (r AlibabaPurBasketMergeAPIRequest) GetParamMallMergeCartRequestDTO() *MallMergeCartRequestDto {
	return r._paramMallMergeCartRequestDTO
}

var poolAlibabaPurBasketMergeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaPurBasketMergeRequest()
	},
}

// GetAlibabaPurBasketMergeRequest 从 sync.Pool 获取 AlibabaPurBasketMergeAPIRequest
func GetAlibabaPurBasketMergeAPIRequest() *AlibabaPurBasketMergeAPIRequest {
	return poolAlibabaPurBasketMergeAPIRequest.Get().(*AlibabaPurBasketMergeAPIRequest)
}

// ReleaseAlibabaPurBasketMergeAPIRequest 将 AlibabaPurBasketMergeAPIRequest 放入 sync.Pool
func ReleaseAlibabaPurBasketMergeAPIRequest(v *AlibabaPurBasketMergeAPIRequest) {
	v.Reset()
	poolAlibabaPurBasketMergeAPIRequest.Put(v)
}
