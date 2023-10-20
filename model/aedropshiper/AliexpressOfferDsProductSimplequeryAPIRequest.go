package aedropshiper

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressOfferDsProductSimplequeryAPIRequest Dropshipper查询单个商品的简易信息 API请求
// aliexpress.offer.ds.product.simplequery
//
// 提供给Dropshipper的通过商品ID查找商品简易信息（包括SKU-价格/库存、产品状态等）的接口，只有特定买家可以使用
type AliexpressOfferDsProductSimplequeryAPIRequest struct {
	model.Params
	// 国家
	_localCountry string
	// 语言
	_localLanguage string
	// 商品ID
	_productId int64
}

// NewAliexpressOfferDsProductSimplequeryRequest 初始化AliexpressOfferDsProductSimplequeryAPIRequest对象
func NewAliexpressOfferDsProductSimplequeryRequest() *AliexpressOfferDsProductSimplequeryAPIRequest {
	return &AliexpressOfferDsProductSimplequeryAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AliexpressOfferDsProductSimplequeryAPIRequest) Reset() {
	r._localCountry = ""
	r._localLanguage = ""
	r._productId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressOfferDsProductSimplequeryAPIRequest) GetApiMethodName() string {
	return "aliexpress.offer.ds.product.simplequery"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressOfferDsProductSimplequeryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressOfferDsProductSimplequeryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLocalCountry is LocalCountry Setter
// 国家
func (r *AliexpressOfferDsProductSimplequeryAPIRequest) SetLocalCountry(_localCountry string) error {
	r._localCountry = _localCountry
	r.Set("local_country", _localCountry)
	return nil
}

// GetLocalCountry LocalCountry Getter
func (r AliexpressOfferDsProductSimplequeryAPIRequest) GetLocalCountry() string {
	return r._localCountry
}

// SetLocalLanguage is LocalLanguage Setter
// 语言
func (r *AliexpressOfferDsProductSimplequeryAPIRequest) SetLocalLanguage(_localLanguage string) error {
	r._localLanguage = _localLanguage
	r.Set("local_language", _localLanguage)
	return nil
}

// GetLocalLanguage LocalLanguage Getter
func (r AliexpressOfferDsProductSimplequeryAPIRequest) GetLocalLanguage() string {
	return r._localLanguage
}

// SetProductId is ProductId Setter
// 商品ID
func (r *AliexpressOfferDsProductSimplequeryAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r AliexpressOfferDsProductSimplequeryAPIRequest) GetProductId() int64 {
	return r._productId
}

var poolAliexpressOfferDsProductSimplequeryAPIRequest = sync.Pool{
	New: func() any {
		return NewAliexpressOfferDsProductSimplequeryRequest()
	},
}

// GetAliexpressOfferDsProductSimplequeryRequest 从 sync.Pool 获取 AliexpressOfferDsProductSimplequeryAPIRequest
func GetAliexpressOfferDsProductSimplequeryAPIRequest() *AliexpressOfferDsProductSimplequeryAPIRequest {
	return poolAliexpressOfferDsProductSimplequeryAPIRequest.Get().(*AliexpressOfferDsProductSimplequeryAPIRequest)
}

// ReleaseAliexpressOfferDsProductSimplequeryAPIRequest 将 AliexpressOfferDsProductSimplequeryAPIRequest 放入 sync.Pool
func ReleaseAliexpressOfferDsProductSimplequeryAPIRequest(v *AliexpressOfferDsProductSimplequeryAPIRequest) {
	v.Reset()
	poolAliexpressOfferDsProductSimplequeryAPIRequest.Put(v)
}
