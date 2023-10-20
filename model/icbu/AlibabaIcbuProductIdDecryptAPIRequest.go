package icbu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaicbuproductiddecryptAPIRequest 商品ID解密 API请求
// alibaba.icbu.product.id.decrypt
//
// 对混淆的产品ID解密
type AlibabaicbuproductiddecryptAPIRequest struct {
	model.Params
	// 语种
	_language string
	// 混淆后的商品ID
	_productId string
}

// NewAlibabaicbuproductiddecryptRequest 初始化AlibabaicbuproductiddecryptAPIRequest对象
func NewAlibabaicbuproductiddecryptRequest() *AlibabaicbuproductiddecryptAPIRequest {
	return &AlibabaicbuproductiddecryptAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaicbuproductiddecryptAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.product.id.decrypt"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaicbuproductiddecryptAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaicbuproductiddecryptAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLanguage is Language Setter
// 语种
func (r *AlibabaicbuproductiddecryptAPIRequest) SetLanguage(_language string) error {
	r._language = _language
	r.Set("language", _language)
	return nil
}

// GetLanguage Language Getter
func (r AlibabaicbuproductiddecryptAPIRequest) GetLanguage() string {
	return r._language
}

// SetProductId is ProductId Setter
// 混淆后的商品ID
func (r *AlibabaicbuproductiddecryptAPIRequest) SetProductId(_productId string) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r AlibabaicbuproductiddecryptAPIRequest) GetProductId() string {
	return r._productId
}
