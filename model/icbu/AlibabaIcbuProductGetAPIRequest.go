package icbu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuProductGetAPIRequest 获得单个商品详情 API请求
// alibaba.icbu.product.get
//
// 获取商品详情
type AlibabaIcbuProductGetAPIRequest struct {
	model.Params
	// 商品语种，目前只支持ENGLISH
	_language string
	// 混淆后的商品ID
	_productId string
}

// NewAlibabaIcbuProductGetRequest 初始化AlibabaIcbuProductGetAPIRequest对象
func NewAlibabaIcbuProductGetRequest() *AlibabaIcbuProductGetAPIRequest {
	return &AlibabaIcbuProductGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIcbuProductGetAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.product.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIcbuProductGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetLanguage is Language Setter
// 商品语种，目前只支持ENGLISH
func (r *AlibabaIcbuProductGetAPIRequest) SetLanguage(_language string) error {
	r._language = _language
	r.Set("language", _language)
	return nil
}

// GetLanguage Language Getter
func (r AlibabaIcbuProductGetAPIRequest) GetLanguage() string {
	return r._language
}

// SetProductId is ProductId Setter
// 混淆后的商品ID
func (r *AlibabaIcbuProductGetAPIRequest) SetProductId(_productId string) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r AlibabaIcbuProductGetAPIRequest) GetProductId() string {
	return r._productId
}
