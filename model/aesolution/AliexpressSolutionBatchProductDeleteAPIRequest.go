package aesolution

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpresssolutionbatchproductdeleteAPIRequest aliexpress.solution.batch.product.delete API请求
// aliexpress.solution.batch.product.delete
//
// Product delete API. Please note that there is no reverse way to retrieve the products which have been deleted. Use this API in cautious.
type AliexpresssolutionbatchproductdeleteAPIRequest struct {
	model.Params
	// maximum 100
	_productIds []int64
}

// NewAliexpresssolutionbatchproductdeleteRequest 初始化AliexpresssolutionbatchproductdeleteAPIRequest对象
func NewAliexpresssolutionbatchproductdeleteRequest() *AliexpresssolutionbatchproductdeleteAPIRequest {
	return &AliexpresssolutionbatchproductdeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpresssolutionbatchproductdeleteAPIRequest) GetApiMethodName() string {
	return "aliexpress.solution.batch.product.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpresssolutionbatchproductdeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpresssolutionbatchproductdeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProductIds is ProductIds Setter
// maximum 100
func (r *AliexpresssolutionbatchproductdeleteAPIRequest) SetProductIds(_productIds []int64) error {
	r._productIds = _productIds
	r.Set("product_ids", _productIds)
	return nil
}

// GetProductIds ProductIds Getter
func (r AliexpresssolutionbatchproductdeleteAPIRequest) GetProductIds() []int64 {
	return r._productIds
}
