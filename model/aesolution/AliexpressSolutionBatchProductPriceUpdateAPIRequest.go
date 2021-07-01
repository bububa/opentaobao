package aesolution

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressSolutionBatchProductPriceUpdateAPIRequest
aliexpress.solution.batch.product.price.update API请求
aliexpress.solution.batch.product.price.update

batch product price update operation for oversea sellers */
type AliexpressSolutionBatchProductPriceUpdateAPIRequest struct {
	model.Params
	// The product list, in which the price needs to be updated. Maximum length:20
	_mutipleProductUpdateList []SynchronizeProductRequestDto
}

// NewAliexpressSolutionBatchProductPriceUpdateRequest 初始化AliexpressSolutionBatchProductPriceUpdateAPIRequest对象
func NewAliexpressSolutionBatchProductPriceUpdateRequest() *AliexpressSolutionBatchProductPriceUpdateAPIRequest {
	return &AliexpressSolutionBatchProductPriceUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressSolutionBatchProductPriceUpdateAPIRequest) GetApiMethodName() string {
	return "aliexpress.solution.batch.product.price.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressSolutionBatchProductPriceUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is MutipleProductUpdateList Setter
// The product list, in which the price needs to be updated. Maximum length:20
func (r *AliexpressSolutionBatchProductPriceUpdateAPIRequest) SetMutipleProductUpdateList(_mutipleProductUpdateList []SynchronizeProductRequestDto) error {
	r._mutipleProductUpdateList = _mutipleProductUpdateList
	r.Set("mutiple_product_update_list", _mutipleProductUpdateList)
	return nil
}

// Get MutipleProductUpdateList Getter
func (r AliexpressSolutionBatchProductPriceUpdateAPIRequest) GetMutipleProductUpdateList() []SynchronizeProductRequestDto {
	return r._mutipleProductUpdateList
}
