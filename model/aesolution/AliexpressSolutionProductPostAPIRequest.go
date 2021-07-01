package aesolution

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressSolutionProductPostAPIRequest
Product posting API API请求
aliexpress.solution.product.post

Product posting API for Oversea merchants, simplifying the complexity of integration that sellers and merchants face. For example, these sellers can use their own category and attributes instead of mapping those from AE. */
type AliexpressSolutionProductPostAPIRequest struct {
	model.Params
	// input param
	_postProductRequest *PostProductRequestDto
}

// New
