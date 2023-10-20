package aesolution

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpresssolutionproductpostAPIRequest Product posting API API请求
// aliexpress.solution.product.post
//
// Product posting API for Oversea merchants, simplifying the complexity of integration that sellers and merchants face. For example, these sellers can use their own category and attributes instead of mapping those from AE.
type AliexpresssolutionproductpostAPIRequest struct {
	model.Params
	// input param
	_postProductRequest *PostProductRequestDto
}

// NewAliexpresssolutionproductpostRequest 初始化AliexpresssolutionproductpostAPIRequest对象
func NewAliexpresssolutionproductpostRequest() *AliexpresssolutionproductpostAPIRequest {
	return &AliexpresssolutionproductpostAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpresssolutionproductpostAPIRequest) GetApiMethodName() string {
	return "aliexpress.solution.product.post"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpresssolutionproductpostAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpresssolutionproductpostAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPostProductRequest is PostProductRequest Setter
// input param
func (r *AliexpresssolutionproductpostAPIRequest) SetPostProductRequest(_postProductRequest *PostProductRequestDto) error {
	r._postProductRequest = _postProductRequest
	r.Set("post_product_request", _postProductRequest)
	return nil
}

// GetPostProductRequest PostProductRequest Getter
func (r AliexpresssolutionproductpostAPIRequest) GetPostProductRequest() *PostProductRequestDto {
	return r._postProductRequest
}
