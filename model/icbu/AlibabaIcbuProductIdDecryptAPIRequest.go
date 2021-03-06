package icbu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuProductIdDecryptAPIRequest 商品ID解密 API请求
// alibaba.icbu.product.id.decrypt
//
// 对混淆的产品ID解密
type AlibabaIcbuProductIdDecryptAPIRequest struct {
	model.Params
	// 语种
	_language string
	// 混淆后的商品ID
	_productId string
}

// NewAlibabaIcbuProductIdDecryptRequest 初始化AlibabaIcbuProductIdDecryptAPIRequest对象
func NewAlibabaIcbuProductIdDecryptRequest() *AlibabaIcbuProductIdDecryptAPIRequest {
	return &AlibabaIcbuProductIdDecryptAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIcbuProductIdDecryptAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.product.id.decrypt"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIcbuProductIdDecryptAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetLanguage is Language Setter
// 语种
func (r *AlibabaIcbuProductIdDecryptAPIRequest) SetLanguage(_language string) error {
	r._language = _language
	r.Set("language", _language)
	return nil
}

// GetLanguage Language Getter
func (r AlibabaIcbuProductIdDecryptAPIRequest) GetLanguage() string {
	return r._language
}

// SetProductId is ProductId Setter
// 混淆后的商品ID
func (r *AlibabaIcbuProductIdDecryptAPIRequest) SetProductId(_productId string) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r AlibabaIcbuProductIdDecryptAPIRequest) GetProductId() string {
	return r._productId
}
