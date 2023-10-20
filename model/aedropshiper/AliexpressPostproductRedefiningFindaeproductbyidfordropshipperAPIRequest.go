package aedropshiper

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpresspostproductredefiningfindaeproductbyidfordropshipperAPIRequest Dropshipper查找商品信息接口 API请求
// aliexpress.postproduct.redefining.findaeproductbyidfordropshipper
//
// 提供给Dropshipper的通过商品ID查找商品信息的接口，只有特定买家可以使用
type AliexpresspostproductredefiningfindaeproductbyidfordropshipperAPIRequest struct {
	model.Params
	// 国家
	_localCountry string
	// 语言
	_localLanguage string
	// 商品ID
	_productId int64
}

// NewAliexpresspostproductredefiningfindaeproductbyidfordropshipperRequest 初始化AliexpresspostproductredefiningfindaeproductbyidfordropshipperAPIRequest对象
func NewAliexpresspostproductredefiningfindaeproductbyidfordropshipperRequest() *AliexpresspostproductredefiningfindaeproductbyidfordropshipperAPIRequest {
	return &AliexpresspostproductredefiningfindaeproductbyidfordropshipperAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpresspostproductredefiningfindaeproductbyidfordropshipperAPIRequest) GetApiMethodName() string {
	return "aliexpress.postproduct.redefining.findaeproductbyidfordropshipper"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpresspostproductredefiningfindaeproductbyidfordropshipperAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpresspostproductredefiningfindaeproductbyidfordropshipperAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLocalCountry is LocalCountry Setter
// 国家
func (r *AliexpresspostproductredefiningfindaeproductbyidfordropshipperAPIRequest) SetLocalCountry(_localCountry string) error {
	r._localCountry = _localCountry
	r.Set("local_country", _localCountry)
	return nil
}

// GetLocalCountry LocalCountry Getter
func (r AliexpresspostproductredefiningfindaeproductbyidfordropshipperAPIRequest) GetLocalCountry() string {
	return r._localCountry
}

// SetLocalLanguage is LocalLanguage Setter
// 语言
func (r *AliexpresspostproductredefiningfindaeproductbyidfordropshipperAPIRequest) SetLocalLanguage(_localLanguage string) error {
	r._localLanguage = _localLanguage
	r.Set("local_language", _localLanguage)
	return nil
}

// GetLocalLanguage LocalLanguage Getter
func (r AliexpresspostproductredefiningfindaeproductbyidfordropshipperAPIRequest) GetLocalLanguage() string {
	return r._localLanguage
}

// SetProductId is ProductId Setter
// 商品ID
func (r *AliexpresspostproductredefiningfindaeproductbyidfordropshipperAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r AliexpresspostproductredefiningfindaeproductbyidfordropshipperAPIRequest) GetProductId() int64 {
	return r._productId
}
