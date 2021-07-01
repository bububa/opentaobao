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

// NewAliexpressSolutionProductSchemaGetRequest 初始化AliexpressSolutionProductSchemaGetAPIRequest对象
func NewAliexpressSolutionProductSchemaGetRequest() *AliexpressSolutionProductSchemaGetAPIRequest {
	return &AliexpressSolutionProductSchemaGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressSolutionProductSchemaGetAPIRequest) GetApiMethodName() string {
	return "aliexpress.solution.product.schema.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressSolutionProductSchemaGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is AliexpressCategoryId Setter
// aliexpress category id. You can get it from category API
func (r *AliexpressSolutionProductSchemaGetAPIRequest) SetAliexpressCategoryId(_aliexpressCategoryId int64) error {
	r._aliexpressCategoryId = _aliexpressCategoryId
	r.Set("aliexpress_category_id", _aliexpressCategoryId)
	return nil
}

// Get AliexpressCategoryId Getter
func (r AliexpressSolutionProductSchemaGetAPIRequest) GetAliexpressCategoryId() int64 {
	return r._aliexpressCategoryId
}
