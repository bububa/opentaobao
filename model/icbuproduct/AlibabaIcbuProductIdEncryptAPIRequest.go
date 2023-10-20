package icbuproduct

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaicbuproductidencryptAPIRequest ICBU国际站商品加密接口 API请求
// alibaba.icbu.product.id.encrypt
//
// ICBU国际站，对混淆的产品ID加密。
type AlibabaicbuproductidencryptAPIRequest struct {
	model.Params
	// 语种
	_language string
	// 明文id
	_productId int64
}

// NewAlibabaicbuproductidencryptRequest 初始化AlibabaicbuproductidencryptAPIRequest对象
func NewAlibabaicbuproductidencryptRequest() *AlibabaicbuproductidencryptAPIRequest {
	return &AlibabaicbuproductidencryptAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaicbuproductidencryptAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.product.id.encrypt"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaicbuproductidencryptAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaicbuproductidencryptAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLanguage is Language Setter
// 语种
func (r *AlibabaicbuproductidencryptAPIRequest) SetLanguage(_language string) error {
	r._language = _language
	r.Set("language", _language)
	return nil
}

// GetLanguage Language Getter
func (r AlibabaicbuproductidencryptAPIRequest) GetLanguage() string {
	return r._language
}

// SetProductId is ProductId Setter
// 明文id
func (r *AlibabaicbuproductidencryptAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r AlibabaicbuproductidencryptAPIRequest) GetProductId() int64 {
	return r._productId
}
