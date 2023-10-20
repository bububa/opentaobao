package aesolution

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressSolutionBatchProductDeleteAPIRequest aliexpress.solution.batch.product.delete API请求
// aliexpress.solution.batch.product.delete
//
// Product delete API. Please note that there is no reverse way to retrieve the products which have been deleted. Use this API in cautious.
type AliexpressSolutionBatchProductDeleteAPIRequest struct {
	model.Params
	// maximum 100
	_productIds []int64
}

// NewAliexpressSolutionBatchProductDeleteRequest 初始化AliexpressSolutionBatchProductDeleteAPIRequest对象
func NewAliexpressSolutionBatchProductDeleteRequest() *AliexpressSolutionBatchProductDeleteAPIRequest {
	return &AliexpressSolutionBatchProductDeleteAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AliexpressSolutionBatchProductDeleteAPIRequest) Reset() {
	r._productIds = r._productIds[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressSolutionBatchProductDeleteAPIRequest) GetApiMethodName() string {
	return "aliexpress.solution.batch.product.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressSolutionBatchProductDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressSolutionBatchProductDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProductIds is ProductIds Setter
// maximum 100
func (r *AliexpressSolutionBatchProductDeleteAPIRequest) SetProductIds(_productIds []int64) error {
	r._productIds = _productIds
	r.Set("product_ids", _productIds)
	return nil
}

// GetProductIds ProductIds Getter
func (r AliexpressSolutionBatchProductDeleteAPIRequest) GetProductIds() []int64 {
	return r._productIds
}

var poolAliexpressSolutionBatchProductDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewAliexpressSolutionBatchProductDeleteRequest()
	},
}

// GetAliexpressSolutionBatchProductDeleteRequest 从 sync.Pool 获取 AliexpressSolutionBatchProductDeleteAPIRequest
func GetAliexpressSolutionBatchProductDeleteAPIRequest() *AliexpressSolutionBatchProductDeleteAPIRequest {
	return poolAliexpressSolutionBatchProductDeleteAPIRequest.Get().(*AliexpressSolutionBatchProductDeleteAPIRequest)
}

// ReleaseAliexpressSolutionBatchProductDeleteAPIRequest 将 AliexpressSolutionBatchProductDeleteAPIRequest 放入 sync.Pool
func ReleaseAliexpressSolutionBatchProductDeleteAPIRequest(v *AliexpressSolutionBatchProductDeleteAPIRequest) {
	v.Reset()
	poolAliexpressSolutionBatchProductDeleteAPIRequest.Put(v)
}
