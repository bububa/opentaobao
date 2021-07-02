package aesolution

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpressSolutionBatchProductInventoryUpdateAPIRequest aliexpress.solution.batch.product.inventory.update API请求
// aliexpress.solution.batch.product.inventory.update
//
// batch product inventory update API for oversea sellers. Sellers could update multiple skus among multiple products in a single call. Maximum 20 products could be updated at the same time and maximum 200 skus could be updated within one product.
type AliexpressSolutionBatchProductInventoryUpdateAPIRequest struct {
	model.Params
	// The product list, in which the inventory needs to be updated. Maximum 20 products.
	_mutipleProductUpdateList []SynchronizeProductRequestDto
}

// NewAliexpressSolutionBatchProductInventoryUpdateRequest 初始化AliexpressSolutionBatchProductInventoryUpdateAPIRequest对象
func NewAliexpressSolutionBatchProductInventoryUpdateRequest() *AliexpressSolutionBatchProductInventoryUpdateAPIRequest {
	return &AliexpressSolutionBatchProductInventoryUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressSolutionBatchProductInventoryUpdateAPIRequest) GetApiMethodName() string {
	return "aliexpress.solution.batch.product.inventory.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressSolutionBatchProductInventoryUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is MutipleProductUpdateList Setter
// The product list, in which the inventory needs to be updated. Maximum 20 products.
func (r *AliexpressSolutionBatchProductInventoryUpdateAPIRequest) SetMutipleProductUpdateList(_mutipleProductUpdateList []SynchronizeProductRequestDto) error {
	r._mutipleProductUpdateList = _mutipleProductUpdateList
	r.Set("mutiple_product_update_list", _mutipleProductUpdateList)
	return nil
}

// Get MutipleProductUpdateList Getter
func (r AliexpressSolutionBatchProductInventoryUpdateAPIRequest) GetMutipleProductUpdateList() []SynchronizeProductRequestDto {
	return r._mutipleProductUpdateList
}
