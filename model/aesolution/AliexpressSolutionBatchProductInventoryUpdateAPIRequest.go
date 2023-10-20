package aesolution

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AliexpressSolutionBatchProductInventoryUpdateAPIRequest) Reset() {
	r._mutipleProductUpdateList = r._mutipleProductUpdateList[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressSolutionBatchProductInventoryUpdateAPIRequest) GetApiMethodName() string {
	return "aliexpress.solution.batch.product.inventory.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressSolutionBatchProductInventoryUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressSolutionBatchProductInventoryUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMutipleProductUpdateList is MutipleProductUpdateList Setter
// The product list, in which the inventory needs to be updated. Maximum 20 products.
func (r *AliexpressSolutionBatchProductInventoryUpdateAPIRequest) SetMutipleProductUpdateList(_mutipleProductUpdateList []SynchronizeProductRequestDto) error {
	r._mutipleProductUpdateList = _mutipleProductUpdateList
	r.Set("mutiple_product_update_list", _mutipleProductUpdateList)
	return nil
}

// GetMutipleProductUpdateList MutipleProductUpdateList Getter
func (r AliexpressSolutionBatchProductInventoryUpdateAPIRequest) GetMutipleProductUpdateList() []SynchronizeProductRequestDto {
	return r._mutipleProductUpdateList
}

var poolAliexpressSolutionBatchProductInventoryUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAliexpressSolutionBatchProductInventoryUpdateRequest()
	},
}

// GetAliexpressSolutionBatchProductInventoryUpdateRequest 从 sync.Pool 获取 AliexpressSolutionBatchProductInventoryUpdateAPIRequest
func GetAliexpressSolutionBatchProductInventoryUpdateAPIRequest() *AliexpressSolutionBatchProductInventoryUpdateAPIRequest {
	return poolAliexpressSolutionBatchProductInventoryUpdateAPIRequest.Get().(*AliexpressSolutionBatchProductInventoryUpdateAPIRequest)
}

// ReleaseAliexpressSolutionBatchProductInventoryUpdateAPIRequest 将 AliexpressSolutionBatchProductInventoryUpdateAPIRequest 放入 sync.Pool
func ReleaseAliexpressSolutionBatchProductInventoryUpdateAPIRequest(v *AliexpressSolutionBatchProductInventoryUpdateAPIRequest) {
	v.Reset()
	poolAliexpressSolutionBatchProductInventoryUpdateAPIRequest.Put(v)
}
