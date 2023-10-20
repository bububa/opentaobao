package aesolution

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpressSolutionSellerCategoryTreeQueryAPIRequest aliexpress.solution.seller.category.tree.query API请求
// aliexpress.solution.seller.category.tree.query
//
// API for seller to query the category tree. Support only displaying the categories which seller have permissions to publish products.
type AliexpressSolutionSellerCategoryTreeQueryAPIRequest struct {
	model.Params
	// parent category ID. To obtain the root categories, setting the category_id = 0
	_categoryId int64
	// filter the categories which seller does not have permission
	_filterNoPermission bool
}

// NewAliexpressSolutionSellerCategoryTreeQueryRequest 初始化AliexpressSolutionSellerCategoryTreeQueryAPIRequest对象
func NewAliexpressSolutionSellerCategoryTreeQueryRequest() *AliexpressSolutionSellerCategoryTreeQueryAPIRequest {
	return &AliexpressSolutionSellerCategoryTreeQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressSolutionSellerCategoryTreeQueryAPIRequest) GetApiMethodName() string {
	return "aliexpress.solution.seller.category.tree.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressSolutionSellerCategoryTreeQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressSolutionSellerCategoryTreeQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCategoryId is CategoryId Setter
// parent category ID. To obtain the root categories, setting the category_id = 0
func (r *AliexpressSolutionSellerCategoryTreeQueryAPIRequest) SetCategoryId(_categoryId int64) error {
	r._categoryId = _categoryId
	r.Set("category_id", _categoryId)
	return nil
}

// GetCategoryId CategoryId Getter
func (r AliexpressSolutionSellerCategoryTreeQueryAPIRequest) GetCategoryId() int64 {
	return r._categoryId
}

// SetFilterNoPermission is FilterNoPermission Setter
// filter the categories which seller does not have permission
func (r *AliexpressSolutionSellerCategoryTreeQueryAPIRequest) SetFilterNoPermission(_filterNoPermission bool) error {
	r._filterNoPermission = _filterNoPermission
	r.Set("filter_no_permission", _filterNoPermission)
	return nil
}

// GetFilterNoPermission FilterNoPermission Getter
func (r AliexpressSolutionSellerCategoryTreeQueryAPIRequest) GetFilterNoPermission() bool {
	return r._filterNoPermission
}
