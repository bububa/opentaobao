package pur

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaPurProductSyncAPIRequest 同步产品 API请求
// alibaba.pur.product.sync
//
// 同步产品
type AlibabaPurProductSyncAPIRequest struct {
	model.Params
	// 产品对象
	_accessProductDtos []AccessProductDto
}

// NewAlibabaPurProductSyncRequest 初始化AlibabaPurProductSyncAPIRequest对象
func NewAlibabaPurProductSyncRequest() *AlibabaPurProductSyncAPIRequest {
	return &AlibabaPurProductSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaPurProductSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.pur.product.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaPurProductSyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is AccessProductDtos Setter
// 产品对象
func (r *AlibabaPurProductSyncAPIRequest) SetAccessProductDtos(_accessProductDtos []AccessProductDto) error {
	r._accessProductDtos = _accessProductDtos
	r.Set("access_product_dtos", _accessProductDtos)
	return nil
}

// Get AccessProductDtos Getter
func (r AlibabaPurProductSyncAPIRequest) GetAccessProductDtos() []AccessProductDto {
	return r._accessProductDtos
}
