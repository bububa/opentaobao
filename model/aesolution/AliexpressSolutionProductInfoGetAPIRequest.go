package aesolution

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpresssolutionproductinfogetAPIRequest Get Single Product Info API请求
// aliexpress.solution.product.info.get
//
// Get Single Product Info
type AliexpresssolutionproductinfogetAPIRequest struct {
	model.Params
	// product ID
	_productId int64
}

// NewAliexpresssolutionproductinfogetRequest 初始化AliexpresssolutionproductinfogetAPIRequest对象
func NewAliexpresssolutionproductinfogetRequest() *AliexpresssolutionproductinfogetAPIRequest {
	return &AliexpresssolutionproductinfogetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpresssolutionproductinfogetAPIRequest) GetApiMethodName() string {
	return "aliexpress.solution.product.info.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpresssolutionproductinfogetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpresssolutionproductinfogetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProductId is ProductId Setter
// product ID
func (r *AliexpresssolutionproductinfogetAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r AliexpresssolutionproductinfogetAPIRequest) GetProductId() int64 {
	return r._productId
}
