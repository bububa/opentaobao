package aesolution

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpresssolutionbatchproductpriceupdateAPIRequest aliexpress.solution.batch.product.price.update API请求
// aliexpress.solution.batch.product.price.update
//
// batch product price update operation for oversea sellers
type AliexpresssolutionbatchproductpriceupdateAPIRequest struct {
	model.Params
	// The product list, in which the price needs to be updated. Maximum length:20
	_mutipleProductUpdateList []SynchronizeProductRequestDto
}

// NewAliexpresssolutionbatchproductpriceupdateRequest 初始化AliexpresssolutionbatchproductpriceupdateAPIRequest对象
func NewAliexpresssolutionbatchproductpriceupdateRequest() *AliexpresssolutionbatchproductpriceupdateAPIRequest {
	return &AliexpresssolutionbatchproductpriceupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpresssolutionbatchproductpriceupdateAPIRequest) GetApiMethodName() string {
	return "aliexpress.solution.batch.product.price.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpresssolutionbatchproductpriceupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpresssolutionbatchproductpriceupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMutipleProductUpdateList is MutipleProductUpdateList Setter
// The product list, in which the price needs to be updated. Maximum length:20
func (r *AliexpresssolutionbatchproductpriceupdateAPIRequest) SetMutipleProductUpdateList(_mutipleProductUpdateList []SynchronizeProductRequestDto) error {
	r._mutipleProductUpdateList = _mutipleProductUpdateList
	r.Set("mutiple_product_update_list", _mutipleProductUpdateList)
	return nil
}

// GetMutipleProductUpdateList MutipleProductUpdateList Getter
func (r AliexpresssolutionbatchproductpriceupdateAPIRequest) GetMutipleProductUpdateList() []SynchronizeProductRequestDto {
	return r._mutipleProductUpdateList
}
