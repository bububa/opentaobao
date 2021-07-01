package icbuproduct

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIcbuProductIdEncryptAPIRequest
ICBU国际站商品加密接口 API请求
alibaba.icbu.product.id.encrypt

ICBU国际站，对混淆的产品ID加密。 */
type AlibabaIcbuProductIdEncryptAPIRequest struct {
	model.Params
	// 语种
	_language string
	// 明文id
	_productId int64
}

// NewAlibabaIcbuProductIdEncryptRequest 初始化AlibabaIcbuProductIdEncryptAPIRequest对象
func NewAlibabaIcbuProductIdEncryptRequest() *AlibabaIcbuProductIdEncryptAPIRequest {
	return &AlibabaIcbuProductIdEncryptAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIcbuProductIdEncryptAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.product.id.encrypt"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIcbuProductIdEncryptAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Language Setter
// 语种
func (r *AlibabaIcbuProductIdEncryptAPIRequest) SetLanguage(_language string) error {
	r._language = _language
	r.Set("language", _language)
	return nil
}

// Get Language Getter
func (r AlibabaIcbuProductIdEncryptAPIRequest) GetLanguage() string {
	return r._language
}

// Set is ProductId Setter
// 明文id
func (r *AlibabaIcbuProductIdEncryptAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// Get ProductId Getter
func (r AlibabaIcbuProductIdEncryptAPIRequest) GetProductId() int64 {
	return r._productId
}
