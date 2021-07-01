package aesolution

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressSolutionSellerCategoryTreeQueryAPIRequest
aliexpress.solution.seller.category.tree.query API请求
aliexpress.solution.seller.category.tree.query

API for seller to query the category tree. Support only displaying the categories which seller have permissions to publish products. */
type AliexpressSolutionSellerCategoryTreeQueryAPIRequest struct {
	model.Params
	// parent category ID. To obtain the root categories, setting the category_id = 0
	_categoryId int64
	// filter the categories which seller does not have permission
	_filterNoPermission bool
}

// New
