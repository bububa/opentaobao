package pur

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabapurproductsyncAPIRequest 同步产品 API请求
// alibaba.pur.product.sync
//
// 同步产品
type AlibabapurproductsyncAPIRequest struct {
	model.Params
	// 产品对象
	_accessProductDtos []AccessProductDto
}

// NewAlibabapurproductsyncRequest 初始化AlibabapurproductsyncAPIRequest对象
func NewAlibabapurproductsyncRequest() *AlibabapurproductsyncAPIRequest {
	return &AlibabapurproductsyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabapurproductsyncAPIRequest) GetApiMethodName() string {
	return "alibaba.pur.product.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabapurproductsyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabapurproductsyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAccessProductDtos is AccessProductDtos Setter
// 产品对象
func (r *AlibabapurproductsyncAPIRequest) SetAccessProductDtos(_accessProductDtos []AccessProductDto) error {
	r._accessProductDtos = _accessProductDtos
	r.Set("access_product_dtos", _accessProductDtos)
	return nil
}

// GetAccessProductDtos AccessProductDtos Getter
func (r AlibabapurproductsyncAPIRequest) GetAccessProductDtos() []AccessProductDto {
	return r._accessProductDtos
}
