package icbuproduct

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuProductIdEncryptAPIRequest ICBU国际站商品加密接口 API请求
// alibaba.icbu.product.id.encrypt
//
// ICBU国际站，对混淆的产品ID加密。
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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIcbuProductIdEncryptAPIRequest) Reset() {
	r._language = ""
	r._productId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIcbuProductIdEncryptAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.product.id.encrypt"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIcbuProductIdEncryptAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIcbuProductIdEncryptAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLanguage is Language Setter
// 语种
func (r *AlibabaIcbuProductIdEncryptAPIRequest) SetLanguage(_language string) error {
	r._language = _language
	r.Set("language", _language)
	return nil
}

// GetLanguage Language Getter
func (r AlibabaIcbuProductIdEncryptAPIRequest) GetLanguage() string {
	return r._language
}

// SetProductId is ProductId Setter
// 明文id
func (r *AlibabaIcbuProductIdEncryptAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r AlibabaIcbuProductIdEncryptAPIRequest) GetProductId() int64 {
	return r._productId
}

var poolAlibabaIcbuProductIdEncryptAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIcbuProductIdEncryptRequest()
	},
}

// GetAlibabaIcbuProductIdEncryptRequest 从 sync.Pool 获取 AlibabaIcbuProductIdEncryptAPIRequest
func GetAlibabaIcbuProductIdEncryptAPIRequest() *AlibabaIcbuProductIdEncryptAPIRequest {
	return poolAlibabaIcbuProductIdEncryptAPIRequest.Get().(*AlibabaIcbuProductIdEncryptAPIRequest)
}

// ReleaseAlibabaIcbuProductIdEncryptAPIRequest 将 AlibabaIcbuProductIdEncryptAPIRequest 放入 sync.Pool
func ReleaseAlibabaIcbuProductIdEncryptAPIRequest(v *AlibabaIcbuProductIdEncryptAPIRequest) {
	v.Reset()
	poolAlibabaIcbuProductIdEncryptAPIRequest.Put(v)
}
