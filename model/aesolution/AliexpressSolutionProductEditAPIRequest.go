package aesolution

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressSolutionProductEditAPIRequest Edit Product API API请求
// aliexpress.solution.product.edit
//
// API for editing product, customized for Oversea merchants. Most of the input fields of this API is similar with aliexpress.solution.product.post. For editing, just fill in the fields you would like to edit, leave other fields to be blank.
type AliexpressSolutionProductEditAPIRequest struct {
	model.Params
	// input param
	_editProductRequest *PostProductRequestDto
}

// NewAliexpressSolutionProductEditRequest 初始化AliexpressSolutionProductEditAPIRequest对象
func NewAliexpressSolutionProductEditRequest() *AliexpressSolutionProductEditAPIRequest {
	return &AliexpressSolutionProductEditAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AliexpressSolutionProductEditAPIRequest) Reset() {
	r._editProductRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressSolutionProductEditAPIRequest) GetApiMethodName() string {
	return "aliexpress.solution.product.edit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressSolutionProductEditAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressSolutionProductEditAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEditProductRequest is EditProductRequest Setter
// input param
func (r *AliexpressSolutionProductEditAPIRequest) SetEditProductRequest(_editProductRequest *PostProductRequestDto) error {
	r._editProductRequest = _editProductRequest
	r.Set("edit_product_request", _editProductRequest)
	return nil
}

// GetEditProductRequest EditProductRequest Getter
func (r AliexpressSolutionProductEditAPIRequest) GetEditProductRequest() *PostProductRequestDto {
	return r._editProductRequest
}

var poolAliexpressSolutionProductEditAPIRequest = sync.Pool{
	New: func() any {
		return NewAliexpressSolutionProductEditRequest()
	},
}

// GetAliexpressSolutionProductEditRequest 从 sync.Pool 获取 AliexpressSolutionProductEditAPIRequest
func GetAliexpressSolutionProductEditAPIRequest() *AliexpressSolutionProductEditAPIRequest {
	return poolAliexpressSolutionProductEditAPIRequest.Get().(*AliexpressSolutionProductEditAPIRequest)
}

// ReleaseAliexpressSolutionProductEditAPIRequest 将 AliexpressSolutionProductEditAPIRequest 放入 sync.Pool
func ReleaseAliexpressSolutionProductEditAPIRequest(v *AliexpressSolutionProductEditAPIRequest) {
	v.Reset()
	poolAliexpressSolutionProductEditAPIRequest.Put(v)
}
