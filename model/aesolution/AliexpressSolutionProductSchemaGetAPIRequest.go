package aesolution

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressSolutionProductSchemaGetAPIRequest
get product schema API请求
aliexpress.solution.product.schema.get

provide a new schema way to post product. With a pair of API, one for getting schema, one for posting instance */
type AliexpressSolutionProductSchemaGetAPIRequest struct {
	model.Params
	// aliexpress category id. You can get it from category API
	_aliexpressCategoryId int64
}

// New
