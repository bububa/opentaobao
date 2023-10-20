package icbu

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIcbuProductIdDecryptAPIRequest) Reset() {
	r._language = ""
	r._productId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIcbuProductIdDecryptAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.product.id.decrypt"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIcbuProductIdDecryptAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIcbuProductIdDecryptAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaIcbuProductIdDecryptAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIcbuProductIdDecryptRequest()
	},
}

// GetAlibabaIcbuProductIdDecryptRequest 从 sync.Pool 获取 AlibabaIcbuProductIdDecryptAPIRequest
func GetAlibabaIcbuProductIdDecryptAPIRequest() *AlibabaIcbuProductIdDecryptAPIRequest {
	return poolAlibabaIcbuProductIdDecryptAPIRequest.Get().(*AlibabaIcbuProductIdDecryptAPIRequest)
}

// ReleaseAlibabaIcbuProductIdDecryptAPIRequest 将 AlibabaIcbuProductIdDecryptAPIRequest 放入 sync.Pool
func ReleaseAlibabaIcbuProductIdDecryptAPIRequest(v *AlibabaIcbuProductIdDecryptAPIRequest) {
	v.Reset()
	poolAlibabaIcbuProductIdDecryptAPIRequest.Put(v)
}
