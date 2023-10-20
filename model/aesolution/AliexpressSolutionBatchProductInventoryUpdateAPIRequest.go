package aesolution

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpresssolutionbatchproductinventoryupdateAPIRequest aliexpress.solution.batch.product.inventory.update API请求
// aliexpress.solution.batch.product.inventory.update
//
// batch product inventory update API for oversea sellers. Sellers could update multiple skus among multiple products in a single call. Maximum 20 products could be updated at the same time and maximum 200 skus could be updated within one product.
type AliexpresssolutionbatchproductinventoryupdateAPIRequest struct {
	model.Params
	// The product list, in which the inventory needs to be updated. Maximum 20 products.
	_mutipleProductUpdateList []SynchronizeProductRequestDto
}

// NewAliexpresssolutionbatchproductinventoryupdateRequest 初始化AliexpresssolutionbatchproductinventoryupdateAPIRequest对象
func NewAliexpresssolutionbatchproductinventoryupdateRequest() *AliexpresssolutionbatchproductinventoryupdateAPIRequest {
	return &AliexpresssolutionbatchproductinventoryupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpresssolutionbatchproductinventoryupdateAPIRequest) GetApiMethodName() string {
	return "aliexpress.solution.batch.product.inventory.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpresssolutionbatchproductinventoryupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpresssolutionbatchproductinventoryupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMutipleProductUpdateList is MutipleProductUpdateList Setter
// The product list, in which the inventory needs to be updated. Maximum 20 products.
func (r *AliexpresssolutionbatchproductinventoryupdateAPIRequest) SetMutipleProductUpdateList(_mutipleProductUpdateList []SynchronizeProductRequestDto) error {
	r._mutipleProductUpdateList = _mutipleProductUpdateList
	r.Set("mutiple_product_update_list", _mutipleProductUpdateList)
	return nil
}

// GetMutipleProductUpdateList MutipleProductUpdateList Getter
func (r AliexpresssolutionbatchproductinventoryupdateAPIRequest) GetMutipleProductUpdateList() []SynchronizeProductRequestDto {
	return r._mutipleProductUpdateList
}
