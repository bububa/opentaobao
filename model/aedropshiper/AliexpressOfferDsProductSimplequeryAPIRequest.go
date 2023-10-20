package aedropshiper

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpressofferdsproductsimplequeryAPIRequest Dropshipper查询单个商品的简易信息 API请求
// aliexpress.offer.ds.product.simplequery
//
// 提供给Dropshipper的通过商品ID查找商品简易信息（包括SKU-价格/库存、产品状态等）的接口，只有特定买家可以使用
type AliexpressofferdsproductsimplequeryAPIRequest struct {
	model.Params
	// 国家
	_localCountry string
	// 语言
	_localLanguage string
	// 商品ID
	_productId int64
}

// NewAliexpressofferdsproductsimplequeryRequest 初始化AliexpressofferdsproductsimplequeryAPIRequest对象
func NewAliexpressofferdsproductsimplequeryRequest() *AliexpressofferdsproductsimplequeryAPIRequest {
	return &AliexpressofferdsproductsimplequeryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressofferdsproductsimplequeryAPIRequest) GetApiMethodName() string {
	return "aliexpress.offer.ds.product.simplequery"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressofferdsproductsimplequeryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressofferdsproductsimplequeryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLocalCountry is LocalCountry Setter
// 国家
func (r *AliexpressofferdsproductsimplequeryAPIRequest) SetLocalCountry(_localCountry string) error {
	r._localCountry = _localCountry
	r.Set("local_country", _localCountry)
	return nil
}

// GetLocalCountry LocalCountry Getter
func (r AliexpressofferdsproductsimplequeryAPIRequest) GetLocalCountry() string {
	return r._localCountry
}

// SetLocalLanguage is LocalLanguage Setter
// 语言
func (r *AliexpressofferdsproductsimplequeryAPIRequest) SetLocalLanguage(_localLanguage string) error {
	r._localLanguage = _localLanguage
	r.Set("local_language", _localLanguage)
	return nil
}

// GetLocalLanguage LocalLanguage Getter
func (r AliexpressofferdsproductsimplequeryAPIRequest) GetLocalLanguage() string {
	return r._localLanguage
}

// SetProductId is ProductId Setter
// 商品ID
func (r *AliexpressofferdsproductsimplequeryAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r AliexpressofferdsproductsimplequeryAPIRequest) GetProductId() int64 {
	return r._productId
}
