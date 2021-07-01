package aesolution

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressSolutionProductEditAPIRequest
Edit Product API API请求
aliexpress.solution.product.edit

API for editing product, customized for Oversea merchants. Most of the input fields of this API is similar with aliexpress.solution.product.post. For editing, just fill in the fields you would like to edit, leave other fields to be blank. */
type AliexpressSolutionProductEditAPIRequest struct {
	model.Params
	// input param
	_editProductRequest *PostProductRequestDto
}

// New
