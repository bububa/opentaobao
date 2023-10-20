package aesolution

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpresssolutionproductlistgetAPIRequest Get product list API请求
// aliexpress.solution.product.list.get
//
// Get product list
type AliexpresssolutionproductlistgetAPIRequest struct {
	model.Params
	// request parameters to query
	_aeopAEProductListQuery *ItemListQuery
}

// NewAliexpresssolutionproductlistgetRequest 初始化AliexpresssolutionproductlistgetAPIRequest对象
func NewAliexpresssolutionproductlistgetRequest() *AliexpresssolutionproductlistgetAPIRequest {
	return &AliexpresssolutionproductlistgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpresssolutionproductlistgetAPIRequest) GetApiMethodName() string {
	return "aliexpress.solution.product.list.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpresssolutionproductlistgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpresssolutionproductlistgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAeopAEProductListQuery is AeopAEProductListQuery Setter
// request parameters to query
func (r *AliexpresssolutionproductlistgetAPIRequest) SetAeopAEProductListQuery(_aeopAEProductListQuery *ItemListQuery) error {
	r._aeopAEProductListQuery = _aeopAEProductListQuery
	r.Set("aeop_a_e_product_list_query", _aeopAEProductListQuery)
	return nil
}

// GetAeopAEProductListQuery AeopAEProductListQuery Getter
func (r AliexpresssolutionproductlistgetAPIRequest) GetAeopAEProductListQuery() *ItemListQuery {
	return r._aeopAEProductListQuery
}
