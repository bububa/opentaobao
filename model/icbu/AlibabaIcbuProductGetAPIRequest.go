package icbu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaicbuproductgetAPIRequest 获得单个商品详情 API请求
// alibaba.icbu.product.get
//
// 获取商品详情
type AlibabaicbuproductgetAPIRequest struct {
	model.Params
	// 商品语种，目前只支持ENGLISH
	_language string
	// 混淆后的商品ID
	_productId string
}

// NewAlibabaicbuproductgetRequest 初始化AlibabaicbuproductgetAPIRequest对象
func NewAlibabaicbuproductgetRequest() *AlibabaicbuproductgetAPIRequest {
	return &AlibabaicbuproductgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaicbuproductgetAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.product.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaicbuproductgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaicbuproductgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLanguage is Language Setter
// 商品语种，目前只支持ENGLISH
func (r *AlibabaicbuproductgetAPIRequest) SetLanguage(_language string) error {
	r._language = _language
	r.Set("language", _language)
	return nil
}

// GetLanguage Language Getter
func (r AlibabaicbuproductgetAPIRequest) GetLanguage() string {
	return r._language
}

// SetProductId is ProductId Setter
// 混淆后的商品ID
func (r *AlibabaicbuproductgetAPIRequest) SetProductId(_productId string) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r AlibabaicbuproductgetAPIRequest) GetProductId() string {
	return r._productId
}
