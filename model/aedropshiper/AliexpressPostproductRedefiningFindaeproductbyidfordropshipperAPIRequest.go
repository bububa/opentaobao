package aedropshiper

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpressPostproductRedefiningFindaeproductbyidfordropshipperAPIRequest Dropshipper查找商品信息接口 API请求
// aliexpress.postproduct.redefining.findaeproductbyidfordropshipper
//
// 提供给Dropshipper的通过商品ID查找商品信息的接口，只有特定买家可以使用
type AliexpressPostproductRedefiningFindaeproductbyidfordropshipperAPIRequest struct {
	model.Params
	// 商品ID
	_productId int64
	// 国家
	_localCountry string
	// 语言
	_localLanguage string
}

// NewAliexpressPostproductRedefiningFindaeproductbyidfordropshipperRequest 初始化AliexpressPostproductRedefiningFindaeproductbyidfordropshipperAPIRequest对象
func NewAliexpressPostproductRedefiningFindaeproductbyidfordropshipperRequest() *AliexpressPostproductRedefiningFindaeproductbyidfordropshipperAPIRequest {
	return &AliexpressPostproductRedefiningFindaeproductbyidfordropshipperAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressPostproductRedefiningFindaeproductbyidfordropshipperAPIRequest) GetApiMethodName() string {
	return "aliexpress.postproduct.redefining.findaeproductbyidfordropshipper"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressPostproductRedefiningFindaeproductbyidfordropshipperAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetProductId is ProductId Setter
// 商品ID
func (r *AliexpressPostproductRedefiningFindaeproductbyidfordropshipperAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r AliexpressPostproductRedefiningFindaeproductbyidfordropshipperAPIRequest) GetProductId() int64 {
	return r._productId
}

// SetLocalCountry is LocalCountry Setter
// 国家
func (r *AliexpressPostproductRedefiningFindaeproductbyidfordropshipperAPIRequest) SetLocalCountry(_localCountry string) error {
	r._localCountry = _localCountry
	r.Set("local_country", _localCountry)
	return nil
}

// GetLocalCountry LocalCountry Getter
func (r AliexpressPostproductRedefiningFindaeproductbyidfordropshipperAPIRequest) GetLocalCountry() string {
	return r._localCountry
}

// SetLocalLanguage is LocalLanguage Setter
// 语言
func (r *AliexpressPostproductRedefiningFindaeproductbyidfordropshipperAPIRequest) SetLocalLanguage(_localLanguage string) error {
	r._localLanguage = _localLanguage
	r.Set("local_language", _localLanguage)
	return nil
}

// GetLocalLanguage LocalLanguage Getter
func (r AliexpressPostproductRedefiningFindaeproductbyidfordropshipperAPIRequest) GetLocalLanguage() string {
	return r._localLanguage
}
