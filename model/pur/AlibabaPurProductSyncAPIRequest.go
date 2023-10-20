package pur

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaPurProductSyncAPIRequest) Reset() {
	r._accessProductDtos = r._accessProductDtos[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaPurProductSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.pur.product.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaPurProductSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaPurProductSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAccessProductDtos is AccessProductDtos Setter
// 产品对象
func (r *AlibabaPurProductSyncAPIRequest) SetAccessProductDtos(_accessProductDtos []AccessProductDto) error {
	r._accessProductDtos = _accessProductDtos
	r.Set("access_product_dtos", _accessProductDtos)
	return nil
}

// GetAccessProductDtos AccessProductDtos Getter
func (r AlibabaPurProductSyncAPIRequest) GetAccessProductDtos() []AccessProductDto {
	return r._accessProductDtos
}

var poolAlibabaPurProductSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaPurProductSyncRequest()
	},
}

// GetAlibabaPurProductSyncRequest 从 sync.Pool 获取 AlibabaPurProductSyncAPIRequest
func GetAlibabaPurProductSyncAPIRequest() *AlibabaPurProductSyncAPIRequest {
	return poolAlibabaPurProductSyncAPIRequest.Get().(*AlibabaPurProductSyncAPIRequest)
}

// ReleaseAlibabaPurProductSyncAPIRequest 将 AlibabaPurProductSyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaPurProductSyncAPIRequest(v *AlibabaPurProductSyncAPIRequest) {
	v.Reset()
	poolAlibabaPurProductSyncAPIRequest.Put(v)
}
