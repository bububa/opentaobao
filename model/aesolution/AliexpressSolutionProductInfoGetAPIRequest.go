package aesolution

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressSolutionProductInfoGetAPIRequest Get Single Product Info API请求
// aliexpress.solution.product.info.get
//
// Get Single Product Info
type AliexpressSolutionProductInfoGetAPIRequest struct {
	model.Params
	// product ID
	_productId int64
}

// NewAliexpressSolutionProductInfoGetRequest 初始化AliexpressSolutionProductInfoGetAPIRequest对象
func NewAliexpressSolutionProductInfoGetRequest() *AliexpressSolutionProductInfoGetAPIRequest {
	return &AliexpressSolutionProductInfoGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AliexpressSolutionProductInfoGetAPIRequest) Reset() {
	r._productId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressSolutionProductInfoGetAPIRequest) GetApiMethodName() string {
	return "aliexpress.solution.product.info.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressSolutionProductInfoGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressSolutionProductInfoGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProductId is ProductId Setter
// product ID
func (r *AliexpressSolutionProductInfoGetAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r AliexpressSolutionProductInfoGetAPIRequest) GetProductId() int64 {
	return r._productId
}

var poolAliexpressSolutionProductInfoGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAliexpressSolutionProductInfoGetRequest()
	},
}

// GetAliexpressSolutionProductInfoGetRequest 从 sync.Pool 获取 AliexpressSolutionProductInfoGetAPIRequest
func GetAliexpressSolutionProductInfoGetAPIRequest() *AliexpressSolutionProductInfoGetAPIRequest {
	return poolAliexpressSolutionProductInfoGetAPIRequest.Get().(*AliexpressSolutionProductInfoGetAPIRequest)
}

// ReleaseAliexpressSolutionProductInfoGetAPIRequest 将 AliexpressSolutionProductInfoGetAPIRequest 放入 sync.Pool
func ReleaseAliexpressSolutionProductInfoGetAPIRequest(v *AliexpressSolutionProductInfoGetAPIRequest) {
	v.Reset()
	poolAliexpressSolutionProductInfoGetAPIRequest.Put(v)
}
